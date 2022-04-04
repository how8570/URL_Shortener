package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/how8570/URL_Shortener/database"
)

func TestHandleRedirect(t *testing.T) {
	database.ConnectDBWithPath("./../../database/urls.db")
	defer database.CloseDB()

	req, err := http.NewRequest("GET", "/v1/redirect/dcard__", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"url": "dcard__"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRedirect)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusFound)
	}

}

func TestHandleRedirect404(t *testing.T) {
	database.ConnectDBWithPath("./../../database/urls.db")
	defer database.CloseDB()

	req, err := http.NewRequest("GET", "/v1/redirect/01234567", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRedirect)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

}
