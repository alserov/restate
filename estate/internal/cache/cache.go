package cache

import (
	"context"
)

type Cache interface {
	Set(ctx context.Context, key string, val any)
	Get(ctx context.Context, key string, dst any) bool
}
