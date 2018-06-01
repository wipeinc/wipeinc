package twitter_test

import (
	"encoding/json"
	"io/ioutil"
	"log"

	twitterGo "github.com/dghubble/go-twitter/twitter"
)

// KimTimeline is for stats test
var KimTimeline []twitterGo.Tweet

func init() {
	KimTimeline = loadTweets("fixtures/kim_timeline.json")
}

func loadTweets(filename string) []twitterGo.Tweet {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("err: %s\n", filename)
		log.Fatalf("error loading fixutre: %s\n", err)
	}

	var tweets []twitterGo.Tweet
	json.Unmarshal(raw, &tweets)
	return tweets
}
