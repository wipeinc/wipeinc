package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/wipeinc/wipeinc/db"
	"github.com/wipeinc/wipeinc/model"
	"github.com/wipeinc/wipeinc/twitter"
)

var twitterAccessToken string
var twitterAccessTokenSecret string
var wipeincDB *db.PGDatabase

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	twitterAccessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	twitterAccessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	wipeincDB, err = db.New(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	err = wipeincDB.Connect()
	if err != nil {
		log.Fatal(err)
	}
}

func Serve() {
	router := mux.NewRouter()
	router.HandleFunc("/profile/{name}", ShowProfile)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:9000"},
		AllowCredentials: true,
		Debug:            false,
	})
	handler := c.Handler(router)
	http.ListenAndServe(":8000", handler)
}

func ShowProfile(w http.ResponseWriter, r *http.Request) {
	var err error
	var user *model.User

	params := mux.Vars(r)

	user, err = wipeincDB.GetUserByScreenName(params["name"])
	if err == sql.ErrNoRows {
		tc := twitter.NewClient(twitterAccessToken, twitterAccessTokenSecret)
		user, err = tc.GetUser(params["name"])
		if err != nil {
			log.Fatal(err)
		}
		err = wipeincDB.NewUser(user)
		if err != nil {
			log.Fatal(err)
		}
	}

	json.NewEncoder(w).Encode(user)
}
