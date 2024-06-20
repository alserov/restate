package async

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/alserov/restate/meetings/internal/log"
	"github.com/alserov/restate/meetings/internal/middleware/wrappers"
)

var _ Producer = &kafka{}

func newKafka(addr string, topic string) *kafka {
	prodCfg := sarama.NewConfig()
	prodCfg.Producer.Partitioner = sarama.NewRandomPartitioner
	prodCfg.Producer.RequiredAcks = sarama.WaitForAll
	prodCfg.Producer.Return.Successes = true

	prod, err := sarama.NewAsyncProducer([]string{addr}, prodCfg)
	if err != nil {
		panic("failed to init producer: " + err.Error())
	}

	return &kafka{prod, topic}
}

type kafka struct {
	sarama.AsyncProducer

	topic string
}

func (k kafka) Produce(ctx context.Context, message any) {
	b, err := json.Marshal(message)
	if err != nil {
		wrappers.ExtractLogger(ctx).Error("failed to unmarshal", log.WithData("error", err.Error()))
	}

	k.AsyncProducer.Input() <- &sarama.ProducerMessage{Value: sarama.StringEncoder(b), Topic: k.topic}
}
