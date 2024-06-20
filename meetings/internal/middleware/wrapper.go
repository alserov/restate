package middleware

import (
	"context"
	"github.com/alserov/restate/meetings/internal/log"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type Wrapper func(ctx context.Context, args ...any) context.Context

// WithWrapper - middleware for grpc handlers, wraps request context with values
func WithWrapper(wrs ...Wrapper) grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		for _, wrapper := range wrs {
			ctx = wrapper(ctx)
		}

		res, err := handler(ctx, req)
		return res, err
	})
}

func WithLogger(ctx context.Context, args ...any) context.Context {
	l, ok := args[0].(log.Logger)
	if !ok {
		panic("invalid argument")
	}

	ctx = context.WithValue(ctx, ContextLogger, l)
	return context.WithValue(ctx, ContextLogger, l)
}

func ExtractLogger(ctx context.Context) log.Logger {
	l, ok := ctx.Value(ContextLogger).(log.Logger)
	if !ok {
		panic("can not get logger from context")
	}

	return l
}

func WithIdempotencyKey(ctx context.Context, args ...any) context.Context {
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
