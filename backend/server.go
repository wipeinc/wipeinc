package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	twitterLogin "github.com/dghubble/gologin/twitter"
	"github.com/dghubble/oauth1"
	oauth1Login "github.com/dghubble/oauth1"
	twitterOAuth1 "github.com/dghubble/oauth1/twitter"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/wipeinc/wipeinc/db"
	"github.com/wipeinc/wipeinc/model"
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
	sessionName               = "wipeinc"
	sessionUserKey            = "twitterID"
	userAccessTokenKey        = "twitterAccessToken"
	userAcccessTokenSecretKey = "twitterAccessTokenSecret"
)

var sessionStore sessions.Store
var sessionSecret string
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
	mux.HandleFunc("/api/myprofile", showMyProfile)
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
		accessToken, acessSecret, err := oauth1Login.AccessTokenFromContext(ctx)
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
		session.Values[userAcccessTokenSecretKey] = accessTokenSecret
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

// requireLogin redirects unauthenticated users to the login route.
func requireLogin(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if !isAuthenticated(req) {
			http.Redirect(w, req, "/", http.StatusFound)
			return
		}
		next.ServeHTTP(w, req)
	}
	return http.HandlerFunc(fn)
}

// isAuthenticated returns true if the user has a signed session cookie.
func isAuthenticated(req *http.Request) bool {
	if _, err := sessionStore.Get(req, sessionName); err == nil {
		return true
	}
	return false
}

// showIndex show empty page with js scripts
func showIndex(w http.ResponseWriter, r *http.Request) {
	index, err := Asset("static/index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	indexReader := bytes.NewBuffer(index)
	io.Copy(w, indexReader)
}

func getUserTwitterClient(req *http.Request) (*twitter.Client, error) {
	session, err := sesessionStore.Get(req, sessionName)
	if err != nil {
		return nil, err
	}
	accessToken := session.Values[userAccessTokenKey]
	accessTokenSecret := session.Values[userAccessTokenSecretKey]
	ctx := req.Context()
	client := twitter.NewUserClient(ctx, accessToken, accessTokenSecret)
}

// showProfile route for /api/profile/{screenName}
func showProfile(w http.ResponseWriter, req *http.Request) {
	var err error
	var user *model.User

	params := mux.Vars(req)
	screenName := params["name"]

	user, err = db.DB.GetUser(screenName)
	if err != nil {
		client, err := getUserTwitterClint(req)
		if err != nil {
			http.Error(w, "unauthroized", http.StatusUnauthorized)
			return
		}
		fetchedUser, err := client.GetUserShow(screenName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = db.DB.AddUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	json.NewEncoder(w).Encode(user)
}
