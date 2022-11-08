# Kolla Go SDK

Kolla helps you manage all your software integrations in one place! See more at [https://getkolla.com](https://getkolla.com)

This SDK is for KollaConnect, which helps you write your software integrations faster by managing all the
authentication and credentials for you, so you can focus on your business logic. Learn more with at the [Kolla Docs Website](https://docs.getkolla.com)

## Usage

### getConsumerToken

`getConsumerToken` takes a `consumer_id` and `consumer_name` and generates a `consumer_token` for use in the front-end web SDK. (see [Installing KollaConnect](https://docs.getkolla.com/kolla/getting-started/installing-kollaconnect) for more info)

```go
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
```

### getCredentials

`getCredentials` returns the credentials for your customers' software integrations. Behind the scenes Kolla manages OAuth2 negotiation and token refreshing so you can simply get your token and get going. It accepts a `connectorID` and a `consumerID`, which was set when creating the consumerToken. (see [Installing KollaConnect](https://docs.getkolla.com/kolla/getting-started/installing-kollaconnect) for more info)

Here is an example of getting a customer's Slack API token and calling the Slack API with the Slack Go SDK:

```go
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
	creds, err := kolla.GetCredentials(ctx, "slack", "CONSUMER_ID") // Use consumer ID set in consumer token
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
```



