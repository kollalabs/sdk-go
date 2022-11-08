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
	creds, err := kolla.GetCredentials(ctx, "slack", "customer-id-kolla")
	if err != nil {
		panic(err)
	}
	//creds.LinkedAccount.AuthData
	slackapi := slack.New(creds.Token)
	_, _, err = slackapi.PostMessage("general", slack.MsgOptionText("Hello world! (Send with Kolla managed token)", false))
	if err != nil {
		panic(err)
	}
}
