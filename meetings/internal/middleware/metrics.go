package middleware

import (
	"context"
	"github.com/alserov/restate/meetings/internal/metrics"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

func WithRequestObserver(metr metrics.Metrics) grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		start := time.Now()

		res, err := handler(ctx, req)

		_ = metr.ObserveRequest(ctx, http.StatusOK, time.Since(start), ExtractIdempotencyKey(ctx))

		return res, err
	})
}
