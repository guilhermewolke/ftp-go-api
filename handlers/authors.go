package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guilhermewolke/fts-go-api/config"
	"github.com/guilhermewolke/fts-go-api/internal/author"
)

func AuthorsList(w http.ResponseWriter, r *http.Request) {
	db, err := config.DBConnect()
	redis := config.RedisConnect()

	if err != nil {
		panic(err)
	}

	authors := author.New(db, redis)
	payload, err := authors.List()

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}

func AuthorGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db, err := config.DBConnect()

	if err != nil {
		panic(err)
	}

	redis := config.RedisConnect()

	log.Printf("vars: %#v", vars)
	param := vars["id"]
	if param == "" {
		panic(err)
	}

	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		panic(err)
	}

	authors := author.New(db, redis)
	payload, err := authors.FindByID(id, false)

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

func AuthorGetBooks(w http.ResponseWriter, r *http.Request) {
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

	authors := author.New(db, redis)
	payload, err := authors.FindByID(id, true)
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
