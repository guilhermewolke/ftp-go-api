package author

import (
	_dto_ "github.com/guilhermewolke/fts-go-api/dto"
	authorDB "github.com/guilhermewolke/fts-go-api/repositories/author"
	bookDB "github.com/guilhermewolke/fts-go-api/repositories/book"
)

func (a *Author) List() ([]_dto_.AuthorList, error) {
	dto := make([]_dto_.AuthorList, 0)
	repository := authorDB.New(a.db, a.cache)

	authors, err := repository.List()

	if err != nil {
		return dto, err
	}

	for _, author := range authors {
		dto = append(dto, _dto_.AuthorList{ID: author.ID, Name: author.Name, Nationality: author.Nationality})
	}

	return dto, nil
}

func (a *Author) FindByID(id int64, withBooks bool) (_dto_.Author, error) {
	repository := authorDB.New(a.db, a.cache)

	author, err := repository.FindByID(id)
	if err != nil {
		return _dto_.Author{}, err
	}

	dto := _dto_.Author{ID: author.ID, Name: author.Name, Nationality: author.Nationality}

	if withBooks {
		bookRepository := bookDB.New(a.db, a.cache)

		books, err := bookRepository.ListByAuthor(id)

		if err != nil {
			return dto, err
		}
		dto.Books = _dto_.ConvertToBooksListDTO(books)

	}

	return dto, nil
}
