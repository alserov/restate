package async

import "context"

type Producer interface {
	Produce(ctx context.Context, message any)
}

type ProducerType int

const (
	Kafka ProducerType = iota
)

func NewProducer(t ProducerType, targetAddr, topic string) Producer {
	switch t {
	case Kafka:
		return newKafka(targetAddr, topic)
	}

	return nil
}
