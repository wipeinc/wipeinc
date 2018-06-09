package twitter

import (
	"time"

	twitterGo "github.com/dghubble/go-twitter/twitter"
	"github.com/wipeinc/wipeinc/model"
)

// GetUserShow get twitter user info by screenName
func (c *Client) GetUserShow(screenName string) (*model.User, error) {
	userShowParams := &twitterGo.UserShowParams{ScreenName: screenName}
	rawUser, _, err := c.client.Users.Show(userShowParams)
	user, err := NewUser(rawUser)
	if err != nil {
		return nil, err
	}
	return user, err
}

// UserLookup lookup for users using IDs
func (c *Client) UserLookup(userID []int64) ([]model.User, *Limit, error) {
	params := &twitterGo.UserLookupParams{
		UserID: userID,
	}
	rawUsers, resp, err := c.client.Users.Lookup(params)
	if err != nil {
		return nil, nil, err
	}
	limits, err := GetLimits(resp)
	users := make([]model.User, 0, len(rawUsers))
	for _, rawUser := range rawUsers {
		user, err := NewUser(&rawUser)
		if err != nil {
			return nil, limits, err
		}
		users = append(users, *user)
	}
	return users, limits, err
}

// NewUser create a new user from anaconda twitter struct
func NewUser(user *twitterGo.User) (*model.User, error) {
	createdAt, err := time.Parse(time.RubyDate, user.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt := time.Now().Round(time.Microsecond).UTC()
	return &model.User{
		CreatedAt:            createdAt,
		Description:          user.Description,
		FavouritesCount:      user.FavouritesCount,
		FollowersCount:       user.FollowersCount,
		FriendsCount:         user.FriendsCount,
		Name:                 user.Name,
		ProfileBannerURL:     user.ProfileBannerURL,
		ProfileImageURLHTTPS: user.ProfileImageURLHttps,
		StatusesCount:        user.StatusesCount,
		UpdatedAt:            updatedAt,
		URL:                  user.URL,
	}, nil
}
