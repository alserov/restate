package grpc

import (
	"context"
	"github.com/alserov/restate/estate/internal/log"
	"github.com/alserov/restate/estate/internal/utils"
	"google.golang.org/grpc"
)

func WithRecover() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if e := recover(); err != nil {
				utils.ExtractLogger(ctx).Error(
					"panic recovery",
					log.WithData("error", e),
				)
			}
		}()

		res, err := handler(ctx, req)
		return res, err
	}
}
