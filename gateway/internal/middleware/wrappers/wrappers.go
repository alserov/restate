package wrappers

import (
	"context"
	"github.com/alserov/restate/gateway/internal/log"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func WithLogger(lg log.Logger) func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(string(ContextLogger), lg)
			return handlerFunc(c)
		}
	}
}

func ExtractLogger(ctx context.Context) log.Logger {
	l, ok := ctx.Value(ContextLogger).(log.Logger)
	if !ok {
		panic("can not get logger from context")
	}

	return l
}

func WithIdempotencyKey(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(string(ContextIdempotencyKey), uuid.NewString())
		return fn(c)
	}
}

func ExtractIdempotencyKey(ctx context.Context) string {
	key, ok := ctx.Value(ContextIdempotencyKey).(string)
	if !ok {
		panic("can not get idempotency key from context")
	}

	return key
}

type ContextKey string

const (
	ContextIdempotencyKey ContextKey = "key"
	ContextLogger         ContextKey = "log"
)
