package main

import (
	"net/http"

	"github.com/guilhermewolke/fts-go-api/handlers"
)

func main() {

	http.Handle("/", handlers.Router())
	http.ListenAndServe(":8080", nil)
}
