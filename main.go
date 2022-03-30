package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/how8570/URL_Shortener/api/v1"
	"github.com/how8570/URL_Shortener/web"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// DB
	db, err := sql.Open("sqlite3", "./database/test.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from test")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var a int
		var b int
		var c string
		err = rows.Scan(&a, &b, &c)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(a, b, c)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/", web.HandleIndex)

	// APIs v1
	router.HandleFunc("/v1/", v1.Handlev1)
	router.HandleFunc("/v1/redirect/{url}", v1.HandleRedirect)

	err = http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
