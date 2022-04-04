package v1

import (
	"fmt"
	"net/http"
)

func Handlev1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello v1, maybe put docs here")
}
