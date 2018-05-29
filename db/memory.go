package db

import (
	"fmt"
	"sync"

	"github.com/wipeinc/wipeinc/model"
)

type memoryDB struct {
	mu    sync.Mutex
	users map[string]*model.User
}

func newMemoryDB() *memoryDB {
	return &memoryDB{
		users: make(map[string]*model.User),
	}
}

// Close closes the database.
func (db *memoryDB) Close() {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.users = nil
}

// GetBook retrieves a book by its ID.
func (db *memoryDB) GetUser(screenName string) (*model.User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	user, ok := db.users[screenName]
	if !ok {
		return nil, fmt.Errorf("memorydb: book not found with screenName %s", screenName)
	}
	return user, nil
}

// AddBook saves a given book, assigning it a new ID.
func (db *memoryDB) AddUser(u *model.User) (err error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.users[u.TwitterUser.ScreenName] = u

	return nil
}
