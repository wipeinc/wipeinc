package twitter_test

import (
	"log"
	"testing"

	"github.com/wipeinc/wipeinc/twitter"
)

func TestAnalyzeTweets(t *testing.T) {
	stats := twitter.AnalyzeUserTweets(KimTimeline)

	const mostPopularTweetID = 997850510219620353
	if len(stats.MostPopularTweets) != 20 {
		t.Errorf("expected %d Most popular tweets got %d\n",
			20, len(stats.MostPopularTweets))
	}

	if stats.MostPopularTweets[0].ID != mostPopularTweetID {
		t.Errorf("expected most pouplar tweet: %d\n", mostPopularTweetID)
		t.Errorf("got: %d\n", stats.MostPopularTweets[0].ID)
	}
	log.Printf("%+v", stats.TopHashtags)
}