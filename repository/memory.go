package repository

import (
	"fmt"
	"sync"

	"github.com/wipeinc/wipeinc/model"
)

type memoryDB struct {
	mu          sync.Mutex
	usersByName map[string]*model.User
	users       map[int64]*model.User
}

func newMemoryDB() *memoryDB {
	return &memoryDB{
		usersByName: make(map[string]*model.User),
		users:       make(map[int64]*model.User),
	}
}

// Close closes the database.
func (db *memoryDB) Close() {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.users = nil
}

// GetUser retrieves a user by its ID.
func (db *memoryDB) GetUser(id int64) (*model.User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	user, ok := db.users[id]
	if !ok {
		return nil, fmt.Errorf("memorydb: book not found with ID %d", id)
	}
	return user, nil
}

// GetUserByScreenName retrieves a user by its screeenName.
func (db *memoryDB) GetUserByScreenName(screenName string) (*model.User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	user, ok := db.usersByName[screenName]
	if !ok {
		return nil, fmt.Errorf("memorydb: book not found with screenName %s", screenName)
	}
	return user, nil
}

// AddUser saves a given user
func (db *memoryDB) AddUser(u *model.User) (err error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.usersByName[u.ScreenName] = u
	db.users[u.ID] = u

	return nil
}
