package tests

import (
	"fmt"
	"net/http"
)

func NameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Who"
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, %s", name)
}