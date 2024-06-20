package wrappers

import (
	"context"
	"google.golang.org/grpc"
)

type Wrapper func(ctx context.Context, args ...any) context.Context

// WithGRPC - middleware for grpc handlers, wraps request context with values
func WithGRPC(wrs ...Wrapper) grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		for _, wrapper := range wrs {
			ctx = wrapper(ctx)
		}

		res, err := handler(ctx, req)
		return res, err
	})
}

type ContextKey string

var (
	ContextLogger         ContextKey = "log"
	ContextIdempotencyKey ContextKey = "key"
)
