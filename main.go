package main

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
)

const (
	consumerKey       = ""
	consumerSecret    = ""
	accessToken       = ""
	accessTokenSecret = ""
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
