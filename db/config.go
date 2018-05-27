package db

import "github.com/wipeinc/wipeinc/model"

var (
	// DB app database
	DB Database
)

// Database interface
type Database interface {
	GetUser(screenName string) (*model.User, error)
	AddUser(u *model.User) error
}

func init() {
	DB = newMemoryDB()
}
