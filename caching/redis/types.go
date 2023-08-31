package caching_redis

import (
	"context"
	"time"

	"github.com/go-redis/cache/v9"
)

type Redis struct {
	Context context.Context
	Cache   *cache.Cache
}

type Options struct {
	TimeAmount   int
	TimeDuration time.Duration
	ServiceName1 string
	ServicePort1 string
}
