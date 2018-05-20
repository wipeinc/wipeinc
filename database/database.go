package database

import (
	"database/sql"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/wipeinc/wipeinc/model"
	// file driver
	_ "github.com/golang-migrate/migrate/source/file"

	// postgres driver
	_ "github.com/lib/pq"
)

// PGDatabase is implementation of Database interface for connecting to database.
type PGDatabase struct {
	url string
	db  *sql.DB
}

// New create new Databse struct
// databaseURL: postgres://localhost:5432/database?sslmode=enable
func New(databaseURL string) (*PGDatabase, error) {
	return &PGDatabase{
		url: databaseURL,
	}, nil
}

// Connect connect to the Postgres database
func (db *PGDatabase) Connect() error {
	var err error
	db.db, err = sql.Open("postgres", db.url)
	return err
}

// Migrate proceed to the migrations
func (db *PGDatabase) Migrate() error {
	driver, err := postgres.WithInstance(db.db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return err
	}
	return m.Steps(1)
}

// InsertUser insert a new twitter user into the database
func (db *PGDatabase) InsertUser(user *model.User) error {
	sqlStatement := `
		INSERT INTO users
		(
			id
		  url
		  name
		  screen_name
		  location
		  description
		  background_image
		  image
		  favourite_count
		  followers_count
		  friends_count
		  updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	stmt, err := db.db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		user.ID,
		user.URL,
		user.Name,
		user.ScreenName,
		user.Location,
		user.Lang,
		user.Description,
		user.BackgroundImage,
		user.Image,
		user.FavouritesCount,
		user.FollowersCount,
		user.FriendsCount,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func scanUser(row *sql.Row) (user *model.User, err error) {
	err = row.Scan(
		user.ID,
		user.URL,
		user.Name,
		user.ScreenName,
		user.Location,
		user.Lang,
		user.Description,
		user.BackgroundImage,
		user.Image,
		user.FavouritesCount,
		user.FollowersCount,
		user.FriendsCount,
		user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, err
}

// GetUser return twitter user from the database
func (db *PGDatabase) GetUser(id int64) (*model.User, error) {
	sqlStatement := `SELECT * from users WHERE id = $1`
	row := db.db.QueryRow(sqlStatement, id)
	return scanUser(row)
}
