package metrics

import (
	"context"
	"github.com/alserov/restate/meetings/internal/async"
	"time"
)

type Metrics interface {
	ObserveRequest(ctx context.Context, status int, dur time.Duration, key string) error
	IncEstateMeeting(ctx context.Context, estateID string) error
}

func NewMetrics(p async.Producer) *metrics {
	return &metrics{p}
}

type metrics struct {
	async.Producer
}

type (
	Message struct {
		Type uint
		Data any
	}

	TimePerRequestData struct {
		ReqName string        `json:"reqName"`
		Time    time.Duration `json:"time"`
	}

	EstateMeeting struct {
		EstateID string `json:"estateID"`
	}

	RequestStatusData struct {
		ReqName string `json:"reqName"`
		Status  uint   `json:"status"`
	}
)

const (
	timePerRequest = iota
	requestStatus
)

const (
	estateMeetingCount = iota
)

func (m *metrics) ObserveRequest(ctx context.Context, status int, dur time.Duration, key string) error {
	m.Producer.Produce(ctx, Message{Type: timePerRequest, Data: TimePerRequestData{ReqName: key, Time: dur}})
	m.Producer.Produce(ctx, Message{Type: requestStatus, Data: RequestStatusData{ReqName: key, Status: uint(status)}})
	return nil
}

func (m *metrics) IncEstateMeeting(ctx context.Context, estateID string) error {
	m.Producer.Produce(ctx, Message{Type: estateMeetingCount, Data: EstateMeeting{EstateID: estateID}})
	return nil
}
