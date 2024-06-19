package async

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/alserov/restate/metrics/pkg/models"
)

var _ Producer = &kafka{}

func newKafka() {
	prodCfg := sarama.NewConfig()
	prodCfg.Producer.Partitioner = sarama.NewRandomPartitioner
	prodCfg.Producer.RequiredAcks = sarama.WaitForAll
	prodCfg.Producer.Return.Successes = true

	prod, err := sarama.NewAsyncProducer([]string{addr}, prodCfg)
	if err != nil {
		panic("failed to init producer: " + err.Error())
	}
}

type kafka struct {
}

func (k kafka) Produce(ctx context.Context, message models.Message, topic string) {
	//TODO implement me
	panic("implement me")
}
