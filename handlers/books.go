package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guilhermewolke/fts-go-api/config"
	"github.com/guilhermewolke/fts-go-api/internal/book"
)

func BooksList(w http.ResponseWriter, r *http.Request) {
	db, err := config.DBConnect()
	if err != nil {
		panic(err)
	}

	redis := config.RedisConnect()

	books := book.New(db, redis)
	payload, err := books.List()

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}

func BookGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db, err := config.DBConnect()

	if err != nil {
		panic(err)
	}

	redis := config.RedisConnect()

	param := vars["id"]
	if param == "" {
		panic(err)
	}

	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		panic(err)
	}

	books := book.New(db, redis)
	payload, err := books.FindByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Registro não encontrado"))
			return
		} else {
			panic(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}
