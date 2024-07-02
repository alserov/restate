package metrics

import (
	"context"
	"encoding/json"
	"github.com/alserov/restate/gateway/internal/async"
	"github.com/alserov/restate/gateway/internal/utils"
	"time"
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

func (m *metrics) ObserveRequest(ctx context.Context, status int, dur time.Duration, name string) error {
	b, err := json.Marshal(nil)
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	m.Producer.Produce(ctx, b)

	return nil
}
