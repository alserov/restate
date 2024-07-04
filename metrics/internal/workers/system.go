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

func NewSystemWorker(consumer async.Consumer, colls *[]prometheus.Collector) *system {
	timePerRequest := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "system",
		Name:      "time_per_request",
	}, []string{"req_name"})

	requestStatus := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "system",
		Name:      "request_status",
	}, []string{"req_name", "status"})

	*colls = append(*colls, requestStatus, timePerRequest)

	return &system{
		timePerRequest: timePerRequest,
		requestStatus:  requestStatus,
		consumer:       consumer,
	}
}

type system struct {
	timePerRequest *prometheus.HistogramVec
	requestStatus  *prometheus.CounterVec

	consumer async.Consumer
}

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
					reqName, ok := m.Data["reqName"].(string)
					if !ok {
						continue
					}

					dur, ok := m.Data["time"].(float64)
					if !ok {
						continue
					}

					s.timePerRequest.With(prometheus.Labels{"req_name": reqName}).Observe(dur)
				case RequestStatus:
					reqName, ok := m.Data["reqName"].(string)
					if !ok {
						continue
					}

					status, ok := m.Data["status"].(float64)
					if !ok {
						continue
					}

					s.requestStatus.With(prometheus.Labels{"req_name": reqName, "status": strconv.Itoa(int(status))}).Inc()
				default:
					l.Error("invalid message type", log.WithData("type", m.Type))
				}
			}
		}()
	}
}
