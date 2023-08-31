package caching_redis

import (
	"context"
	"encoding/json"

	"github.com/go-redis/cache/v9"
	"github.com/guilhermewolke/fts-go-api/caching"
	"github.com/redis/go-redis/v9"
)

func New(options Options) *Redis {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			options.ServiceName1: options.ServicePort1}})

	return &Redis{
		Context: context.TODO(),
		Cache: cache.New(&cache.Options{
			Redis:      ring,
			LocalCache: cache.NewTinyLFU(options.TimeAmount, options.TimeDuration)})}
}

func (r *Redis) Set(key string, value interface{}) error {
	b, err := json.Marshal(value)

	if err != nil {
		return err
	}

	err = r.Cache.Set(&cache.Item{
		Ctx:   r.Context,
		Key:   key,
		Value: string(b),
		TTL:   caching.TTL})

	return err

}

func (r *Redis) Get(key string) (string, error) {
	var (
		value []byte
	)

	err := r.Cache.Get(r.Context, key, &value)

	return string(value), err
}
