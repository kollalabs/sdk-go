package kc

import (
	"context"
	"os"
	"testing"

	"github.com/kollalabs/sdk-go/kc/swagger"
)

// These are integration tests that require a valid API key to be set
// in the environment variable KC_API_KEY
func TestClient_ConsumerToken(t *testing.T) {

	// Get API key from environment variable
	apiKey := os.Getenv("KC_API_KEY")
	if apiKey == "" {
		t.Skip("Skipping integration test. Set KC_API_KEY environment variable to run integration tests")
	}

	type fields struct {
		APIKey        string
		OpenAPIClient *swagger.APIClient
	}
	type args struct {
		ctx          context.Context
		consumerID   string
		consumerName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				APIKey:        tt.fields.APIKey,
				OpenAPIClient: tt.fields.OpenAPIClient,
			}
			got, err := c.ConsumerToken(tt.args.ctx, tt.args.consumerID, tt.args.consumerName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ConsumerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.ConsumerToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
