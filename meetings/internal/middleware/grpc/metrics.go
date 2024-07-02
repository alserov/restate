package grpc

import (
	"context"
	"github.com/alserov/restate/meetings/internal/metrics"
	"github.com/alserov/restate/meetings/internal/utils"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

func WithRequestObserver(metr metrics.Metrics) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		start := time.Now()

		res, err := handler(ctx, req)

		_ = metr.ObserveRequest(ctx, http.StatusOK, time.Since(start), utils.ExtractIdempotencyKey(ctx))

		return res, err
	}
}
