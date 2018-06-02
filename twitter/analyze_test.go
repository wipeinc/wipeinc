package twitter_test

import (
	"testing"

	twitterGo "github.com/dghubble/go-twitter/twitter"
	"github.com/stretchr/testify/assert"
	"github.com/wipeinc/wipeinc/twitter"
)

func tweetWithMention(mention string) twitterGo.Tweet {
	if mention == "" {
		mention = "mention"
	}
	tweet := twitterGo.Tweet{}
	tweet.Entities = &twitterGo.Entities{}
	entity := twitterGo.MentionEntity{IDStr: mention}
	tweet.Entities.UserMentions = append(tweet.Entities.UserMentions, entity)
	return tweet
}

var tweetWithMentionTests = []struct {
	description string
	in          []twitterGo.Tweet
	out         []twitter.Freq
}{
	{
		"tweets with no mention",
		[]twitterGo.Tweet{twitterGo.Tweet{}, twitterGo.Tweet{}},
		[]twitter.Freq{},
	},
	{
		"tweets with one mention",
		[]twitterGo.Tweet{
			tweetWithMention(""),
			tweetWithMention(""),
			twitterGo.Tweet{},
		},
		[]twitter.Freq{twitter.Freq{Value: "mention", F: 2}},
	},
}

func TestAnalyzeTweetWithMentions(t *testing.T) {
	for _, tt := range tweetWithMentionTests {
		stats := twitter.NewTweetStats()
		stats.AnalyzeTweets(tt.in)
		if !assert.Equal(t, tt.out, stats.TopMentions(0)) {
			t.Errorf("test: '%s' failed\n", tt.description)
		}
	}
}

func TestAnalyzeTweetsIntegration(t *testing.T) {
	stats := twitter.NewTweetStats()
	stats.AnalyzeTweets(KimTimeline)

	const mostPopularTweetID = 997850510219620353
	if len(stats.MostPopularTweets) != 20 {
		t.Errorf("expected %d Most popular tweets got %d\n",
			20, len(stats.MostPopularTweets))
	}

	if stats.MostPopularTweets[0].ID != mostPopularTweetID {
		t.Errorf("expected most pouplar tweet: %d\n", mostPopularTweetID)
		t.Errorf("got: %d\n", stats.MostPopularTweets[0].ID)
	}
}
