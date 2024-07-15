package utils

import (
	"context"
	"time"
)

func NewLimiter(ctx context.Context, count int64, tick time.Duration) Limiter {
	t := make(chan struct{}, count)

	c := int(count)
	for i := 0; i < c; i++ {
		t <- struct{}{}
	}

	ticker := time.NewTicker(tick)

	go func() {
		for {
			select {
			case <-ticker.C:
				select {
				case t <- struct{}{}:
				default:
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return Limiter{
		count:  count,
		tokens: t,
	}
}

type tokens chan struct{}

type Limiter struct {
	count  int64
	tokens tokens
}

func (l Limiter) Allow(ctx context.Context) bool {
	select {
	case <-l.tokens:
		return true
	case <-ctx.Done():
		return false
	}
}
