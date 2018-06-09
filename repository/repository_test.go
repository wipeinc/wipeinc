package repository

import (
	"reflect"
	"testing"
	"time"

	"github.com/wipeinc/wipeinc/model"
)

var u = &model.User{
	ID:                   123,
	CreatedAt:            time.Now().Round(time.Microsecond).UTC(),
	FavouritesCount:      1,
	FollowersCount:       1,
	FriendsCount:         1,
	Name:                 "name",
	ProfileBannerURL:     "https://profilebannerurl.png",
	ProfileImageURLHTTPS: "https://profileimage.png",
	ScreenName:           "screenName",
	StatusesCount:        1,
	Description:          "Description",
	UpdatedAt:            time.Now().Round(time.Microsecond).UTC(),
	URL:                  "https://url",
}

func testDB(t *testing.T, db Database) {
	defer db.Close()
	err := db.AddUser(u)
	if err != nil {
		t.Fatalf("error on calling db.AddUser: %s", err.Error())
	}

	//try to fetch the inserted user by name
	var fetchedUser *model.User
	fetchedUser, err = db.GetUserByScreenName(u.ScreenName)
	if err != nil {
		t.Fatalf("error on calling db.GetUserByScreenName: %s", err.Error())
	}
	if !reflect.DeepEqual(fetchedUser, u) {
		t.Errorf("expected \n%q\n", u)
		t.Errorf("got:\n%q\n", fetchedUser)
		t.Fail()
	}

	// try to fetch the inserted user by ID
	fetchedUser, err = db.GetUser(u.ID)
	if err != nil {
		t.Fatalf("error on calling db.GetUser: %s", err.Error())
	}
	if !reflect.DeepEqual(fetchedUser, u) {
		t.Errorf("expected \n%q\n", u)
		t.Errorf("got:\n%q\n", fetchedUser)
		t.Fail()
	}

	// fetch an user by ID that does not exist
	_, err = db.GetUser(123456)
	if err == nil {
		t.Fatal("Expected GetUser to fail on unexisting user")
	}

	// fetch an user by name that does not exist
	_, err = db.GetUserByScreenName("i dont exist")
	if err == nil {
		t.Fatalf("Expected GetUserByScreenName to fail on unexisting user")
	}
}

func TestMemoryDB(t *testing.T) {
	testDB(t, newMemoryDB())
}
