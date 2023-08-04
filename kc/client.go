package kc

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/kollalabs/sdk-go/kc/swagger"
)

const baseURL = "https://api.getkolla.com/connect"

var userAgent = setUserAgent()

// Client struct for kc
type Client struct {
	// Swagger Client
	openAPIClient *swagger.APIClient
}

// New returns a new KC client, a configuration option modifier function
// may be provided to configure the underlying client
func New(apikey string, opts ...ConfigurationOption) (*Client, error) {

	client := &Client{}

	// Initialize Swagger Client and store it
	configuration := swagger.NewConfiguration()
	configuration.BasePath = baseURL

	for _, v := range opts {
		// apply configuration options
		v(configuration)
	}

	// override the useragent and authorization header
	configuration.AddDefaultHeader("Authorization", "Bearer "+apikey)
	configuration.UserAgent = userAgent
	client.openAPIClient = swagger.NewAPIClient(configuration)

	return client, nil
}

// ConfigurationOption is a function that can be used to configure the swagger client
type ConfigurationOption func(*swagger.Configuration)

// ConsumerToken fetches a consumer token from the KC api which is used to initiate the embedded marketplace
func (c *Client) ConsumerToken(ctx context.Context, consumerID string, consumerName string) (string, error) {
	if c == nil {
		return "", fmt.Errorf("kolla sdk client is nil")
	}
	if c.openAPIClient == nil || c.openAPIClient.ConnectApi == nil {
		return "", fmt.Errorf("internal kolla http client is not configured")
	}

	// Create consumer token request
	req := swagger.ConsumerTokenRequest{
		ConsumerId: consumerID,
		Metadata: &swagger.ConsumerMetadata{
			Title: consumerName,
		},
	}

	// Get consumer token
	consumerTokenResponse, _, err := c.openAPIClient.ConnectApi.ConnectConsumerToken(ctx, req)
	if err != nil {
		return "", err
	}
	return consumerTokenResponse.Token, nil
}

// LinkedAccount
type LinkedAccount struct {
	swagger.LinkedAccount
}

// LinkedAccount Credentials
type Credentials struct {
	swagger.CredentialsResponseCredentials
	LinkedAccount *LinkedAccount
}

// Credentials returns the credentials for a given consumer and connector
func (c *Client) Credentials(ctx context.Context, connectorID string, consumerID string) (*Credentials, error) {
	if c == nil || c.openAPIClient == nil {
		return nil, fmt.Errorf("kolla sdk client is nil")
	}
	if c.openAPIClient == nil || c.openAPIClient.ConnectApi == nil {
		return nil, fmt.Errorf("internal kolla http client is not configured")
	}

	creds := &Credentials{}

	req := swagger.CredentialsRequest{
		ConsumerId:    consumerID,
		LinkedAccount: "",
	}

	// Get credentials
	lacreds, _, err := c.openAPIClient.ConnectApi.ConnectCredentials(ctx, req, connectorID, "-")
	if err != nil {
		return creds, err
	}

	creds.Token = lacreds.Credentials.Token
	creds.ExpiryTime = lacreds.Credentials.ExpiryTime
	creds.Secrets = lacreds.Credentials.Secrets

	creds.LinkedAccount = &LinkedAccount{
		LinkedAccount: *lacreds.LinkedAccount,
	}

	return creds, nil
}

func setUserAgent() string {
	userAgent := ""
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		userAgent = "kolla-sdk-go/unknown"
	} else {
		// look for the kolla sdk dep
		for _, v := range buildInfo.Deps {
			if v.Path == "github.com/kollalabs/sdk-go" {
				userAgent = "kolla-sdk-go/" + v.Version + " go/" + buildInfo.GoVersion
				break
			}
		}
	}
	return userAgent
}
