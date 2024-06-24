package metrics

import (
	"context"
	"time"
)

import (
	"github.com/alserov/restate/estate/internal/async"
)

type Metrics interface {
	ObserveRequest(ctx context.Context, status int, dur time.Duration, name string) error
}

func NewMetrics(p async.Producer) *metrics {
	return &metrics{p}
}

type metrics struct {
	async.Producer
}

type (
	TimePerRequestData struct {
		ReqName string        `json:"reqName"`
		Time    time.Duration `json:"time"`
	}

	RequestStatusData struct {
		ReqName string `json:"reqName"`
		Status  int    `json:"status"`
	}
)

func (m *metrics) ObserveRequest(ctx context.Context, status int, dur time.Duration, name string) error {
	m.Producer.Produce(ctx, TimePerRequestData{ReqName: name, Time: dur})
	m.Producer.Produce(ctx, RequestStatusData{ReqName: name, Status: status})
	return nil
}
