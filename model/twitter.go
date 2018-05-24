package model

import (
	"time"

	"github.com/wow-sweetlie/anaconda"
)

// User model
type User struct {
	ID              int64  `json:"id"`
	URL             string `json:"url"`
	Name            string `json:"name"`
	ScreenName      string `json:"screenName"`
	Location        string `json:"location,omitempty"`
	Lang            string `json:"lang,omitempty"`
	Description     string `json:"description,omitempty"`
	BackgroundImage string `json:"backgroundImage,omitempty"`
	Image           string `json:"image,omitempty"`
	FavouritesCount int    `json:"favorites"`
	FollowersCount  int    `json:"followers"`
	FriendsCount    int    `json:"friends"`

	UpdatedAt time.Time `json:"updatedAt"`
}

// NewUser create a new user from anaconda twitter struct
func NewUser(user anaconda.User) (*User, error) {
	// createdAt, err := time.Parse(time.RubyDate, user.CreatedAt)
	updatedAt := time.Now().Round(time.Microsecond).UTC()
	// if err != nil {
	// 	return nil, err
	// }
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

		FavouritesCount: user.FavouritesCount,
		FollowersCount:  user.FollowersCount,
		FriendsCount:    user.FriendsCount,
		UpdatedAt:       updatedAt,
	}, nil
}
