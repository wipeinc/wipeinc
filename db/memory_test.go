// +build test

package db

import (
	"reflect"
	"testing"
	"time"

	"github.com/wipeinc/wipeinc/model"
)

func TestMemoryDB(t *testing.T) {
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
		CreatedAt:       time.Now().Round(time.Microsecond).UTC(),
		UpdatedAt:       time.Now().Round(time.Microsecond).UTC(),
	}

	db := newMemoryDB()
	defer db.Close()

	err := db.AddUser(user)
	if err != nil {
		t.Fatalf("error on calling NewUser; %#v", err)
	}

	fetchUser, err := db.GetUser(user.ScreenName)
	if err != nil {
		t.Fatalf("error on calling NewUser; %#v", err)
	}
	if !reflect.DeepEqual(fetchUser, user) {
		t.Errorf("expected \n%q", user)
		t.Errorf("\ngot:\n%q", fetchUser)
	}

	_, err = db.GetUser("does not exist")
	if err == nil {
		t.Fatalf("expected Error got; %#v", err)
	}
}
