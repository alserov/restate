package workers

import (
	"context"
	"github.com/alserov/restate/metrics/internal/async"
	"github.com/prometheus/client_golang/prometheus"
)

type Worker interface {
	Metrics() []prometheus.Collector
	Run(ctx context.Context, workersAmount int)
}

func NewWorker(t WorkerType, consumer async.Consumer, colls *[]prometheus.Collector) Worker {
	switch t {
	case Business:
		return nil
	case System:
		return NewSystemWorker(consumer, colls)
	default:
		panic("invalid worker type")
	}
}

type (
	WorkerType uint
)

const (
	System WorkerType = iota
	Business
)
