package twitter_test

import (
	"encoding/json"
	"io/ioutil"
	"log"

	twitterGo "github.com/dghubble/go-twitter/twitter"
)

// KimTimeline is for stats test
var KimTimeline []twitterGo.Tweet

// TwitBird is a tweet for tests
var TwitBird twitterGo.Tweet

var ITweets map[string]twitterGo.Tweet

func init() {
	KimTimeline = loadTweets("fixtures/kim_timeline.json")
	ITweets = map[string]twitterGo.Tweet{
		"bird":      loadTweet("fixtures/tweetbird.json"),
		"overwatch": loadTweet("fixtures/overwatch.json"),
	}
}

func loadTweet(filename string) twitterGo.Tweet {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("err: %s\n", filename)
		log.Fatalf("error loading fixture: %s\n", err)
	}

	var tweet twitterGo.Tweet
	json.Unmarshal(raw, &tweet)
	return tweet
}

func loadTweets(filename string) []twitterGo.Tweet {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("err: %s\n", filename)
		log.Fatalf("error loading fixture: %s\n", err)
	}

	var tweets []twitterGo.Tweet
	json.Unmarshal(raw, &tweets)
	return tweets
}
