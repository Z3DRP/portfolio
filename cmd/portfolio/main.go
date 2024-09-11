package main

import (
	"net/http"

	"github.com/Z3DRP/portfolio/internal/routes"
)

func run() {
	mux := routes.NewRouter()

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func main() {
	run()
}
