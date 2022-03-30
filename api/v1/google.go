package v1

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]

	fmt.Println("Redirecting to: ", url)
	

	http.Redirect(w, r, "https://"+url, http.StatusFound)
}
