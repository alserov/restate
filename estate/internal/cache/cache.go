package cache

import (
	"context"
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
		return newRedis(addr[0])
	default:
		panic("invalid type")
	}
}
