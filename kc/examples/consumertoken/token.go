package main

import (
	"context"
	"log"
	"os"

	"github.com/kollalabs/sdk-go/kc"
)

func main() {
	// Get api key from environment variable
	apiKey := os.Getenv("KC_API_KEY")
	if apiKey == "" {
		log.Fatal("KC_API_KEY not set")
	}

	// Create a new client
	kolla, err := kc.New(apiKey)
	if err != nil {
		log.Fatalf("unable to create kolla connect client: %s\n", err)
	}

	// Get consumer token
	ctx := context.Background()
	consumerToken, err := kolla.GetConsumerToken(ctx, "CONSUMER_ID", "CONSUMER_NAME")
	if err != nil {
		log.Fatalf("unable to get consumer token: %s\n", err)
	}

	log.Printf("ConsumerToken: %s\n", consumerToken)
}
