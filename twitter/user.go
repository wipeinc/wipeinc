package twitter

import (
	"time"

	twitterGo "github.com/dghubble/go-twitter/twitter"
	"github.com/wipeinc/wipeinc/entity"
)

func parseDate(date string) (time.Time, error) {
	return time.Parse(time.RubyDate, date)
}

// GetUserShow get twitter user info by screenName
func (c *Client) GetUserShow(screenName string) (*entity.User, error) {
	userShowParams := &twitterGo.UserShowParams{ScreenName: screenName}
	rawUser, _, err := c.client.Users.Show(userShowParams)
	user, err := NewUser(rawUser)
	if err != nil {
		return nil, err
	}
	return user, err
}

// UserLookup lookup for users using IDs
func (c *Client) UserLookup(userID []int64) ([]entity.User, *Limit, error) {
	params := &twitterGo.UserLookupParams{
		UserID: userID,
	}
	rawUsers, resp, err := c.client.Users.Lookup(params)
	if err != nil {
		return nil, nil, err
	}
	limits, err := GetLimits(resp)
	users := make([]entity.User, 0, len(rawUsers))
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
func NewUser(user *twitterGo.User) (*entity.User, error) {
	createdAt, err := parseDate(user.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt := time.Now().Round(time.Microsecond).UTC()
	return &entity.User{
		ID:                   user.ID,
		CreatedAt:            createdAt,
		Description:          user.Description,
		FavouritesCount:      user.FavouritesCount,
		FollowersCount:       user.FollowersCount,
		FriendsCount:         user.FriendsCount,
		Name:                 user.Name,
		ProfileBannerURL:     user.ProfileBannerURL,
		ProfileImageURLHTTPS: user.ProfileImageURLHttps,
		ScreenName:           user.ScreenName,
		StatusesCount:        user.StatusesCount,
		UpdatedAt:            updatedAt,
		URL:                  user.URL,
	}, nil
}
