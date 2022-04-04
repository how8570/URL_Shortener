package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

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
	log.Printf("RemoteAddr: %s Request url: %s Method: %s", r.RemoteAddr, r.URL, r.Method)
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

	// trim URL, ExpireDate to correct format
	request.Url = strings.TrimSpace(request.Url)
	request.ExpireAt = strings.TrimSpace(request.ExpireAt)
	u, _ := url.Parse(request.Url)
	request.Url = u.Host + u.Path
	if u.RawQuery != "" {
		request.Url += "?" + u.RawQuery
	}
	if u.Fragment != "" {
		request.Url += "#" + u.Fragment
	}

	if request.Url == "" {
		http.Error(w, "url is empty", http.StatusBadRequest)
		return
	}

	log.Printf("RemoteAddr: %s | Json decode value{request url: %s , expireAt: %s}", r.RemoteAddr, request.Url, request.ExpireAt)

	// check db exist
	if database.UrlsDB == nil {
		log.Println("database is NOT connected, please run `go run main.go`")
		http.Error(w, "database is NOT connected", http.StatusInternalServerError)
		return
	}

	// check is url is exis
	stmt := "SELECT shortUrl, expireAt FROM urls WHERE originUrl = ?"
	sqlStmt, err := database.UrlsDB.Prepare(stmt)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlStmt.Close()
	q, err := sqlStmt.Query(request.Url)
	if err != nil {
		log.Fatal(err)
	}

	var extend bool
	var expireAt int64
	if q.Next() {
		q.Scan(&shortUrl, &expireAt)
		extend = true
	}
	q.Close()

	if extend {
		if expireAt <= toUnixTime(request.ExpireAt) {
			stmt = "UPDATE urls SET expireAt = ? WHERE shortUrl = ? ;"
			sqlStmt, err = database.UrlsDB.Prepare(stmt)
			if err != nil {
				log.Fatal(err)
			}
			defer sqlStmt.Close()
			res, err := sqlStmt.Exec(toUnixTime(request.ExpireAt), shortUrl)
			if err != nil {
				log.Fatal(err)
			}

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
		return
	}

	shortUrl = Hash(request.Url) // TODO: check hash is unique

	// insert short url to db
	stmt = `INSERT INTO urls(originUrl, shortUrl, expireAt, times)
			VALUES(?, ?, ?, ?);`
	sqlStmt, err = database.UrlsDB.Prepare(stmt)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlStmt.Close()

	res, err := sqlStmt.Exec(request.Url, shortUrl, toUnixTime(request.ExpireAt), 0)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
	}

	response.Id = shortUrl
	response.ShortUrl = "http://localhost/v1/redirect/" + shortUrl

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func toUnixTime(expireAt string) int64 {
	if expireAt == "" {
		return time.Now().AddDate(0, 0, 7).Unix()
	}

	var t time.Time
	var err error

	const format string = "2006-01-02T15:04:05Z"
	t, err = time.Parse(format, expireAt)
	if err == nil {
		return t.Unix()
	}

	t, err = time.Parse("2006-01-02", expireAt)
	if err == nil {
		return t.Unix()
	}

	unixTimestmp, err := strconv.ParseInt(expireAt, 10, 64)
	if err == nil && unixTimestmp >= 0 && unixTimestmp <= 2147483647 {
		return time.Unix(unixTimestmp, 0).Unix()
	}

	return time.Now().AddDate(0, 0, 7).Unix()
}
