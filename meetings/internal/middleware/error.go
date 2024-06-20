package middleware

import (
	"context"
	"github.com/alserov/restate/meetings/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func WithErrorHandler() grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		res, err := handler(ctx, req)
		if err != nil {
			msg, st := utils.FromError(err)
			return nil, status.Error(st, msg)
		}

		return res, nil
	})
}
