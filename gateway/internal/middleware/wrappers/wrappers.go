package wrappers

import (
	"context"
	"github.com/alserov/restate/gateway/internal/log"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func Ctx(c echo.Context) context.Context {
	return c.Get(string(Context)).(context.Context)
}

func WithLogger(lg log.Logger) func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Get(string(Context))
			if ctx == nil {
				ctx = context.WithValue(context.Background(), ContextLogger, lg)
			} else {
				ctx = context.WithValue(ctx.(context.Context), ContextLogger, lg)
			}

			c.Set(string(Context), ctx)

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

func WithIdempotencyKey(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Get(string(Context))
		if ctx == nil {
			ctx = context.WithValue(context.Background(), ContextIdempotencyKey, uuid.NewString())
		} else {
			ctx = context.WithValue(ctx.(context.Context), ContextIdempotencyKey, uuid.NewString())
		}

		c.Set(string(Context), ctx)

		return handlerFunc(c)
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

	Context ContextKey = "ctx"
)
