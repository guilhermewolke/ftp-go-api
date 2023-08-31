package caching

import "time"

type Cacheable interface {
	Set(key string, value interface{}) error
	Get(key string) (string, error)
}

const (
	TTL                 = time.Second * 10
	CacheAuthorListKey  = "author.list"
	CacheAuthorKey      = "author.%d"
	CacheAuthorBooksKey = "author.%d.book"
	CacheBookListKey    = "book.list"
	CacheBookKey        = "book.%d"
)
