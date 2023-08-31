package dto

import (
	authorDB "github.com/guilhermewolke/fts-go-api/repositories/author"
)

type AuthorList struct {
	ID          int64  `json:"id"`
	Name        string `json:"nome"`
	Nationality string `json:"pais_de_origem"`
}

type Author struct {
	ID          int64            `json:"id"`
	Name        string           `json:"nome"`
	Nationality string           `json:"pais_de_origem"`
	Books       []AuthorBookList `json:"livros,omitempty"`
}

func ConvertToAuthorDTO(author authorDB.Author) BookAuthor {
	return BookAuthor{ID: author.ID, Name: author.Name, Nationality: author.Nationality}
}
