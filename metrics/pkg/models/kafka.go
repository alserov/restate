package models

type MetricType uint

type Message struct {
	Type MetricType
	Data []byte
}

const (
	MetricsTopic = "metrics"

	TimePerRequest MetricType = iota
	RequestStatus
)
