package authorDB

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/guilhermewolke/fts-go-api/caching"
)

func New(db *sql.DB, cache caching.Cacheable) *AuthorDB {
	return &AuthorDB{db: db, cache: cache}
}

func (a *AuthorDB) List() ([]Author, error) {
	authors := make([]Author, 0)
	cache, err := a.cache.Get(caching.CacheAuthorListKey)

	if err != nil && err.Error() != "cache: key is missing" {

		log.Fatalf("Ocorreu um erro ao recuperar a chave '%s' do cache: %v", caching.CacheAuthorListKey, err)

	} else if cache != "" {
		log.Println("Recuperado do cache")
		return cachedAuthorList(cache), nil
	}

	query := `
		SELECT
			id,
			name,
			nationality
		FROM
			author
		ORDER BY
			name ASC;
	`
	rows, err := a.db.Query(query)

	if err != nil {
		return authors, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id                sql.NullInt64
			name, nationality sql.NullString
		)

		if err = rows.Scan(&id, &name, &nationality); err != nil {
			return authors, err
		}

		authors = append(authors, Author{ID: id.Int64, Name: name.String, Nationality: nationality.String})
	}

	//Sempre salvar a última consulta no banco de dados em cache
	a.cache.Set(caching.CacheAuthorListKey, authors)
	log.Println("Recuperado do Banco de dados e salvo no cache")
	return authors, nil
}

func (a *AuthorDB) FindByID(id int64) (Author, error) {
	key := fmt.Sprintf(caching.CacheAuthorKey, id)

	cache, err := a.cache.Get(key)

	if err != nil {
		if err.Error() != "cache: key is missing" {
			log.Fatalf("Ocorreu um erro ao recuperar a chave '%s' do cache: %v", key, err)
		}
	} else if cache != "" {
		log.Printf("Recuperado do cache")
		return cachedAuthor(cache), nil

	}

	query := `
		SELECT
			name,
			nationality
		FROM
			author
		WHERE
			id = ?;
		`
	row := a.db.QueryRow(query, id)

	var (
		name, nationality sql.NullString
	)

	err = row.Scan(&name, &nationality)

	if err != nil {
		return Author{}, err
	}

	author := Author{ID: id, Name: name.String, Nationality: nationality.String}

	a.cache.Set(key, author)
	log.Println("Recuperado do Banco de dados e salvo no cache")
	return author, nil
}

func cachedAuthorList(cache string) []Author {
	authors := make([]Author, 0)

	err := json.Unmarshal([]byte(cache), &authors)
	if err != nil {
		panic(err)
	}

	return authors
}

func cachedAuthor(cache string) Author {
	author := Author{}

	err := json.Unmarshal([]byte(cache), &author)
	if err != nil {
		panic(err)
	}

	return author
}
