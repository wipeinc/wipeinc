package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wipeinc/wipeinc/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.New(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	// anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	// anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	// client := twitter.NewClient(os.Getenv("TESTCLIENT_ACCESS_TOKEN"), os.Getenv("TESTCLIENT_ACCESS_TOKEN_SECRET"))
	// err = client.BlockUserFollowers("E8Emma")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// client := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	// url, credentials, err := client.AuthorizationURL("oob")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(url)
	// var code string
	// if _, err = fmt.Scan(&code); err != nil {
	// 	log.Fatal(err)
	// }
	// credentials, _, err = client.GetCredentials(credentials, code)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%#v\n", credentials)
	// client.BlockUserFollowers("E8Emma")
}
