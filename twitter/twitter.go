package twitter

import (
	"context"
	"os"

	twitterGo "github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

type Client struct {
	client *twitterGo.Client
}

var AppClient *Client
var credentials *clientcredentials.Config

var accessToken string
var token string

func init() {
	accessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	credentials = newCredentials()
}

func newCredentials() *clientcredentials.Config {
	return &clientcredentials.Config{
		ClientID:     os.Getenv("TWITTER_CONSUMER_KEY"),
		ClientSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
}

func NewAppClient(ctx context.Context) *Client {
	config := newCredentials()
	httpClient := config.Client(ctx)

	// Twitter client
	return &Client{
		client: twitterGo.NewClient(httpClient),
	}
}

func (c *Client) GetUserShow(screenName string) (*twitterGo.User, error) {
	userShowParams := &twitterGo.UserShowParams{ScreenName: screenName}
	user, _, err := c.client.Users.Show(userShowParams)
	return user, err
}
