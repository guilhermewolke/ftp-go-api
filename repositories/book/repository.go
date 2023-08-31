package bookDB

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/guilhermewolke/fts-go-api/caching"
)

func New(db *sql.DB, cache caching.Cacheable) *BookDB {
	return &BookDB{db: db, cache: cache}
}

func (b *BookDB) List() ([]Book, error) {
	books := make([]Book, 0)
	cache, err := b.cache.Get(caching.CacheBookListKey)

	if err != nil && err.Error() != "cache: key is missing" {
		log.Fatalf("Ocorreu um erro ao recuperar a chave '%s' do cache: %v", caching.CacheBookListKey, err)
	} else if cache != "" {
		log.Println("Recuperado do cache")
		return cachedBookList(cache), nil
	}

	query := `
		SELECT
			id,
			title
		FROM
			book
		ORDER BY
			title ASC;
	`
	rows, err := b.db.Query(query)

	if err != nil {
		return books, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id    sql.NullInt64
			title sql.NullString
		)

		if err = rows.Scan(&id, &title); err != nil {
			return books, err
		}

		books = append(books, Book{ID: id.Int64, Title: title.String})
	}

	//Sempre salvar a última consulta no banco de dados em cache
	b.cache.Set(caching.CacheBookListKey, books)
	log.Println("Recuperado do Banco de dados e salvo no cache")
	return books, nil
}

func (b *BookDB) FindByID(id int64) (Book, error) {
	key := fmt.Sprintf(caching.CacheBookKey, id)

	cache, err := b.cache.Get(key)

	if err != nil && err.Error() != "cache: key is missing" {
		log.Fatalf("Ocorreu um erro ao recuperar a chave '%s' do cache: %v", key, err)
	} else if cache != "" {
		log.Println("Recuperado do cache")
		return cachedBook(cache), nil
	}

	query := `
		SELECT
			author_id,
			title
		FROM
			book
		WHERE
			id = ?;
		`
	row := b.db.QueryRow(query, id)

	var (
		author_id sql.NullInt64
		title     sql.NullString
	)

	err = row.Scan(&author_id, &title)

	if err != nil {
		return Book{}, err
	}

	book := Book{ID: id, Title: title.String, AuthorID: author_id.Int64}

	//Sempre salvar a última consulta no banco de dados em cache
	b.cache.Set(key, book)
	log.Println("Recuperado do Banco de dados e salvo no cache")
	return book, nil
}

func (b *BookDB) ListByAuthor(id int64) ([]Book, error) {
	key := fmt.Sprintf(caching.CacheAuthorBooksKey, id)

	cache, err := b.cache.Get(key)

	if err != nil && err.Error() != "cache: key is missing" {
		log.Fatalf("Ocorreu um erro ao recuperar a chave '%s' do cache: %v", key, err)
	} else if cache != "" {
		log.Println("Recuperado do cache")
		return cachedBookList(cache), nil
	}

	books := make([]Book, 0)
	query := `
		SELECT
			id,
			title
		FROM
			book
		WHERE
			author_id = ?
		ORDER BY
			title ASC;
	`
	rows, err := b.db.Query(query, id)

	if err != nil {
		return books, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id    sql.NullInt64
			title sql.NullString
		)

		if err = rows.Scan(&id, &title); err != nil {
			return books, err
		}

		books = append(books, Book{ID: id.Int64, Title: title.String})
	}

	//Sempre salvar a última consulta no banco de dados em cache
	b.cache.Set(key, books)
	log.Println("Recuperado do Banco de dados e salvo no cache")
	return books, nil
}

func cachedBookList(cache string) []Book {
	books := make([]Book, 0)

	err := json.Unmarshal([]byte(cache), &books)
	if err != nil {
		panic(err)
	}

	return books
}

func cachedBook(cache string) Book {
	book := Book{}

	err := json.Unmarshal([]byte(cache), &book)
	if err != nil {
		panic(err)
	}

	return book
}
