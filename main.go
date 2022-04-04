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
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// init random seed
	rand.Seed(time.Now().UnixNano())

	// Connect to the database
	database.ConnectDB()
	defer database.CloseDB()
	log.Println("database connected")

	// Router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", web.HandleIndex)

	// APIs v1
	router.HandleFunc("/v1/", v1.Handlev1)
	router.HandleFunc("/v1/urls", v1.HandleUrls)
	router.HandleFunc("/v1/redirect/{url}", v1.HandleRedirect)

	log.Println("Start listening on port 80...")
	err := http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
