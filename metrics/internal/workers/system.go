// nolint
package workers

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/alserov/restate/metrics/internal/log"
	"github.com/alserov/restate/metrics/pkg/models"
	"github.com/prometheus/client_golang/prometheus"
	"os"
	"strconv"
	"time"
)

var _ Worker = &system{}

func NewSystemWorker() *system {
	consumer, err := sarama.NewConsumer([]string{os.Getenv("KAFKA_ADDR")}, sarama.NewConfig())
	if err != nil {
		panic("failed to init consumer: " + err.Error())
	}

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
	consumer sarama.Consumer
}

func (s *system) Run(ctx context.Context) {
	partitions, _ := s.consumer.Partitions(models.MetricsTopic)

	pcons, err := s.consumer.ConsumePartition(models.MetricsTopic, partitions[0], sarama.OffsetNewest)
	if nil != err {
		panic("failed to consume: " + err.Error())
	}

	ch := pcons.Messages()
	l := log.FromCtx(ctx)

	for i := 0; i < 5; i++ {
		go func() {
			for msg := range ch {
				var m models.Message
				if err = json.Unmarshal(msg.Value, &m); err != nil {
					l.Error("failed to unmarshal", log.WithData("error", err.Error()))
				}

				switch m.Type {
				case models.TimePerRequest:
					var data TimePerRequestData
					if err = json.Unmarshal(m.Data, &data); err != nil {
						l.Error("failed to unmarshal", log.WithData("error", err.Error()))
					}

					timePerRequest.With(prometheus.Labels{"req_name": data.ReqName}).Observe(float64(data.Time.Milliseconds()))
				case models.RequestStatus:
					var data RequestStatusData
					if err = json.Unmarshal(m.Data, &data); err != nil {
						l.Error("failed to unmarshal", log.WithData("error", err.Error()))
					}

					requestStatus.With(prometheus.Labels{"req_name": data.ReqName, "status": strconv.Itoa(int(data.Status))}).Inc()
				default:
					l.Error("invalid message type", log.WithData("type", m.Type))
				}
			}
		}()
	}

	select {
	case <-ctx.Done():
		s.consumer.Close()
	}
}
