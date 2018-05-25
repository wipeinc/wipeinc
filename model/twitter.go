package model

import (
	"fmt"
	"time"

	"github.com/wow-sweetlie/anaconda"
)

// User model
type User struct {
	ID              int64     `json:"id"`
	URL             string    `json:"url"`
	Name            string    `json:"name"`
	ScreenName      string    `json:"screenName"`
	Location        string    `json:"location,omitempty"`
	Lang            string    `json:"lang,omitempty"`
	Description     string    `json:"description,omitempty"`
	BackgroundImage string    `json:"backgroundImage,omitempty"`
	Image           string    `json:"image,omitempty"`
	Banner          string    `json:"banner,omitempty"`
	StatusesCount   int64     `json:"statuses"`
	FavouritesCount int       `json:"favorites"`
	FollowersCount  int       `json:"followers"`
	FriendsCount    int       `json:"friends"`
	CreatedAt       time.Time `json:"createdAt"`

	UpdatedAt time.Time `json:"updatedAt"`
}

// NewUser create a new user from anaconda twitter struct
func NewUser(user anaconda.User) (*User, error) {
	// createdAt, err := time.Parse(time.RubyDate, user.CreatedAt)
	updatedAt := time.Now().Round(time.Microsecond).UTC()
	createdAt, err := time.Parse(time.RubyDate, user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("can't parse create date %s of user %s, err: %q",
			user.CreatedAt,
			user.ScreenName,
			err)
	}

	return &User{
		ID: user.Id,

		URL: user.URL,

		Name:            user.Name,
		ScreenName:      user.ScreenName,
		Location:        user.Location,
		Lang:            user.Lang,
		Description:     user.Description,
		BackgroundImage: user.ProfileBackgroundImageUrlHttps,
		Image:           user.ProfileImageUrlHttps,
		Banner:          user.ProfileBannerURL,
		StatusesCount:   user.StatusesCount,
		FavouritesCount: user.FavouritesCount,
		FollowersCount:  user.FollowersCount,
		FriendsCount:    user.FriendsCount,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}, nil
}
