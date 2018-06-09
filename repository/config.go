package repository

import "github.com/wipeinc/wipeinc/model"

var (
	// DB app database
	DB Database
)

// Database interface
type Database interface {
	AddUser(u *model.User) error
	GetUserByScreenName(screenName string) (*model.User, error)
	GetUser(id int64) (*model.User, error)
	Close()
}

func init() {
	DB = newMemoryDB()
}
