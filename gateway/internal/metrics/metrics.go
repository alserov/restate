package metrics

import (
	"context"
	"time"
)

type Metrics interface {
	ObserveRequest(ctx context.Context, status int, dur time.Duration, name string)
}

func NewMetrics() Metrics {
	return &metrics{}
}

type metrics struct {
}

func (m metrics) ObserveRequest(ctx context.Context, status int, dur time.Duration, name string) {
	//TODO implement me
	panic("implement me")
}
