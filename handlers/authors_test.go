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

func TestAuthorList(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/autores/", nil)
	w := httptest.NewRecorder()

	router := Router()

	router.ServeHTTP(w, req)

	res := w.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected StatusCode to be '%d' but got %d", http.StatusOK, res.StatusCode)
	}

}

func TestAuthorGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/autores/1", nil)
	w := httptest.NewRecorder()

	router := Router()

	router.ServeHTTP(w, req)

	res := w.Result()

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		t.Fatalf("Erro ao ler o body da resposta: %#v", err)
	}

	var author dto.Author

	err = json.Unmarshal(body, &author)

	if err != nil {
		t.Fatalf("Erro ao converter o body da resposta: %#v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected StatusCode to be '%d' but got %d", http.StatusOK, res.StatusCode)
	}

	expected := dto.Author{
		ID:          1,
		Name:        "Fiodor Dostoievski",
		Nationality: "Rússia"}

	if !reflect.DeepEqual(expected, author) {
		t.Fatalf("Expected Author to be '%v' but got '%v'", expected, author)
	}

}

func TestAuthorGetBooks(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/autores/1/livros", nil)
	w := httptest.NewRecorder()

	router := Router()

	router.ServeHTTP(w, req)

	res := w.Result()

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		t.Fatalf("Erro ao ler o body da resposta: %#v", err)
	}

	var author dto.Author

	err = json.Unmarshal(body, &author)

	if err != nil {
		t.Fatalf("Erro ao converter o body da resposta: %#v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected StatusCode to be '%d' but got %d", http.StatusOK, res.StatusCode)
	}

	books := make([]dto.AuthorBookList, 0)

	books = append(books, dto.AuthorBookList{ID: 3, Title: "Crime e Castigo"}, dto.AuthorBookList{ID: 1, Title: "Humilhados e Ofendidos"}, dto.AuthorBookList{ID: 2, Title: "Noites Brancas"})

	expected := dto.Author{
		ID:          1,
		Name:        "Fiodor Dostoievski",
		Nationality: "Rússia",
		Books:       books}

	if !reflect.DeepEqual(expected, author) {
		t.Fatalf("Expected Author to be '%#v' but got '%#v'", expected, author)
	}

}
