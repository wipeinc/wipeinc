package twitter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wipeinc/wipeinc/entity"
	"github.com/wipeinc/wipeinc/twitter"
)

func tweetWithMentions(userMentions []int64) entity.Tweet {
	tweet := entity.Tweet{}
	tweet.User = &entity.User{ID: 5555555}
	tweet.UserMentions = make([]*entity.User, 0, len(userMentions))
	for _, userMention := range userMentions {
		tweet.UserMentions = append(tweet.UserMentions,
			&entity.User{ID: userMention})
	}
	return tweet
}

func tweetWithMention(userMention int64) entity.Tweet {
	if userMention == 0 {
		userMention = 12345
	}
	tweet := entity.Tweet{}
	tweet.User = &entity.User{ID: 5555555}
	tweet.UserMentions = make([]*entity.User, 0, 1)
	tweet.UserMentions = append(tweet.UserMentions, &entity.User{ID: userMention})
	return tweet
}

var tweetWithHashtagTests = []struct {
	description string
	len         int
	in          []entity.Tweet
	out         []twitter.Freq
}{
	{
		"tweets with no hashtag",
		0,
		[]entity.Tweet{entity.Tweet{}, entity.Tweet{}},
		[]twitter.Freq{},
	},
}

var tweetWithMentionTests = []struct {
	description string
	len         int
	in          []entity.Tweet
	out         []twitter.Freq
}{
	{
		"tweets with no mention",
		0,
		[]entity.Tweet{entity.Tweet{}, entity.Tweet{}},
		[]twitter.Freq{},
	},
	{
		"tweets with one mention",
		0,
		[]entity.Tweet{
			tweetWithMention(1),
			tweetWithMention(1),
			entity.Tweet{},
		},
		[]twitter.Freq{twitter.Freq{Value: "1", F: 2}},
	},
	{
		"with 3 mentions",
		0,
		[]entity.Tweet{
			tweetWithMention(1),
			tweetWithMention(1),
			tweetWithMention(1),
			tweetWithMention(2),
			tweetWithMention(2),
			tweetWithMention(3),
			entity.Tweet{},
		},
		[]twitter.Freq{
			twitter.Freq{Value: "1", F: 3},
			twitter.Freq{Value: "2", F: 2},
			twitter.Freq{Value: "3", F: 1},
		},
	},
	{
		"with 3 tweets, 2 on the same, mentions",
		0,
		[]entity.Tweet{
			tweetWithMentions([]int64{1, 2}),
			tweetWithMention(1),
			tweetWithMention(1),
			tweetWithMention(2),
			tweetWithMention(3),
			entity.Tweet{},
		},
		[]twitter.Freq{
			twitter.Freq{Value: "1", F: 3},
			twitter.Freq{Value: "2", F: 2},
			twitter.Freq{Value: "3", F: 1},
		},
	},
	{
		"with 3 mentions and len 2",
		2,
		[]entity.Tweet{
			tweetWithMention(1),
			tweetWithMention(1),
			tweetWithMention(1),
			tweetWithMention(2),
			tweetWithMention(2),
			tweetWithMention(3),
			entity.Tweet{},
		},
		[]twitter.Freq{
			twitter.Freq{Value: "1", F: 3},
			twitter.Freq{Value: "2", F: 2},
		},
	},
}

func TestAnalyzeTweetWithHashtags(t *testing.T) {
	for _, tt := range tweetWithHashtagTests {
		stats := twitter.NewTweetStats()
		stats.AnalyzeTweets(tt.in)
		if !assert.Equal(t, tt.out, stats.TopHashtags(tt.len)) {
			t.Errorf("test: '%s' failed\n", tt.description)
		}
	}
}

func TestAnalyzeTweetWithMentions(t *testing.T) {
	for _, tt := range tweetWithMentionTests {
		stats := twitter.NewTweetStats()
		stats.AnalyzeTweets(tt.in)
		if !assert.Equal(t, tt.out, stats.TopMentions(tt.len)) {
			t.Errorf("test: '%s' failed\n", tt.description)
		}
	}
}

func TestAnalyzeTweetsIntegration(t *testing.T) {
	stats := twitter.NewTweetStats()
	timeline := make([]entity.Tweet, 0, len(KimTimeline))
	for _, apiTweet := range KimTimeline {
		tweet, err := twitter.NewTweet(apiTweet)
		if err != nil {
			t.Fatalf("error loding kim timeline tweet: %d, err: %s",
				apiTweet.ID, err.Error())
		}
		timeline = append(timeline, *tweet)
	}
	stats.AnalyzeTweets(timeline)

	const mostPopularTweetID = 997850510219620353
	if len(stats.MostPopularTweets) != twitter.MostPopularTweetsLen {
		t.Errorf("expected %d Most popular tweets got %d\n",
			twitter.MostPopularTweetsLen, len(stats.MostPopularTweets))
	}

	if stats.MostPopularTweets[0].ID != mostPopularTweetID {
		t.Errorf("expected most pouplar tweet: %d\n", mostPopularTweetID)
		t.Errorf("got: %d\n", stats.MostPopularTweets[0].ID)
	}
}
