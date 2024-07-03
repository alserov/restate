// nolint
package workers

import (
	"context"
	"encoding/json"
	"github.com/alserov/restate/metrics/internal/async"
	"github.com/alserov/restate/metrics/internal/log"
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
	MetricType uint
	Message    struct {
		Type MetricType
		Data map[string]any
	}

	TimePerRequestData struct {
		ReqName string        `json:"reqName"`
		Time    time.Duration `json:"time"`
	}

	RequestStatusData struct {
		ReqName string `json:"reqName"`
		Status  uint   `json:"status"`
	}
)

const (
	TimePerRequest MetricType = iota
	RequestStatus
)

type system struct {
	consumer async.Consumer
}

func (s *system) Run(ctx context.Context, workersAmount int) {
	l := log.FromCtx(ctx)

	for i := 0; i < workersAmount; i++ {
		go func() {
			for msg := range s.consumer.Consume(ctx) {
				var m Message
				if err := json.Unmarshal(msg, &m); err != nil {
					l.Error("failed to unmarshal", log.WithData("error", err.Error()))
				}

				switch m.Type {
				case TimePerRequest:
					timePerRequest.With(prometheus.Labels{"req_name": m.Data["reqName"].(string)}).Observe(m.Data["time"].(float64))
				case RequestStatus:
					requestStatus.With(prometheus.Labels{"req_name": m.Data["reqName"].(string), "status": strconv.Itoa(int(m.Data["status"].(float64)))}).Inc()
				default:
					l.Error("invalid message type", log.WithData("type", m.Type))
				}
			}
		}()
	}
}
