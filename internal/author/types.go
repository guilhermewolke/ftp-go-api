package author

import (
	"database/sql"

	"github.com/guilhermewolke/fts-go-api/caching"
)

type Author struct {
	db    *sql.DB
	cache caching.Cacheable
}

func New(db *sql.DB, cache caching.Cacheable) *Author {
	return &Author{db: db, cache: cache}
}
