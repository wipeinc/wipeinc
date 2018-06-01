package twitter

import (
	"sort"

	twitterGo "github.com/dghubble/go-twitter/twitter"
)

const favoriteRetweetRatio = 9

// TweetStats struct returned for twitter statistic anlytics
type TweetStats struct {
	MostPopularTweets []twitterGo.Tweet
}

func tweetPopularityScore(tweet twitterGo.Tweet) int {
	return tweet.FavoriteCount + tweet.RetweetCount*9
}

// NewTweetStats Create a New TweetStats struct
func NewTweetStats() *TweetStats {
	return &TweetStats{
		MostPopularTweets: []twitterGo.Tweet{},
	}
}

// AnalyzeUserTweets return a TweetStats structure of the analyzed tweets
func AnalyzeUserTweets(tweets []twitterGo.Tweet) *TweetStats {
	s := NewTweetStats()
	for _, tweet := range tweets {
		s.AnalyzeTweet(tweet)
	}
	return s
}

func (s *TweetStats) updateTopTweets(tweet twitterGo.Tweet) {
	score := tweetPopularityScore(tweet)
	index := sort.Search(len(s.MostPopularTweets), func(i int) bool {
		return tweetPopularityScore(s.MostPopularTweets[i]) < score
	})
	if index < len(s.MostPopularTweets) {
		copy(s.MostPopularTweets[index+1:], s.MostPopularTweets[index:])
		s.MostPopularTweets[index] = tweet
		return
	}
	if len(s.MostPopularTweets) < 20 {
		s.MostPopularTweets = append(s.MostPopularTweets, tweet)
	}
}

// AnalyzeTweet Analyze a single Tweet for statistics
func (s *TweetStats) AnalyzeTweet(tweet twitterGo.Tweet) {
	if tweet.RetweetedStatus == nil {
		s.updateTopTweets(tweet)
	}
}
