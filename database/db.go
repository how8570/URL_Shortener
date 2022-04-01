package database

import "database/sql"

var UrlsDB *sql.DB
var err error

func ConnectDB() {
	UrlsDB, err = sql.Open("sqlite3", "./database/urls.db")
	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	UrlsDB.Close()
}
