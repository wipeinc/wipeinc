package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/wipeinc/wipeinc/db"
	"github.com/wipeinc/wipeinc/model"
	"github.com/wipeinc/wipeinc/twitter"
)

var twitterAccessToken string
var twitterAccessTokenSecret string

func init() {
	twitterAccessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	twitterAccessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	router := mux.NewRouter()
	router.HandleFunc("/profile/{name}", ShowProfile)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://wipeinc.io"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	http.Handle("/", handler)
}

func ShowProfile(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	params := mux.Vars(r)

	user, err = db.DB.GetUser(params["name"])
	if err != nil {
		tc := twitter.NewClient(twitterAccessToken, twitterAccessTokenSecret)
		user, err = tc.GetUser(params["name"])
		if err != nil {
			log.Fatal(err)
		}
		err = db.DB.AddUser(user)
		if err != nil {
			log.Fatal(err)
		}
	}

	json.NewEncoder(w).Encode(user)
}
