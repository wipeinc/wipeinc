package twitter

import (
	"net/url"
	"sort"

	"log"

	"github.com/wipeinc/wipeinc/entity"
)

const favoriteRetweetRatio = 9
const topHashtagsLen = 10
const topMentionsLen = 10
const topDomainsLen = 10
const topRetweetsLen = 10

// MostPopularTweetsLen is default Maximum size for most popular tweets
const MostPopularTweetsLen = 5

// Freq is the structure for sorting
type Freq struct {
	Value interface{}
	F     int
}

// TweetStats struct returned for twitter statistic anlytics
type TweetStats struct {
	MostPopularTweets []entity.Tweet
	mentionsCount     map[string]int
	retweetsCount     map[string]int
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

// NewTweetStats Create a New TweetStats struct
func NewTweetStats() *TweetStats {
	return &TweetStats{
		MostPopularTweets: make([]entity.Tweet, MostPopularTweetsLen),
		hashtagsCount:     make(map[string]int),
		mentionsCount:     make(map[string]int),
		domainsCount:      make(map[string]int),
		retweetsCount:     make(map[string]int),
	}
}

// AnalyzeTweets return a TweetStats structure of the analyzed tweets
func (s *TweetStats) AnalyzeTweets(tweets []entity.Tweet) {
	for _, tweet := range tweets {
		s.AnalyzeTweet(tweet)
	}
}

// TopDomains Return top len hashtags
func (s *TweetStats) TopDomains(len int) []Freq {
	if len == 0 {
		len = topDomainsLen
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

// TopRetweets return the most len user retweeted
func (s *TweetStats) TopRetweets(len int) []Freq {
	if len == 0 {
		len = topRetweetsLen
	}
	return top(s.retweetsCount, len)
}

// Min return the Min between to integer
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func top(elements map[string]int, maxLen int) []Freq {
	len := Min(len(elements), maxLen)
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

func (s *TweetStats) updateMostPopularTweets(tweet entity.Tweet) {
	score := tweet.PopularityScore()
	index := sort.Search(len(s.MostPopularTweets), func(i int) bool {
		return s.MostPopularTweets[i].PopularityScore() < score
	})

	if index < MostPopularTweetsLen {
		copy(s.MostPopularTweets[index+1:], s.MostPopularTweets[index:MostPopularTweetsLen-1])
		s.MostPopularTweets[index] = tweet
	}
}

// AnalyzeTweet Analyze a single Tweet for statistics
func (s *TweetStats) AnalyzeTweet(tweet entity.Tweet) {
	if tweet.RetweetedStatus == nil {
		s.updateMostPopularTweets(tweet)
	} else if tweet.User.ID != tweet.RetweetedStatus.User.ID {
		s.retweetsCount[tweet.RetweetedStatus.User.IDStr()]++
	}
	for _, hashtag := range tweet.Hashtags {
		s.hashtagsCount[hashtag]++
	}
	for _, mention := range tweet.UserMentions {
		if mention.ID != tweet.User.ID {
			s.mentionsCount[mention.IDStr()]++
		}
	}

	for _, sURL := range tweet.URLS {
		u, err := url.Parse(sURL)
		if err == nil {
			if !isBlacklisted(u.Hostname()) {
				s.domainsCount[u.Hostname()]++
			}
		} else {
			log.Printf("failed to parse url: %s\n", sURL)
		}
	}
}
