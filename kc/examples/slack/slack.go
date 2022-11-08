package main

import (
	"context"
	"log"
	"os"

	"github.com/kollalabs/sdk-go/kc"
	"github.com/slack-go/slack"
)

func main() {
	ctx := context.Background()

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

	creds, err := kolla.GetCredentials(ctx, "slack", "customer-id-kolla")
	if err != nil {
		log.Fatalf("unable to load credentials %s\n", err)
	}

	slackapi := slack.New(creds.Token)
	_, _, err = slackapi.PostMessage("general", slack.MsgOptionText("Hello world! (Send with Kolla managed token)", false))
	if err != nil {
		log.Fatalf("unable to post slack message: %s\n", err)
	}
}
