package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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

	t.Log(rr.Body.String())

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}
