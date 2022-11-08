package kc

import (
	"context"

	"github.com/kollalabs/sdk-go/kc/swagger"
)

const BaseURL = "https://api.getkolla.com/connect"

// Client struct for kc
type Client struct {
	APIKey string
	// Swagger Client
	OpenAPIClient *swagger.APIClient
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

func New(apikey string) *Client {
	client := &Client{}
	client.APIKey = apikey

	// Initialize Swagger Client and store it
	configuration := swagger.NewConfiguration()
	configuration.AddDefaultHeader("Authorization", "Bearer "+apikey)
	configuration.BasePath = BaseURL

	client.OpenAPIClient = swagger.NewAPIClient(configuration)

	return client
}

// Get consumer token from KC api
func (c *Client) GetConsumerToken(ctx context.Context, consumerID string, consumerName string) (string, error) {
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

func (c *Client) GetCredentials(ctx context.Context, connectorID string, consumerID string) (*Credentials, error) {
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
