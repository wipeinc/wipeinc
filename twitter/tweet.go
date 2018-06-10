package twitter

import (
	twitterGo "github.com/dghubble/go-twitter/twitter"
	"github.com/wipeinc/wipeinc/entity"
)

// NewTweet fill entity tweet structure from a a fetched twitter tweet
func NewTweet(apiTweet twitterGo.Tweet) (*entity.Tweet, error) {
	createdAt, err := apiTweet.CreatedAtTime()
	if err != nil {
		return nil, err
	}
	tweet := &entity.Tweet{
		ID:            apiTweet.ID,
		CreatedAt:     createdAt,
		FavoriteCount: apiTweet.FavoriteCount,
		FullText:      apiTweet.FullText,
		RetweetCount:  apiTweet.RetweetCount,
		Text:          apiTweet.Text,
	}
	if apiTweet.Entities != nil {
		entities := apiTweet.Entities
		tweet.Hashtags = make([]string, 0, len(entities.Hashtags))
		for _, hashtag := range entities.Hashtags {
			tweet.Hashtags = append(tweet.Hashtags, hashtag.Text)
		}
		tweet.UserMentions = make([]*entity.User, 0, len(entities.UserMentions))
		for _, mention := range entities.UserMentions {
			mentionedUser := &entity.User{
				ID:         mention.ID,
				ScreenName: mention.ScreenName,
				Name:       mention.Name,
			}
			tweet.UserMentions = append(tweet.UserMentions, mentionedUser)
		}
		tweet.URLS = make([]string, 0, len(entities.Urls))
		for _, urlEntity := range entities.Urls {
			tweet.URLS = append(tweet.URLS, urlEntity.ExpandedURL)
		}
	}

	tweet.User, err = NewUser(apiTweet.User)
	if err != nil {
		return nil, err
	}

	return tweet, nil
}
