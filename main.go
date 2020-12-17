package main

import (
	"fmt"

	gt "github.com/bas24/googletranslatefree"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	// setup
	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("accessToken", "accessSecret")
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// get tweet
	tweet, resp, err := client.Statuses.Show(585613041028431872, nil)

	// translate twice
	content := tweet[-1]
	content = gt.Translate(content, "ja", "en")
	content = gt.Translate(content, "en", "ja")
	fmt.Println(content)

	// tweet the generated content
	tweet, resp, err := client.Statuses.Update(content, nil)
}
