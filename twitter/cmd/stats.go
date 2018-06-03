package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	twitterGo "github.com/dghubble/go-twitter/twitter"
	"github.com/olekukonko/tablewriter"
	"github.com/wipeinc/wipeinc/twitter"
)

func printTop(name string, top []twitter.Freq) {
	data := make([][]string, 0, len(top))
	for _, freq := range top {
		data = append(data, []string{freq.Value, strconv.Itoa(freq.F)})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{name, "occ"})
	table.AppendBulk(data)
	table.Render()

}

func printTopTweets(tweets []twitterGo.Tweet) {
	data := make([][]string, 0, len(tweets))
	for _, tweet := range tweets {
		data = append(data, []string{
			tweet.Text,
			strconv.Itoa(tweet.FavoriteCount),
			strconv.Itoa(tweet.RetweetCount)})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"tweet", "fav", "rt"})
	table.AppendBulk(data)
	table.Render()

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: cmd <screen_name>")
	}
	user := os.Args[1]
	client := twitter.NewAppClient(context.Background())
	var since int64
	since = 0
	stats := twitter.NewTweetStats()
	for i := 0; i < 4; i++ {
		fmt.Printf("[%d/4]fetching user since: %d\n", i+1, since)
		tweets, limit, err := client.GetUserTimeline(user, since)
		if err != nil {
			fmt.Println("error")
			fmt.Printf("%+v\n", err)
			break
		}
		stats.AnalyzeTweets(tweets)
		if len(tweets) < 199 {
			fmt.Printf("got %d tweets\n", len(tweets))
			break
		}
		since = tweets[len(tweets)-1].ID
		limit.Delay()
	}
	top := stats.TopHashtags(0)
	printTop("hashtag", top)
	top = stats.TopRetweets(0)
	printTop("retweeted", top)
	top = stats.TopMentions(0)
	printTop("mention", top)
	top = stats.TopDomains(0)
	printTop("domain", top)
	printTopTweets(stats.MostPopularTweets[:5])
}
