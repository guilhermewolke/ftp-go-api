package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/guilhermewolke/fts-go-api/dto"
)

func TestBooksList(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/livros/", nil)
	w := httptest.NewRecorder()

	router := Router()

	router.ServeHTTP(w, req)

	res := w.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected StatusCode to be '%d' but got %d", http.StatusOK, res.StatusCode)
	}

}

func TestBookGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/livros/1", nil)
	w := httptest.NewRecorder()

	router := Router()

	router.ServeHTTP(w, req)

	res := w.Result()

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		t.Fatalf("Erro ao ler o body da resposta: %#v", err)
	}

	var book dto.Book

	err = json.Unmarshal(body, &book)

	if err != nil {
		t.Fatalf("Erro ao converter o body da resposta: %#v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected StatusCode to be '%d' but got %d", http.StatusOK, res.StatusCode)
	}

	expected := dto.Book{
		ID:    1,
		Title: "Humilhados e Ofendidos",
		Author: dto.BookAuthor{
			ID:          1,
			Name:        "Fiodor Dostoievski",
			Nationality: "Rússia"}}

	if !reflect.DeepEqual(expected, book) {
		t.Fatalf("Expected Book to be '%v' but got '%v'", expected, book)
	}

}
