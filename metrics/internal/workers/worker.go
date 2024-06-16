package workers

import "context"

type Worker interface {
	Run(ctx context.Context)
}

func NewWorker(t WorkerType, brokerAddr string) Worker {
	switch t {
	case Business:
		return nil
	case System:
		return NewSystemWorker(brokerAddr)
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
