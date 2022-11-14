package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Store interface {
	Open() error
	Close() error

	GetMovies() ([]*Movie, error)
	GetMovieById(id int64) (*Movie, error)
	CreateMovie(*Movie) error
	GetUsers() ([]*User, error)
	FindUser(username string, password string) (bool, error)
}

type dbStore struct {
	db *sqlx.DB
}

var schema = `
	CREATE TABLE IF NOT EXISTS movie
	(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		release_date TEXT,
		duration INTEGER, 
		trailer_url TEXT
	);
	CREATE TABLE IF NOT EXISTS user
	(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		password TEXT
	);
`

func (store *dbStore) Open() error {
	db, err := sqlx.Connect("sqlite3", "goflix.db")
	if err != nil {
		return err
	}
	log.Println("Connected to DB")
	db.MustExec(schema)
	store.db = db
	return nil
}

func (store *dbStore) Close() error {
	return store.db.Close()
}

func (store *dbStore) GetMovies() ([]*Movie, error) {
	var movies []*Movie
	err := store.db.Select(&movies, "SELECT * FROM movie")
	if err != nil {
		return movies, err
	}
	return movies, nil
}

func (store *dbStore) GetMovieById(id int64) (*Movie, error) {
	var movie = &Movie{}

	err := store.db.Get(movie, "SELECT * FROM movie WHERE id = $1", id)

	if err != nil {
		return movie, err
	}
	return movie, nil
}

func (store *dbStore) CreateMovie(m *Movie) error {
	res, err := store.db.Exec("INSERT INTO movie (title, release_date, duration,trailer_url ) VALUES (?, ?, ?, ?)",
		m.Title, m.ReleaseDate, m.Duration, m.TrailerUrl)
	if err != nil {
		return err
	}

	m.ID, err = res.LastInsertId()
	return err
}

func (store *dbStore) GetUsers() ([]*User, error) {
	var users []*User
	err := store.db.Select(&users, "SELECT * FROM user")
	if err != nil {
		return users, err
	}
	return users, nil
}

func (store *dbStore) FindUser(username string, password string) (bool, error) {
	var count int
	err := store.db.Get(&count, "SELECT COUNT(id) FROM user WHERE username = $1 AND password = $2", username, password)

	if err != nil {
		return false, err
	}
	return count == 1, nil
}
