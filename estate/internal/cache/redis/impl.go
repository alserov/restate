package redis

import (
	"context"
	"github.com/alserov/restate/estate/internal/cache"
	"github.com/goccy/go-json"
	"time"

	rd "github.com/redis/go-redis/v9"
)

const exp = time.Minute * 15

func NewCache(addr string) *redis {
	return &redis{mustConnect(addr)}
}

var _ cache.Cache = &redis{}

type redis struct {
	cl *rd.Client
}

func (r *redis) Set(ctx context.Context, key string, val any) {
	r.cl.Set(ctx, key, val, exp)
}

func (r *redis) Get(ctx context.Context, key string, dst any) bool {
	val, err := r.cl.Get(ctx, key).Result()
	if err != nil {
		return false
	}

	if err = json.Unmarshal([]byte(val), &dst); err != nil {
		return false
	}

	return true
}
