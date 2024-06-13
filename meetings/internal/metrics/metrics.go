package metrics

import (
	"context"
	"github.com/alserov/restate/metrics/pkg/models"
)

type Producer interface {
	Produce(ctx context.Context, message models.Message, topic string)
}

type Metrics interface {
	Produce(ctx context.Context, message models.Message)
}

func NewMetrics(p Producer) *metrics {
	return &metrics{p}
}

type metrics struct {
	Producer
}

func (m *metrics) Produce(ctx context.Context, message models.Message) {
	m.Producer.Produce(ctx, message, models.MetricsTopic)
}
