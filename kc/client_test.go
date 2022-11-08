package kc

import (
	"context"
	"os"
	"testing"
)

// These are integration tests that require a valid API key to be set
// in the environment variable KC_API_KEY

// TestClientConsumerToken tests getting a kc consumer token
func TestClientConsumerToken(t *testing.T) {

	// Get API key from environment variable
	apiKey := os.Getenv("KC_API_KEY")
	if apiKey == "" {
		t.Skip("Skipping integration test. Set KC_API_KEY environment variable to run integration tests")
	}

	ctx := context.Background()

	type args struct {
		ctx          context.Context
		consumerID   string
		consumerName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestClientConsumerTokenValid",
			args: args{
				ctx:          ctx,
				consumerID:   "test-consumer-id",
				consumerName: "test-consumer-name",
			},
			wantErr: false,
		},
		{
			name: "TestClientConsumerTokenNoConsumerID",
			args: args{
				ctx:          ctx,
				consumerID:   "",
				consumerName: "",
			},
			wantErr: true,
		},
	}

	c, err := New(apiKey)
	if err != nil {
		t.Fatalf("unable to load kc client: %s\n", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.ConsumerToken(tt.args.ctx, tt.args.consumerID, tt.args.consumerName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ConsumerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Errorf("Client.ConsumerToken() = [%v]", got)
			}
		})
	}
}
