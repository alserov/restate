package grpc

import (
	"context"
	"github.com/alserov/restate/meetings/internal/utils"
	"google.golang.org/grpc"
)

// WithWrappers - middleware for grpc handlers, wraps request context with values
func WithWrappers(wrs ...utils.Wrapper) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		for _, wrapper := range wrs {
			ctx = wrapper(ctx)
		}

		res, err := handler(ctx, req)
		return res, err
	}
}
