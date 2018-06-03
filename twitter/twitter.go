package twitter

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	twitterGo "github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

// Client wrapper for twitter
type Client struct {
	client *twitterGo.Client
}

type Limit struct {
	Remaining int
	Limit     int
	Reset     int64
}

var credentials *clientcredentials.Config
var accessToken string

func init() {
	accessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	credentials = newCredentials()
}

func (l *Limit) TimeLeft() time.Duration {
	reset := time.Unix(l.Reset, 0)
	return reset.Sub(time.Now())
}

func (l *Limit) Delay() time.Duration {
	return time.Duration(float64(l.TimeLeft()) / (float64(l.Remaining) + 1))
}

func GetLimits(resp *http.Response) (*Limit, error) {
	remainingStr := resp.Header.Get("x-rate-limit-remaining")
	remaining, err := strconv.Atoi(remainingStr)
	if err != nil {
		return nil, err
	}

	limitStr := resp.Header.Get("x-rate-limit-limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, err
	}

	resetStr := resp.Header.Get("x-rate-limit-reset")
	reset, err := strconv.ParseInt(resetStr, 10, 64)
	if err != nil {
		return nil, err
	}

	return &Limit{
		Limit:     limit,
		Remaining: remaining,
		Reset:     reset,
	}, nil
}

func newCredentials() *clientcredentials.Config {
	return &clientcredentials.Config{
		ClientID:     os.Getenv("TWITTER_CONSUMER_KEY"),
		ClientSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
}

// NewAppClient create a new Client struct with app credentials
func NewAppClient(ctx context.Context) *Client {
	config := newCredentials()
	httpClient := config.Client(ctx)

	return &Client{
		client: twitterGo.NewClient(httpClient),
	}
}

// GetUserShow get twitter user info by screenName
func (c *Client) GetUserShow(screenName string) (*twitterGo.User, error) {
	userShowParams := &twitterGo.UserShowParams{ScreenName: screenName}
	user, _, err := c.client.Users.Show(userShowParams)
	return user, err
}

// GetUserTweetsStats return statistics about user tweets
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
