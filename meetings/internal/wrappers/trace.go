package wrappers

import (
	"context"
	"github.com/google/uuid"
)

func WithIdempotencyKey(ctx context.Context, args ...any) context.Context {
	return context.WithValue(ctx, ContextIdempotencyKey, uuid.NewString())
}

func ExtractIdempotencyKey(ctx context.Context) string {
	return ctx.Value(ContextIdempotencyKey).(string)
}
