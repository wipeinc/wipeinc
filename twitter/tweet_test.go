package twitter_test

import (
	"testing"

	"github.com/sanity-io/litter"
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
		litter.Dump(tweet)
	}
}
