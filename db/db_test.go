package db_test

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	// postgres driver
	_ "github.com/lib/pq"
	"github.com/wipeinc/wipeinc/db"
	"github.com/wipeinc/wipeinc/model"
)

const testDatabaseURL = "postgres://localhost:5432/wipeinc_test?sslmode=disable"

var tdb *sql.DB

func init() {
	var err error
	tdb, err = sql.Open("postgres", testDatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
}

func teardown(table, where string, params ...interface{}) func(t *testing.T) {
	return func(t *testing.T) {
		t.Logf("Treadown: DELETE FROM \"%s\" WHERE %s\tparams: %v", table, where, params)
		_, err := tdb.Exec(
			fmt.Sprintf("DELETE FROM \"%s\" WHERE %s", table, where),
			params...)
		if err != nil {
			// Usually I'd return the error, but using panic means less code in
			// my tests, which is more important.
			t.Fatalf("Problem tearing down %s data: %v", table, err)
		}
	}
}

func TestNewUser(t *testing.T) {
	user := &model.User{
		ID:              123,
		URL:             "url",
		Name:            "name",
		ScreenName:      "screenName",
		Location:        "Location",
		Lang:            "Lang",
		Description:     "Description",
		BackgroundImage: "BackgroundImage",
		Image:           "Image",
		FavouritesCount: 1,
		FollowersCount:  1,
		FriendsCount:    1,
		UpdatedAt:       time.Now().Round(time.Microsecond).UTC(),
	}
	teardownCase := teardown("users", "ID = $1", user.ID)
	defer teardownCase(t)

	db, _ := db.NewFromDB(tdb)
	err := db.NewUser(user)
	if err != nil {
		t.Fatalf("error on calling NewUser; %#v", err)
	}

	fetchUser, err := db.GetUser(123)
	if err != nil {
		t.Fatalf("error on calling NewUser; %#v", err)
	}
	if !reflect.DeepEqual(fetchUser, user) {
		t.Errorf("expected \n%q", user)
		t.Errorf("\ngot:\n%q", fetchUser)
	}
}
