package config

import (
	"time"

	"github.com/guilhermewolke/fts-go-api/caching"
	caching_redis "github.com/guilhermewolke/fts-go-api/caching/redis"
)

func RedisConnect() caching.Cacheable {
	options := caching_redis.Options{
		TimeAmount:   10,
		TimeDuration: time.Second,
		ServiceName1: "server1",
		ServicePort1: "redis-fts:6379"}

	redis := caching_redis.New(options)
	return redis
}
