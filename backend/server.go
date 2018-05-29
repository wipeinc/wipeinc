package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/wipeinc/wipeinc/db"
	"github.com/wipeinc/wipeinc/model"
	"github.com/wipeinc/wipeinc/twitter"
	"google.golang.org/appengine"
)

var twitterAccessToken string
var twitterAccessTokenSecret string

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/profile/{name}", ShowProfile)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://wipeinc.io"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	http.Handle("/", handler)
	appengine.Main()
}

func ShowProfile(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	params := mux.Vars(r)

	user, err = db.DB.GetUser(params["name"])
	if err != nil {
		ctx := appengine.NewContext(r)
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
