package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/how8570/URL_Shortener/database"
)

func TestHandleUrlsGetMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/urls", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUrls)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestHandleUrlsDelMethod(t *testing.T) {
	req, err := http.NewRequest("DEL", "/v1/urls", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUrls)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestHandleUrlsPostMethod(t *testing.T) {
	database.ConnectDBWithPath("./../../database/urls.db")
	defer database.CloseDB()

	mcPostBody := map[string]interface{}{
		"url":      "http://about.dcard.tw",
		"expireAt": "2038-01-01",
	}
	body, _ := json.Marshal(mcPostBody)

	req, err := http.NewRequest("POST", "/v1/urls", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUrls)
	handler.ServeHTTP(rr, req)
	expectRes := `{
		"id":"dcard__",
		"shortUrl":"http://localhost/v1/redirect/dcard__"
		}`

	var j1, j2 Response
	json.Unmarshal([]byte(rr.Body.Bytes()), &j1)
	json.Unmarshal([]byte(expectRes), &j2)

	if j1 != j2 {
		t.Errorf("handler returned wrong body: got %v want %v",
			rr.Body.String(), expectRes)
	}

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestPostEmptyUrl(t *testing.T) {
	database.ConnectDBWithPath("./../../database/urls.db")
	defer database.CloseDB()

	mcPostBody := `"url": "", ExpireAt: "1700000000"`
	body := []byte(mcPostBody)

	req, err := http.NewRequest("POST", "/v1/urls", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUrls)
	handler.ServeHTTP(rr, req)

	t.Log(rr.Body.String())

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
