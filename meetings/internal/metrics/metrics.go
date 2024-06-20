package metrics

import (
	"context"
	"github.com/alserov/restate/meetings/internal/async"
	"time"
)

type Metrics interface {
	ObserveRequest(ctx context.Context, status int, dur time.Duration, key string) error
}

func NewMetrics(p async.Producer) *metrics {
	return &metrics{p}
}

type metrics struct {
	async.Producer
}

type (
	TimePerRequestData struct {
		Key  string        `json:"reqName"`
		Time time.Duration `json:"time"`
	}

	RequestStatusData struct {
		Key    string `json:"reqName"`
		Status int    `json:"status"`
	}
)

func (m *metrics) ObserveRequest(ctx context.Context, status int, dur time.Duration, key string) error {
	m.Producer.Produce(ctx, TimePerRequestData{Key: key, Time: dur})
	m.Producer.Produce(ctx, RequestStatusData{Key: key, Status: status})
	return nil
}
