package twitter_test

import (
	"testing"

	"github.com/wipeinc/wipeinc/twitter"
)

func TestNewTweet(t *testing.T) {
	for name, apiTweet := range ITweets {
		tweet, err := twitter.NewTweet(apiTweet)
		if err != nil {
			t.Fatalf("unexpected error from NewTweet(%s): %s\n",
				name,
				err.Error())
		}
		if tweet.User == nil || tweet.User.ID == 0 {
			t.Fatalf("tweet User is invalid NewTweet(%s)", name)
		}
	}
}
