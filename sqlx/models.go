package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

type Song struct {
	ID     int64
	Title  string
	Artist string
	Year   int
}

func init() {
	var err error
	db, err = sqlx.Connect("sqlite3", "./songs.db")
	if err != nil {
		log.Fatalf("Fail to connect to database: %v", err)
	}

	db.MustExec(`
CREATE TABLE IF NOT EXISTS song (
    id INTEGER NOT NULL PRIMARY KEY ASC,
    title TEXT,
    artist TEXT,
    year INTEGER
);`)
}
