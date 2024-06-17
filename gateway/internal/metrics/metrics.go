package metrics

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/alserov/restate/gateway/internal/utils"
	"time"
)

type Metrics interface {
	ObserveRequest(ctx context.Context, status int, dur time.Duration, name string) error
}

var _ Metrics = &metrics{}

func NewMetrics(addr string) *metrics {
	cfg := sarama.NewConfig()

	prod, err := sarama.NewAsyncProducer([]string{addr}, cfg)
	if err != nil {
		panic("failed to init metrics: " + err.Error())
	}

	return &metrics{p: prod}
}

type metrics struct {
	p sarama.AsyncProducer
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

// ObserveRequest sends kafka message to broker, observes TimePerRequest and RequestStatus metrics
func (m metrics) ObserveRequest(ctx context.Context, status int, dur time.Duration, name string) error {
	timePerReq := TimePerRequestData{
		ReqName: name,
		Time:    dur,
	}

	b, err := json.Marshal(timePerReq)
	if err != nil {
		return utils.NewError(fmt.Sprintf("failed to marshal data: %v", err), utils.Internal)
	}

	m.p.Input() <- &sarama.ProducerMessage{Value: sarama.StringEncoder(b)}

	// ====================

	statusReq := RequestStatusData{
		ReqName: name,
		Status:  status,
	}

	b, err = json.Marshal(statusReq)
	if err != nil {
		return utils.NewError(fmt.Sprintf("failed to marshal data: %v", err), utils.Internal)
	}

	m.p.Input() <- &sarama.ProducerMessage{Value: sarama.StringEncoder(b)}

	return nil
}
