package v1

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/how8570/URL_Shortener/database"
)

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]

	stmt := "SELECT originUrl FROM urls WHERE shortUrl = ?"
	sqlStmt, err := database.UrlsDB.Prepare(stmt)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlStmt.Close()
	q, err := sqlStmt.Query(url)
	if err != nil {
		log.Fatal(err)
	}

	var originUrl string
	if q.Next() {
		q.Scan(&originUrl)
		http.Redirect(w, r, "https://"+originUrl, http.StatusFound)
		return
	}

	http.NotFound(w, r)
}
