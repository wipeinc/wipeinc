package twitter

import (
	"net/url"
	"sort"

	"log"

	twitterGo "github.com/dghubble/go-twitter/twitter"
)

const favoriteRetweetRatio = 9
const topHashtagsLen = 10
const topMentionsLen = 10
const topDomains = 10
const mostPopularTweetsLen = 20

// Freq is the structure for sorting
type Freq struct {
	Value string
	F     int
}

// TweetStats struct returned for twitter statistic anlytics
type TweetStats struct {
	MostPopularTweets []twitterGo.Tweet
	mentionsCount     map[string]int
	domainsCount      map[string]int
	hashtagsCount     map[string]int
}

var blacklisted = struct{}{}
var blacklistedDomains = map[string]struct{}{
	"bit.ly":      blacklisted,
	"twitter.com": blacklisted,
}

func isBlacklisted(domain string) bool {
	_, c := blacklistedDomains[domain]
	return c
}

func tweetPopularityScore(tweet twitterGo.Tweet) int {
	return tweet.FavoriteCount + tweet.RetweetCount*favoriteRetweetRatio
}

// NewTweetStats Create a New TweetStats struct
func NewTweetStats() *TweetStats {
	return &TweetStats{
		MostPopularTweets: make([]twitterGo.Tweet, mostPopularTweetsLen),
		hashtagsCount:     make(map[string]int),
		mentionsCount:     make(map[string]int),
		domainsCount:      make(map[string]int),
	}
}

// AnalyzeTweets return a TweetStats structure of the analyzed tweets
func (s *TweetStats) AnalyzeTweets(tweets []twitterGo.Tweet) {
	for _, tweet := range tweets {
		s.AnalyzeTweet(tweet)
	}
}

// TopDomains Return top len hashtags
func (s *TweetStats) TopDomains(len int) []Freq {
	if len == 0 {
		len = topDomains
	}
	return top(s.domainsCount, len)
}

// TopHashtags Return top len hashtags
func (s *TweetStats) TopHashtags(len int) []Freq {
	if len == 0 {
		len = topHashtagsLen
	}
	return top(s.hashtagsCount, len)
}

// TopMentions return the most len user mentionned
func (s *TweetStats) TopMentions(len int) []Freq {
	if len == 0 {
		len = topMentionsLen
	}
	return top(s.mentionsCount, len)
}

func top(elements map[string]int, len int) []Freq {
	top := make([]Freq, len)
	for value, count := range elements {
		insert := Freq{Value: value, F: count}
		index := sort.Search(len, func(i int) bool {
			return top[i].F < insert.F
		})
		if index < len {
			copy(top[index+1:], top[index:len-1])
			top[index] = insert
		}
	}
	index := sort.Search(len, func(i int) bool {
		return top[i].Value == ""
	})
	if (index + 1) < len {
		return top[:index]
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
	if tweet.Entities != nil {
		for _, hashtag := range tweet.Entities.Hashtags {
			s.hashtagsCount[hashtag.Text]++
		}
		for _, mention := range tweet.Entities.UserMentions {
			s.mentionsCount[mention.IDStr]++
		}

		for _, urlEntity := range tweet.Entities.Urls {
			u, err := url.Parse(urlEntity.ExpandedURL)
			if err == nil {
				if !isBlacklisted(u.Hostname()) {
					s.domainsCount[u.Hostname()]++
				}
			} else {
				log.Printf("failed to parse url: %s\n", urlEntity.ExpandedURL)
			}
		}
	}
}
