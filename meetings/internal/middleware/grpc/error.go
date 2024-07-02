package grpc

import (
	"context"
	"github.com/alserov/restate/meetings/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WithErrorHandler() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		res, err := handler(ctx, req)
		if err != nil {
			msg, st := utils.FromError(err)
			if st == codes.Internal {
				utils.ExtractLogger(ctx).Error(err.Error(), nil)
			}
			return nil, status.Error(st, msg)
		}

		return res, nil
	}
}
