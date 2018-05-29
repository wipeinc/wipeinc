package model

import (
	"time"

	twitterGo "github.com/dghubble/go-twitter/twitter"
)

// User model
type User struct {
	TwitterUser *twitterGo.User `json:"twitterUser"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}

// NewUser create a new user from anaconda twitter struct
func NewUser(user *twitterGo.User) *User {
	// createdAt, err := time.Parse(time.RubyDate, user.CreatedAt)
	updatedAt := time.Now().Round(time.Microsecond).UTC()
	return &User{
		TwitterUser: user,
		UpdatedAt:   updatedAt,
	}
}
