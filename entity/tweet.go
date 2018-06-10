package entity

import (
	"strconv"
	"time"
)

// Tweet entity
type Tweet struct {
	ID              int64     `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	FavoriteCount   int       `json:"favorite_count"`
	FullText        string    `json:"full_text"`
	Lang            string    `json:"lang"`
	Links           []string  `json:"links"`
	Hashtags        []string  `json:"hashtags"`
	QuotedStatus    *Tweet    `json:"quote_status"`
	UserMentions    []*User   `json:"mentions"`
	RetweetCount    int       `json:"retweet_count"`
	RetweetedStatus *Tweet    `json:"retweeted_status"`
	Text            string    `json:"text"`
	Truncated       bool      `json:"truncated"`
	URLS            []string  `json:"urls"`
	User            *User     `json:"user"`
}

const favoriteRetweetRatio = 9

// IDStr is string version of the User ID
func (t *Tweet) IDStr() string {
	return strconv.FormatInt(t.ID, 10)
}

// FullTweet return always the most complete text
func (t *Tweet) FullTweet() string {
	if t.Truncated {
		return t.FullText
	}
	return t.Text
}

// PopularityScore aproximate the popularity of a tweet in fuction of
// favorites and retweets counts
func (t *Tweet) PopularityScore() int {
	return t.FavoriteCount + t.RetweetCount*favoriteRetweetRatio
}
