package twitter

import (
	"context"
	"log"
	"os"

	twitterGo "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"golang.org/x/oauth2/clientcredentials"
)

// Client wrapper for twitter
type Client struct {
	client *twitterGo.Client
}

var appCredentilsConfig *clientcredentials.Config
var userCredentialsConfig *oauth1.Config
var accessToken string
var consumerKey string
var consumerSecret string

func init() {
	accessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("TWITTER_ACCESS_TOKEN is not set\n")
	}
	consumerKey = os.Getenv("TWITTER_CONSUMER_KEY")
	if consumerKey == "" {
		log.Fatal("TWITTER_CONSUMER_KEY is not set\n")
	}
	consumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	if consumerSecret == "" {
		log.Fatal("TWITTER_CONSUMER_SECRET is not set\n")
	}

	appCredentilsConfig = &clientcredentials.Config{
		ClientID:     consumerKey,
		ClientSecret: consumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	userCredentialsConfig = oauth1.NewConfig(consumerKey, consumerSecret)
}

// NewAppClient create a new Client struct with app credentials
func NewAppClient(ctx context.Context) *Client {
	httpClient := appCredentilsConfig.Client(ctx)

	return &Client{
		client: twitterGo.NewClient(httpClient),
	}
}

// NewUserClient create a new Client using user credentials
func NewUserClient(ctx context.Context, accessToken string, accessSecret string) *Client {
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := userCredentialsConfig.Client(ctx, token)
	return &Client{
		client: twitterGo.NewClient(httpClient),
	}
}

// GetUserTimeline return user timeline
func (c *Client) GetUserTimeline(screenName string, after int64) ([]twitterGo.Tweet, *Limit, error) {
	params := &twitterGo.UserTimelineParams{
		ScreenName:      screenName,
		Count:           200,
		IncludeRetweets: twitterGo.Bool(true),
	}
	if after != 0 {
		params.MaxID = after
	}
	tweets, resp, err := c.client.Timelines.UserTimeline(params)
	if err != nil {
		return nil, nil, err
	}
	limits, err := GetLimits(resp)
	return tweets, limits, err
}
