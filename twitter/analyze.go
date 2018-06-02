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
	mentionsCount     map[int64]int
	mentions          map[int64]twitterGo.MentionEntity
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
		mentions:          make(map[int64]twitterGo.MentionEntity),
		mentionsCount:     make(map[int64]int),
	}
}

// AnalyzeUserTweets return a TweetStats structure of the analyzed tweets
func (s *TweetStats) AnalyzeUserTweets(tweets []twitterGo.Tweet) {
	for _, tweet := range tweets {
		s.AnalyzeTweet(tweet)
	}
}

// Return top len hashtags
func (s *TweetStats) TopHashtags(len int) []freq {
	if len == 0 {
		len = TopHashtagsLen
	}
	topHashtags := make([]freq, len)
	for hashtag, seen := range s.hashtags {
		topHashtags = updateTop(topHashtags, freq{value: hashtag, f: seen})
	}
	return topHashtags
}

func updateTop(top []freq, insert freq) []freq {
	index := sort.Search(len(top), func(i int) bool {
		return top[i].f < insert.f
	})
	if index < len(top) {
		copy(top[index+1:], top[index:TopHashtagsLen-1])
		top[index] = insert
	}
	return top
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
	for _, mention := range tweet.Entities.UserMentions {
		s.mentions[mention.ID] = mention
		s.mentionsCount[mention.ID]++
	}
}
