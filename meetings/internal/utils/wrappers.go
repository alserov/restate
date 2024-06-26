package utils

import (
	"context"
	"github.com/alserov/restate/meetings/internal/log"
	"github.com/google/uuid"
)

type Wrapper func(ctx context.Context) context.Context

func WithLogger(l log.Logger) Wrapper {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ContextLogger, l)
	}
}

func ExtractLogger(ctx context.Context) log.Logger {
	l, ok := ctx.Value(ContextLogger).(log.Logger)
	if !ok {
		panic("can not get logger from context")
	}

	return l
}

func WithIdempotencyKey(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextIdempotencyKey, uuid.NewString())
}

func ExtractIdempotencyKey(ctx context.Context) string {
	return ctx.Value(ContextIdempotencyKey).(string)
}

type ContextKey string

var (
	ContextLogger         ContextKey = "log"
	ContextIdempotencyKey ContextKey = "key"
)
