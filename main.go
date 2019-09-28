package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
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

	stream := api.PublicStreamFilter(url.Values{
		"track": []string{"#world"},
	})

	defer stream.Stop()

	for v := range stream.C {
		fmt.Println("starting")

		t, ok := v.(anaconda.Tweet)

		if !ok {
			logrus.Warningf("%T", v)
		}
		fmt.Printf("%s\n", t.Text)
	}

	// res, _ := api.GetSearch("golang", nil)
	// fmt.Print(res)
	// for _, tweet := range res.Statuses {
	// 	fmt.Print(tweet.Text)
	// }
}
