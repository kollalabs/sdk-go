package kc

import (
	"context"
	"runtime/debug"

	"github.com/kollalabs/sdk-go/kc/swagger"
)

const BaseURL = "https://api.getkolla.com/connect"

var userAgent = setUserAgent()

// Client struct for kc
type Client struct {
	APIKey string
	// Swagger Client
	OpenAPIClient *swagger.APIClient
}

func New(apikey string) (*Client, error) {

	client := &Client{}
	client.APIKey = apikey

	// Initialize Swagger Client and store it
	configuration := swagger.NewConfiguration()
	configuration.AddDefaultHeader("Authorization", "Bearer "+apikey)
	configuration.BasePath = BaseURL
	configuration.UserAgent = userAgent

	client.OpenAPIClient = swagger.NewAPIClient(configuration)

	return client, nil
}

// ConsumerToken fetches a consumer token from the KC api which is used to initiate the embedded marketplace
func (c *Client) ConsumerToken(ctx context.Context, consumerID string, consumerName string) (string, error) {
	// Create consumer token request
	req := swagger.ConsumerTokenRequest{
		ConsumerId: consumerID,
		Metadata: &swagger.ConsumerTokenRequestConsumerMetadata{
			TenantDisplayName: consumerName,
		},
	}

	// Get consumer token
	consumerTokenResponse, _, err := c.OpenAPIClient.ConnectApi.ConnectConsumerToken(ctx, req)
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
	creds := &Credentials{}

	req := swagger.CredentialsRequest{
		ConsumerId:    consumerID,
		LinkedAccount: "",
	}

	// Get credentials
	lacreds, _, err := c.OpenAPIClient.ConnectApi.ConnectCredentials(ctx, req, connectorID, "-")
	if err != nil {
		return creds, err
	}
	creds.Token = lacreds.Credentials.Token
	creds.ExpiryTime = lacreds.Credentials.ExpiryTime
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
