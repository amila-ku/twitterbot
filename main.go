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

	log := &logger{logrus.New()}
	api.SetLogger(log)

	stream := api.PublicStreamFilter(url.Values{
		"track": []string{"#golang"},
	})

	defer stream.Stop()

	for v := range stream.C {
		fmt.Println("starting")

		t, ok := v.(anaconda.Tweet)

		if !ok {
			logrus.Warningf("%T", v)
			continue
		}
		_, err := api.Retweet(t.Id, false)
		if err != nil {
			logrus.Errorf("Error RT %d: %v", t.Id, err)
			continue
		}
		logrus.Infof("Retweeted %d", t.Id)
		fmt.Printf("%s\n", t.Text)
	}

	// for v := range stream.C {
	// 	fmt.Printf("%v\n", v)
	// }

	// res, _ := api.GetSearch("golang", nil)
	// fmt.Print(res)
	// for _, tweet := range res.Statuses {
	// 	fmt.Print(tweet.Text)
	// }
}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{}) {
	log.Error(args...)
}

func (log *logger) Criticalf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func (log *logger) Notice(args ...interface{}) {
	log.Error(args...)
}

func (log *logger) Noticef(format string, args ...interface{}) {
	log.Errorf(format, args...)
}
