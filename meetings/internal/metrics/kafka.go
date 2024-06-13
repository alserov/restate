package metrics

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/alserov/restate/meetings/internal/log"
	"github.com/alserov/restate/metrics/pkg/models"
)

func NewKafkaProducer(addr string) Producer {
	prodCfg := sarama.NewConfig()
	prodCfg.Producer.Partitioner = sarama.NewRandomPartitioner
	prodCfg.Producer.RequiredAcks = sarama.WaitForAll
	prodCfg.Producer.Return.Successes = true

	prod, err := sarama.NewAsyncProducer([]string{addr}, prodCfg)
	if err != nil {
		panic("failed to init producer: " + err.Error())
	}

	return &kafka{prod}
}

type kafka struct {
	sarama.AsyncProducer
}

func (k *kafka) Produce(ctx context.Context, m models.Message, topic string) {
	b, err := json.Marshal(m)
	if err != nil {
		log.FromCtx(ctx).Error("failed to unmarshal", log.WithData("error", err.Error()))
	}

	k.AsyncProducer.Input() <- &sarama.ProducerMessage{Value: sarama.StringEncoder(b), Topic: topic}
}
