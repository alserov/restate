package models

type MetricType uint

type Message struct {
	Type MetricType
	Data any
}

const (
	TimePerRequest MetricType = iota
	RequestStatus
)
