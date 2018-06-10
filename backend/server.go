package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	oauth1Login "github.com/dghubble/gologin/oauth1"
	twitterLogin "github.com/dghubble/gologin/twitter"
	"github.com/dghubble/oauth1"
	twitterOAuth1 "github.com/dghubble/oauth1/twitter"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/wipeinc/wipeinc/model"
	"github.com/wipeinc/wipeinc/repository"
	"github.com/wipeinc/wipeinc/twitter"
)

// Config struct for backend server
type Config struct {
	TwitterConsumerKey    string
	TwitterConsumerSecret string
	TwitterCallbackURL    string
	CookieSecretToken     string
}

const (
	sessionName              = "wipeinc"
	sessionUserKey           = "twitterID"
	userAccessTokenKey       = "twitterAccessToken"
	userAccessTokenSecretKey = "twitterAccessTokenSecret"
)

var sessionStore sessions.Store
var config *Config

func init() {
	config = &Config{
		TwitterConsumerKey:    os.Getenv("TWITTER_CONSUMER_KEY"),
		TwitterConsumerSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
		TwitterCallbackURL:    os.Getenv("TWITTER_CALLBACK_URL"),
	}
	if config.TwitterConsumerKey == "" {
		log.Fatal("twitter consumer key not set")
	}
	if config.TwitterConsumerSecret == "" {
		log.Fatal("twitter consumer secret not set")
	}
	if _, err := url.Parse(config.TwitterCallbackURL); err != nil {
		log.Fatalf("invalid twitter callback url : %s", err.Error())
	}
	sessionSecret := os.Getenv("SESSION_SECRET_KEY")
	if sessionSecret == "" {
		log.Fatal("session secret not set")
	}
	if len(sessionSecret) != 32 && len(sessionSecret) != 64 {
		log.Fatalf("invalid session secret size: %d\n", len(sessionSecret))
	}
	sessionStore = sessions.NewCookieStore([]byte(sessionSecret), nil)

}

func main() {
	oauth1Config := &oauth1.Config{
		ConsumerKey:    config.TwitterConsumerKey,
		ConsumerSecret: config.TwitterConsumerSecret,
		CallbackURL:    config.TwitterCallbackURL,
		Endpoint:       twitterOAuth1.AuthorizeEndpoint,
	}
	mux := mux.NewRouter()
	mux.Handle("/twitter/login", twitterLogin.LoginHandler(oauth1Config, nil))
	mux.Handle("/twitter/callback", twitterLogin.CallbackHandler(oauth1Config, issueSession(), nil))
	mux.HandleFunc("/api/profile/{name}", showProfile)
	mux.HandleFunc("/api/profile/{name}/analyze", showProfile)
	mux.HandleFunc("/api/sessions/logout", logoutHandler)
	mux.PathPrefix("/").HandlerFunc(showIndex)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// issueSession issues a cookie session after successful Twitter login
func issueSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		twitterUser, err := twitterLogin.UserFromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		accessToken, accessTokenSecret, err := oauth1Login.AccessTokenFromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session, err := sessionStore.New(req, sessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values[sessionUserKey] = twitterUser.ID
		session.Values[userAccessTokenKey] = accessToken
		session.Values[userAccessTokenSecretKey] = accessTokenSecret
		err = session.Save(req, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, req, "/", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

// logoutHandler destroys the session on POSTs and redirects to home.
func logoutHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		session, err := sessionStore.Get(req, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		delete(session.Values, sessionUserKey)
		session.Options.MaxAge = -1
		err = session.Save(req, w)
		if err != nil {
			data, errJSON := json.Marshal(session)
			if errJSON != nil {
				log.Printf("could not delete session %s\n", data)
			}
			log.Printf("could not delete session err: %s\n", err)
		}
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

// showIndex show empty page with js scripts
func showIndex(w http.ResponseWriter, r *http.Request) {
	index, err := Asset("static/index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	indexReader := bytes.NewBuffer(index)
	if _, err := io.Copy(w, indexReader); err != nil {
		log.Printf("error writing index to http connection: %s\n", err)
	}
}

func getUserTwitterClient(req *http.Request) (*twitter.Client, error) {
	session, err := sessionStore.Get(req, sessionName)
	if err != nil {
		return nil, err
	}
	var accessToken string
	var accessTokenSecret string
	var ok bool

	val := session.Values[userAccessTokenKey]
	if accessToken, ok = val.(string); !ok {
		return nil, errors.New("user token key absent")
	}
	val = session.Values[userAccessTokenSecretKey]
	if accessTokenSecret, ok = val.(string); !ok {
		return nil, errors.New("user token key secret absent")
	}
	ctx := req.Context()
	return twitter.NewUserClient(ctx, accessToken, accessTokenSecret), nil
}

// showProfile route for /api/profile/{screenName}
func showProfile(w http.ResponseWriter, req *http.Request) {
	var err error
	var user *model.User

	params := mux.Vars(req)
	screenName := params["name"]

	user, err = repository.DB.GetUserByScreenName(screenName)
	if err == nil {
		log.Println("cache hit")
	}
	if err != nil {
		var client *twitter.Client
		client, err = getUserTwitterClient(req)
		if err != nil {
			http.Error(w, "unauthroized", http.StatusUnauthorized)
			return
		}
		user, err = client.GetUserShow(screenName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		go saveUser(*user)
	}

	if err = json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("error trying to serialize twitter user profile: %s", err.Error())
	}
}

// analyzeProfile route for /api/profile/{screenName}/analyze
func analyzeProfile(w http.ResponseWriter, req *http.Request) {
	client, err := getUserTwitterClient(req)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
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

}

func saveUser(user model.User) {
	err := repository.DB.AddUser(&user)
	if err != nil {
		log.Printf("Error callling AddUser: %s", err.Error())
	}

}
