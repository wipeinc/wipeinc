package model

import (
	"time"

	"github.com/ChimeraCoder/anaconda"
)

// User model
type User struct {
	ID              int64
	URL             string
	Name            string
	ScreenName      string
	Location        string
	Lang            string
	Description     string
	BackgroundImage string
	Image           string
	FavouritesCount int
	FollowersCount  int
	FriendsCount    int

	UpdatedAt time.Time
}

// NewUser create a new user from anaconda twitter struct
func NewUser(user anaconda.User) (*User, error) {
	createdAt, err := time.Parse(time.RubyDate, user.CreatedAt)
	if err != nil {
		return nil, err
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

		FavouritesCount: user.FavouritesCount,
		FollowersCount:  user.FollowersCount,
		FriendsCount:    user.FriendsCount,
		UpdatedAt:       createdAt,
	}, nil
}
