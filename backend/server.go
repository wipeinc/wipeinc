package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wipeinc/wipeinc/db"
	"github.com/wipeinc/wipeinc/model"
	"github.com/wipeinc/wipeinc/twitter"
	"google.golang.org/appengine"
)

var twitterAccessToken string
var twitterAccessTokenSecret string

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/profile/{name}", ShowProfile)
	http.Handle("/", router)
	appengine.Main()
}

// ShowProfile route for /api/profile/{screenName}
func ShowProfile(w http.ResponseWriter, r *http.Request) {
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
	if appengine.IsDevAppServer() {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}

	json.NewEncoder(w).Encode(user)
}
