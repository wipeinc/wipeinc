package twitter

import (
	"sort"

	twitterGo "github.com/dghubble/go-twitter/twitter"
)

const favoriteRetweetRatio = 9
const topHashtagsLen = 10
const topMentionsLen = 10
const mostPopularTweetsLen = 20

// HashtagFreq is the structure for top hashtags
type HashtagFreq struct {
	Value string
	F     int
}

// MentionFreq is the structure for top hashtags
type MentionFreq struct {
	Value twitterGo.MentionEntity
	F     int
}

// TweetStats struct returned for twitter statistic anlytics
type TweetStats struct {
	MostPopularTweets []twitterGo.Tweet
	mentionsCount     map[int64]int
	mentions          map[int64]twitterGo.MentionEntity
	hashtags          map[string]int
}

func tweetPopularityScore(tweet twitterGo.Tweet) int {
	return tweet.FavoriteCount + tweet.RetweetCount*favoriteRetweetRatio
}

// NewTweetStats Create a New TweetStats struct
func NewTweetStats() *TweetStats {
	return &TweetStats{
		MostPopularTweets: make([]twitterGo.Tweet, mostPopularTweetsLen),
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

// TopHashtags Return top len hashtags
func (s *TweetStats) TopHashtags(len int) []HashtagFreq {
	if len == 0 {
		len = topHashtagsLen
	}
	top := make([]HashtagFreq, len)
	for hashtag, seen := range s.hashtags {
		insert := HashtagFreq{Value: hashtag, F: seen}
		index := sort.Search(len, func(i int) bool {
			return top[i].F < insert.F
		})
		if index < len {
			copy(top[index+1:], top[index:len-1])
			top[index] = insert
		}
	}
	return top
}

// TopMentions return the most len user mentionned
func (s *TweetStats) TopMentions(len int) []MentionFreq {
	if len == 0 {
		len = topMentionsLen
	}
	top := make([]MentionFreq, len)
	for ID, mention := range s.mentions {
		insert := MentionFreq{Value: mention, F: s.mentionsCount[ID]}
		index := sort.Search(len, func(i int) bool {
			return top[i].F < insert.F
		})
		if index < len {
			copy(top[index+1:], top[index:len-1])
			top[index] = insert
		}
	}
	return top
}

func (s *TweetStats) updateMostPopularTweets(tweet twitterGo.Tweet) {
	score := tweetPopularityScore(tweet)
	index := sort.Search(len(s.MostPopularTweets), func(i int) bool {
		return tweetPopularityScore(s.MostPopularTweets[i]) < score
	})

	if index < mostPopularTweetsLen {
		copy(s.MostPopularTweets[index+1:], s.MostPopularTweets[index:mostPopularTweetsLen-1])
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
