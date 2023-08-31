package authorDB

import (
	"database/sql"

	"github.com/guilhermewolke/fts-go-api/caching"
)

type AuthorDB struct {
	db    *sql.DB
	cache caching.Cacheable
}

type Author struct {
	ID          int64
	Name        string
	Nationality string
}
