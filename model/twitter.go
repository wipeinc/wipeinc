package model

import (
	"time"
)

// User model
type User struct {
	ID                   int64     `json:"id"`
	CreatedAt            time.Time `json:"created_at"`
	Description          string    `json:"description"`
	FavouritesCount      int       `json:"favourites_count"`
	FollowersCount       int       `json:"followers_count"`
	FriendsCount         int       `json:"friends_count"`
	Name                 string    `json:"name"`
	ProfileBannerURL     string    `json:"profile_banner_url"`
	ProfileImageURLHTTPS string    `json:"profile_image_url_https"`
	ScreenName           string    `json:"screen_name"`
	StatusesCount        int       `json:"statuses_count"`
	UpdatedAt            time.Time `json:"updatedAt"`
	URL                  string    `json:"url"`
}
