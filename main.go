package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

var (
	consumerKey       = os.Getenv("CONSUMER_KEY")
	consumerSecret    = os.Getenv("CONSUMER_SECRET")
	accessToken       = os.Getenv("ACCESS_TOKEN")
	accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
)

func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	res, _ := api.GetSearch("golang", nil)
	fmt.Print(res)
	for _, tweet := range res.Statuses {
		fmt.Print(tweet.Text)
	}
}
