package async

import (
	"context"
	"github.com/alserov/restate/metrics/pkg/models"
)

type Producer interface {
	Produce(ctx context.Context, message models.Message, topic string)
}

type ProducerType int

const (
	Kafka ProducerType = iota
)

func NewProducer(t ProducerType, targetAddr string) Producer {
	switch t {
	case Kafka:
		return newKafka(targetAddr)
	}

	return nil
}
