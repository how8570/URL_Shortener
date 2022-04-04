package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var UrlsDB *sql.DB
var err error

func ConnectDB() {
	UrlsDB, err = sql.Open("sqlite3", "./database/urls.db")
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDBWithPath(path string) {
	UrlsDB, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	UrlsDB.Close()
}
