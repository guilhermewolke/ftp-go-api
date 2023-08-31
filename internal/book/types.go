package book

import (
	"database/sql"

	"github.com/guilhermewolke/fts-go-api/caching"
)

type Book struct {
	db    *sql.DB
	cache caching.Cacheable
}

func New(db *sql.DB, cache caching.Cacheable) *Book {
	return &Book{db: db, cache: cache}
}
