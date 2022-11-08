package main

import (
	"context"
	"os"

	"github.com/kollalabs/sdk-go/kc"
)

func main() {
	// Get api key from environment variable
	apiKey := os.Getenv("KC_API_KEY")
	// Create a new client
	kolla := kc.New(apiKey)

	// Get consumer token
	ctx := context.Background()
	consumerToken, err := kolla.GetConsumerToken(ctx, "CONSUMER_ID", "CONSUMER_NAME")
	if err != nil {
		panic(err)
	}
	println(consumerToken)
}
