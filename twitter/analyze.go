package twitter

import (
	"sort"

	twitterGo "github.com/dghubble/go-twitter/twitter"
)

const favoriteRetweetRatio = 9
const TopHashtagsLen = 10
const MostPopularTweetsLen = 20

type freq struct {
	value string
	f     int
}

// TweetStats struct returned for twitter statistic anlytics
type TweetStats struct {
	MostPopularTweets []twitterGo.Tweet
	TopHashtags       []freq
	hashtags          map[string]int
}

func tweetPopularityScore(tweet twitterGo.Tweet) int {
	return tweet.FavoriteCount + tweet.RetweetCount*9
}

// NewTweetStats Create a New TweetStats struct
func NewTweetStats() *TweetStats {
	return &TweetStats{
		MostPopularTweets: make([]twitterGo.Tweet, MostPopularTweetsLen),
		hashtags:          make(map[string]int),
		TopHashtags:       make([]freq, TopHashtagsLen),
	}
}

// AnalyzeUserTweets return a TweetStats structure of the analyzed tweets
func AnalyzeUserTweets(tweets []twitterGo.Tweet) *TweetStats {
	s := NewTweetStats()
	for _, tweet := range tweets {
		s.AnalyzeTweet(tweet)
	}
	for hashtag, seen := range s.hashtags {
		s.updateTopHashtags(freq{value: hashtag, f: seen})
	}

	return s
}

func (s *TweetStats) updateTopHashtags(hashtag freq) {
	index := sort.Search(TopHashtagsLen, func(i int) bool {
		return s.TopHashtags[i].f < hashtag.f
	})
	if index < TopHashtagsLen {
		copy(s.TopHashtags[index+1:], s.TopHashtags[index:TopHashtagsLen-1])
		s.TopHashtags[index] = hashtag
	}
}

func (s *TweetStats) updateMostPopularTweets(tweet twitterGo.Tweet) {
	score := tweetPopularityScore(tweet)
	index := sort.Search(len(s.MostPopularTweets), func(i int) bool {
		return tweetPopularityScore(s.MostPopularTweets[i]) < score
	})

	if index < MostPopularTweetsLen {
		copy(s.MostPopularTweets[index+1:], s.MostPopularTweets[index:MostPopularTweetsLen-1])
		s.MostPopularTweets[index] = tweet
	}
}

// AnalyzeTweet Analyze a single Tweet for statistics
func (s *TweetStats) AnalyzeTweet(tweet twitterGo.Tweet) {
	if tweet.RetweetedStatus == nil {
		s.updateMostPopularTweets(tweet)
	}
	for _, hashtag := range tweet.Entities.Hashtags {
		s.hashtags[hashtag.Text]++
	}
}
