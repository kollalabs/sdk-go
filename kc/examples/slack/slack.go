package main

import (
	"context"
	"os"

	"github.com/kollalabs/sdk-go/kc"
	"github.com/slack-go/slack"
)

func main() {
	// Get api key from environment variable
	apiKey := os.Getenv("KC_API_KEY")
	// Create a new client
	kolla := kc.New(apiKey)

	ctx := context.Background()
	creds, err := kolla.GetCredentials(ctx, "slack", "CONSUMER_ID")
	if err != nil {
		panic(err)
	}

	slackapi := slack.New(creds.Token)
	slackapi.PostMessage("CHANNEL_ID", slack.MsgOptionText("Hello world", false))
}
