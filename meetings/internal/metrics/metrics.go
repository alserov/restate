package metrics

import (
	"context"
	"github.com/alserov/restate/gateway/internal/async"
	"github.com/alserov/restate/metrics/pkg/models"
)

type Metrics interface {
	Produce(ctx context.Context, message models.Message)
}

func NewMetrics(p async.Producer) *metrics {
	return &metrics{p}
}

type metrics struct {
	async.Producer
}

func (m *metrics) Produce(ctx context.Context, message models.Message) {
	m.Producer.Produce(ctx, message, models.MetricsTopic)
}
