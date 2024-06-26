package models

type MetricType uint

type Message struct {
	Type MetricType
	Data []byte
}

const (
	TopicMetrics = "metrics"

	TimePerRequest MetricType = iota
	RequestStatus
)
