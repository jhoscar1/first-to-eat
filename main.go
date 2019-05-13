package main

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	cKey := os.Getenv("CONSUMER_KEY")
	cSecret := os.Getenv("CONSUMER_SECRET")
	authToken := os.Getenv("ACCESS_TOKEN")
	authSecret := os.Getenv("ACCESS_SECRET")
	config := oauth1.NewConfig(cKey, cSecret)
	token := oauth1.NewToken(authToken, authSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
}
