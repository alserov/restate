package grpc

import (
	"context"
	"github.com/alserov/restate/estate/internal/metrics"
	"github.com/alserov/restate/estate/internal/utils"
	"google.golang.org/grpc"
	"time"
)

func WithRequestObserver(metr metrics.Metrics) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		start := time.Now()

		res, err := handler(ctx, req)

		_, st := utils.FromError(err)

		_ = metr.ObserveRequest(ctx, int(st), time.Since(start), "")

		return res, err
	}
}
