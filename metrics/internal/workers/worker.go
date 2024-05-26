package workers

import "context"

type Worker interface {
	Run(ctx context.Context)
}

func NewWorker(t WorkerType) Worker {
	switch t {
	case Business:
		return nil
	case System:
		return NewSystemWorker()
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
