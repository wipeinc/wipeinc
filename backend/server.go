package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/profiler"
	"github.com/gorilla/mux"
	"github.com/wipeinc/wipeinc/db"
	"github.com/wipeinc/wipeinc/model"
	"github.com/wipeinc/wipeinc/twitter"
	"google.golang.org/appengine"
)

var twitterAccessToken string
var twitterAccessTokenSecret string

func main() {
	err := profiler.Start(profiler.Config{
		Service: "wipeinc",
	})
	if err != nil {
		log.Println(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/profile/{name}", ShowProfile)
	router.PathPrefix("/").HandlerFunc(ShowIndex)
	http.Handle("/", router)
	appengine.Main()
}

func ShowIndex(w http.ResponseWriter, r *http.Request) {
	index, err := Asset("static/index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	indexReader := bytes.NewBuffer(index)
	io.Copy(w, indexReader)
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

	json.NewEncoder(w).Encode(user)
}
