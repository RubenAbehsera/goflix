package main

import "fmt"

type Movie struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	ReleaseDate string `db:"release_date"`
	Duration    int    `db:"duration"`
	TrailerUrl  string `db:"trailer_url"`
}

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func (m Movie) String() string {
	return fmt.Sprintf("id=%v, title=%v, releaseDate=%v, duration=%v, trailerUrl=%v", m.ID, m.Title, m.ReleaseDate, m.Duration, m.TrailerUrl)
}

func (u User) String() string {
	return fmt.Sprintf("id=%v, username=%v, password=%v", u.ID, u.Username, u.Password)
}
