package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/setup", Setup).Methods(http.MethodGet)
	r.HandleFunc("/autores", AuthorsList).Methods(http.MethodGet)
	r.HandleFunc("/autores/", AuthorsList).Methods(http.MethodGet)
	r.HandleFunc("/autores/{id}", AuthorGet).Methods(http.MethodGet)
	r.HandleFunc("/autores/{id}/livros", AuthorGetBooks).Methods(http.MethodGet)
	r.HandleFunc("/livros", BooksList).Methods(http.MethodGet)
	r.HandleFunc("/livros/", BooksList).Methods(http.MethodGet)
	r.HandleFunc("/livros/{id}", BookGet).Methods(http.MethodGet)

	return r
}
