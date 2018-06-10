package entity

import (
	"time"
)

// Tweet entity
type Tweet struct {
	ID              int64     `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	FavoriteCount   int       `json:"favorite_count"`
	FullText        string    `json:"full_text"`
	Links           []string  `json:"links"`
	Hashtags        []string  `json:"hashtags"`
	UserMentions    []*User   `json:"mentions"`
	RetweetCount    int       `json:"retweet_count"`
	RetweetedStatus *Tweet    `json:"retweeted_status"`
	Text            string    `json:"text"`
	URLS            []string  `json:"urls"`
	User            *User     `json:"user"`
}
