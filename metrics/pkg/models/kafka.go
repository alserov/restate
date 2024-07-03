package models

type MetricType uint

type Message struct {
	Type MetricType
	Data []byte
}

const (
	TimePerRequest MetricType = iota
	RequestStatus
)
