package metrics

import (
	"context"
	"time"
)

type Metrics interface {
	TimePerRequest(ctx context.Context, duration time.Duration, handlerName string)
}

var _ Metrics = &metrics{}

func NewMetrics() *metrics {
	return &metrics{}
}

type metrics struct {
}

func (m *metrics) TimePerRequest(ctx context.Context, duration time.Duration, handlerName string) {
	//TODO implement me
	panic("implement me")
}
