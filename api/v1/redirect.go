package v1

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/how8570/URL_Shortener/database"
)

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	log.Printf("RemoteAddr: %s request url: %s", r.RemoteAddr, r.URL)
	vars := mux.Vars(r)
	url := vars["url"]
	
	stmt := "SELECT originUrl, expireAt FROM urls WHERE shortUrl = ?;"

	sqlStmt, err := database.UrlsDB.Prepare(stmt)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlStmt.Close()
	q, err := sqlStmt.Query(url)
	if err != nil {
		log.Fatal(err)
	}
	defer q.Close()

	var originUrl string
	var expireAt int64
	if q.Next() {
		q.Scan(&originUrl, &expireAt)
		if time.Now().Unix() <= expireAt {
			http.Redirect(w, r, "http://"+originUrl, http.StatusFound)
			return
		}
	}

	http.NotFound(w, r)
}
