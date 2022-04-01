package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/how8570/URL_Shortener/database"
)

type Request struct {
	Url      string `json:"url"`
	ExpireAt string `json:"expireAt"`
}

type Response struct {
	Id       string `json:"id"`
	ShortUrl string `json:"shortUrl"`
}

func HandleUrls(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	var request Request
	var response Response
	var shortUrl string
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&request)
	if err != nil {
		log.Println("json parse faild", err)
		http.Error(w, "wrong json format", http.StatusBadRequest)
		return
	}

	// TODO: check URL, ExpireDate correct format
	if request.Url == "" || request.ExpireAt == "" {
		http.Error(w, "url or expireAt is empty", http.StatusBadRequest)
		return
	}

	// check is url is exis
	stmt := "SELECT shortUrl FROM urls WHERE originUrl = ?"
	sqlStmt, err := database.UrlsDB.Prepare(stmt)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlStmt.Close()
	q, err := sqlStmt.Query(request.Url)
	if err != nil {
		log.Fatal(err)
	}

	if q.Next() {
		q.Scan(&shortUrl)

		response.Id = shortUrl
		response.ShortUrl = "http://localhost/v1/redirect/" + shortUrl

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	shortUrl = Hash(request.Url) // TODO: check hash is unique

	// insert short url to db
	stmt = `INSERT INTO urls(originUrl, shortUrl, expireAt, times)
			VALUES(?, ?, ?, ?)`
	sqlStmt, err = database.UrlsDB.Prepare(stmt)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlStmt.Close()

	res, err := sqlStmt.Exec(request.Url, shortUrl, request.ExpireAt, 0)

	if err != nil {
		log.Fatal(err)
	} else {
		rows, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		if rows != 1 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
		}
	}

	response.Id = shortUrl
	response.ShortUrl = "http://localhost/v1/redirect/" + shortUrl

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
