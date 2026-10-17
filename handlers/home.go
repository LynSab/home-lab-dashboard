package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>Hello, World!</p>")
	w.Header().Add("Content-Type", "text/html")
}
