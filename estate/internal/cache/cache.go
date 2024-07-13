package cache

import (
	"context"
	"github.com/alserov/restate/estate/internal/cache/redis"
)

type Cache interface {
	Set(ctx context.Context, key string, val any)
	Get(ctx context.Context, key string, dst any) bool
}

type Type uint

const (
	Redis Type = iota
)

func NewCache(t Type, addr ...string) Cache {
	switch t {
	case Redis:
		return redis.NewCache(addr[0])
	default:
		panic("invalid type")
	}
}
