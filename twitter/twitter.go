package twitter

import (
	"net/url"

	"github.com/wow-sweetlie/anaconda"
)

// Client Twitter API wrapper
type Client struct {
	API *anaconda.TwitterApi
}

func (c *Client) BlockUserFollowers(screenName string) error {
	friends, err := c.GetFriendsIds()
	if err != nil {
		return err
	}
	v := url.Values{}
	v.Set("screen_name", screenName)
	followers, err := c.GetFollowersIds(v)
	if err != nil {
		return err
	}
	followersToBan := MinusIDList(followers, friends)
	c.BlockUserIds(followersToBan)
	return nil
}

func (c *Client) BlockUserIds(ids []int64) error {
	for _, id := range ids {
		_, err := c.API.BlockUserId(id, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetFriendsIds return a slice of friends ids
func (c *Client) GetFriendsIds() ([]int64, error) {
	ids := []int64{}
	pages := c.API.GetFriendsIdsAll(nil)
	for page := range pages {
		if page.Error != nil {
			return nil, page.Error
		}
		ids = append(ids, page.Ids...)
	}
	return ids, nil
}

// GetFollowersIds return a slice of followers ids
func (c *Client) GetFollowersIds(v url.Values) ([]int64, error) {
	ids := []int64{}
	pages := c.API.GetFollowersIdsAll(v)
	for page := range pages {
		if page.Error != nil {
			return nil, page.Error
		}
		ids = append(ids, page.Ids...)
	}
	return ids, nil
}

// NewClient return a new client based on token credentials
func NewClient(accessToken string, accessTokenSecret string) *Client {
	return &Client{
		API: anaconda.NewTwitterApi(accessToken, accessTokenSecret),
	}
}
