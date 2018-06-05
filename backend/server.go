package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"cloud.google.com/go/profiler"
	twitterLogin "github.com/dghubble/gologin/twitter"
	"github.com/dghubble/oauth1"
	twitterOAuth1 "github.com/dghubble/oauth1/twitter"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/wipeinc/wipeinc/db"
	"github.com/wipeinc/wipeinc/model"
	"github.com/wipeinc/wipeinc/twitter"
	"google.golang.org/appengine"
)

// Config struct for backend server
type Config struct {
	TwitterConsumerKey    string
	TwitterConsumerSecret string
	TwitterCallbackURL    string
	CookieSecretToken     string
}

const (
	sessionName    = "wipeinc"
	sessionUserKey = "twitterID"
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
	err := profiler.Start(profiler.Config{
		Service: "wipeinc",
	})
	if err != nil {
		log.Println(err)
	}
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
	mux.HandleFunc("/api/sessions/logout", logoutHandler)
	mux.PathPrefix("/").HandlerFunc(showIndex)
	http.Handle("/", mux)
	appengine.Main()
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
		session, err := sessionStore.New(req, sessionName)
		if err != nil {
			log.Printf("Error creating session: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		session.Values[sessionUserKey] = twitterUser.ID
		err = session.Save(req, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.Redirect(w, req, "/profile", http.StatusFound)
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

// showProfile route for /api/profile/{screenName}
func showProfile(w http.ResponseWriter, r *http.Request) {
	if appengine.IsDevAppServer() {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
	var err error
	var user *model.User

	params := mux.Vars(r)

	user, err = db.DB.GetUser(params["name"])
	if err != nil {
		ctx := context.Background()
		appClient := twitter.NewAppClient(ctx)
		fetchedUser, err := appClient.GetUserShow(params["name"])
		if err != nil {
			log.Fatal(err)
		}
		user = model.NewUser(fetchedUser)
		err = db.DB.AddUser(user)
		if err != nil {
			log.Fatal(err)
		}
	}

	json.NewEncoder(w).Encode(user)
}
