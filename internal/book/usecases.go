package book

import (
	_dto_ "github.com/guilhermewolke/fts-go-api/dto"
	authorDB "github.com/guilhermewolke/fts-go-api/repositories/author"
	bookDB "github.com/guilhermewolke/fts-go-api/repositories/book"
)

func (b *Book) List() ([]_dto_.BookList, error) {
	dto := make([]_dto_.BookList, 0)
	repository := bookDB.New(b.db, b.cache)

	books, err := repository.List()

	if err != nil {
		return dto, err
	}

	for _, book := range books {
		dto = append(dto, _dto_.BookList{ID: book.ID, Title: book.Title})
	}

	return dto, nil
}

func (b *Book) FindByID(id int64) (_dto_.Book, error) {
	repository := bookDB.New(b.db, b.cache)

	book, err := repository.FindByID(id)
	if err != nil {
		return _dto_.Book{}, err
	}

	authorRepository := authorDB.New(b.db, b.cache)
	bookAuthor, err := authorRepository.FindByID(book.AuthorID)

	if err != nil {
		return _dto_.Book{}, err
	}

	dto := _dto_.Book{ID: book.ID, Title: book.Title, Author: _dto_.ConvertToAuthorDTO(bookAuthor)}

	return dto, nil
}
