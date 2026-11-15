package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LynSab/home-lab-dashboard/handlers"
)

func main() {
	var port string
	if os.Getenv("PORT") != "" {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	} else {
		port = ":8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.HomeHandler)

	log.Fatal(http.ListenAndServe(port, mux))
}
