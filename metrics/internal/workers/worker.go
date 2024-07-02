package workers

import (
	"context"
	"github.com/alserov/restate/metrics/internal/async"
)

type Worker interface {
	Run(ctx context.Context, workersAmount int)
}

func NewWorker(t WorkerType, consumer async.Consumer) Worker {
	switch t {
	case Business:
		return nil
	case System:
		return NewSystemWorker(consumer)
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
