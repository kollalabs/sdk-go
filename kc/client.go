package kc

import (
	"context"

	"github.com/kollalabs/sdk-go/kc/swagger"
)

// Client struct for kc
type Client struct {
	APIKey string
	// Swagger Client
	sc *swagger.APIClient
}

// LinkedAccount
type LinkedAccount struct {
	swagger.LinkedAccount
}

// LinkedAccount Credentials
type Credentials struct {
	swagger.LinkedAccountCredentialsResponseCredentials
	LinkedAccount *LinkedAccount
}

func New(apikey string) *Client {
	client := &Client{}
	client.APIKey = apikey

	// Initialize Swagger Client and store it
	configuration := swagger.NewConfiguration()
	configuration.AddDefaultHeader("Authorization", "Bearer "+apikey)
	client.sc = swagger.NewAPIClient(configuration)

	return client
}

// Get consumer token from KC api
func (c *Client) GetConsumerToken(consumerID string, consumerName string) (string, error) {
	// Create empty context
	ctx := context.Background()
	// Create consumer token request
	req := swagger.ConsumerTokenRequest{
		ConsumerId: consumerID,
		Metadata: &swagger.ConsumerTokenRequestConsumerMetadata{
			TenantDisplayName: consumerName,
		},
	}

	// Get consumer token
	consumerTokenResponse, _, err := c.sc.ConnectApi.ConnectConsumerToken(ctx, req)
	if err != nil {
		return "", err
	}
	return consumerTokenResponse.Token, nil
}

func (c *Client) GetCredentials(connectorID string, consumerID string) (*Credentials, error) {
	creds := &Credentials{}
	// Create empty context
	ctx := context.Background()

	req := swagger.LinkedAccountCredentialsRequest{
		ConsumerId:    consumerID,
		LinkedAccount: "",
	}

	// Get credentials
	lacreds, _, err := c.sc.ConnectApi.ConnectLinkedAccountCredentials(ctx, req, connectorID, "-")
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
