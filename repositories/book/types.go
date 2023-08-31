package bookDB

import (
	"database/sql"

	"github.com/guilhermewolke/fts-go-api/caching"
)

type BookDB struct {
	db    *sql.DB
	cache caching.Cacheable
}

type Book struct {
	ID       int64
	AuthorID int64
	Title    string
}
