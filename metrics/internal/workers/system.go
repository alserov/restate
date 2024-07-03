// nolint
package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alserov/restate/metrics/internal/async"
	"github.com/alserov/restate/metrics/internal/log"
	"github.com/alserov/restate/metrics/pkg/models"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

var _ Worker = &system{}

func NewSystemWorker(consumer async.Consumer) *system {
	return &system{
		consumer: consumer,
	}
}

var (
	timePerRequest = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "system",
		Name:      "time_per_request",
	}, []string{"req_name"})

	requestStatus = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "system",
		Name:      "request_status",
	}, []string{"req_name", "status"})
)

type (
	TimePerRequestData struct {
		ReqName string        `json:"reqName"`
		Time    time.Duration `json:"time"`
	}

	RequestStatusData struct {
		ReqName string `json:"reqName"`
		Status  uint   `json:"status"`
	}
)

type system struct {
	consumer async.Consumer
}

func (s *system) Run(ctx context.Context, workersAmount int) {
	l := log.FromCtx(ctx)

	for i := 0; i < workersAmount; i++ {
		go func() {
			for msg := range s.consumer.Consume(ctx) {
				var m models.Message
				if err := json.Unmarshal(msg, &m); err != nil {
					l.Error("failed to unmarshal", log.WithData("error", err.Error()))
				}

				fmt.Println(m)

				switch m.Type {
				case models.TimePerRequest:
					var data TimePerRequestData
					if err := json.Unmarshal(m.Data, &data); err != nil {
						l.Error("failed to unmarshal", log.WithData("error", err.Error()))
					}

					timePerRequest.With(prometheus.Labels{"req_name": data.ReqName}).Observe(float64(data.Time.Milliseconds()))
				case models.RequestStatus:
					var data RequestStatusData
					if err := json.Unmarshal(m.Data, &data); err != nil {
						l.Error("failed to unmarshal", log.WithData("error", err.Error()))
					}

					requestStatus.With(prometheus.Labels{"req_name": data.ReqName, "status": strconv.Itoa(int(data.Status))}).Inc()
				default:
					l.Error("invalid message type", log.WithData("type", m.Type))
				}
			}
		}()
	}
}
