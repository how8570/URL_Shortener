package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	v1 "github.com/how8570/URL_Shortener/api/v1"
	"github.com/how8570/URL_Shortener/database"
	"github.com/how8570/URL_Shortener/web"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// init random seed
	rand.Seed(time.Now().UnixNano())

	// Connect to the database
	database.ConnectDB()
	defer database.CloseDB()

	// Router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", web.HandleIndex)

	// APIs v1
	router.HandleFunc("/v1/", v1.Handlev1)
	router.HandleFunc("/v1/urls", v1.HandleUrls)
	router.HandleFunc("/v1/redirect/{url}", v1.HandleRedirect)

	err := http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
