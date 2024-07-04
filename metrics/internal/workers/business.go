package workers

import (
	"context"
	"encoding/json"
	"github.com/alserov/restate/metrics/internal/async"
	"github.com/alserov/restate/metrics/internal/log"
	"github.com/prometheus/client_golang/prometheus"
)

var _ Worker = &business{}

func NewBusinessWorker(cons async.Consumer, colls *[]prometheus.Collector) Worker {
	estateMeetings := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "business",
		Name:      "estate_meetings",
	}, []string{"estateID"})

	*colls = append(*colls, estateMeetings)

	return &business{
		estateMeetings: estateMeetings,

		consumer: cons,
	}
}

const (
	EstateMeetings = iota
)

type business struct {
	estateMeetings *prometheus.CounterVec

	consumer async.Consumer
}

func (b *business) Run(ctx context.Context, workersAmount int) {
	l := log.FromCtx(ctx)

	for i := 0; i < workersAmount; i++ {
		go func() {
			for msg := range b.consumer.Consume(ctx) {
				var m Message
				if err := json.Unmarshal(msg, &m); err != nil {
					l.Error("failed to unmarshal", log.WithData("error", err.Error()))
				}

				switch m.Type {
				case EstateMeetings:
					estateID, ok := m.Data["estateID"].(string)
					if !ok {
						continue
					}

					b.estateMeetings.With(prometheus.Labels{"estateID": estateID}).Inc()
				default:
					l.Error("invalid message type", log.WithData("type", m.Type))
				}
			}
		}()
	}
}
