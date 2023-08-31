package dto

import bookDB "github.com/guilhermewolke/fts-go-api/repositories/book"

type Book struct {
	ID     int64      `json:"id"`
	Title  string     `json:"titulo"`
	Author BookAuthor `json:"autor"`
}

type BookList struct {
	ID    int64  `json:"id"`
	Title string `json:"titulo"`
}

type AuthorBookList struct {
	ID    int64  `json:"id"`
	Title string `json:"titulo"`
}

type BookAuthor struct {
	ID          int64  `json:"id"`
	Name        string `json:"nome"`
	Nationality string `json:"pais_de_origem"`
}

func ConvertToBooksListDTO(books []bookDB.Book) []AuthorBookList {
	list := make([]AuthorBookList, 0)

	for _, book := range books {
		list = append(list, AuthorBookList{ID: book.ID, Title: book.Title})
	}

	return list
}
