// nolint
package workers

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/alserov/restate/metrics/internal/log"
	"os"
)

var _ Worker = &system{}

func NewSystemWorker() *system {
	consumer, err := sarama.NewConsumer([]string{os.Getenv("KAFKA_ADDR")}, sarama.NewConfig())
	if err != nil {
		panic("failed to init consumer: " + err.Error())
	}

	return &system{
		consumer: consumer,
	}
}

type system struct {
	consumer sarama.Consumer
}

func (s *system) Run(ctx context.Context) {
	partitions, _ := s.consumer.Partitions(MetricsTopic)
	pcons, err := s.consumer.ConsumePartition(MetricsTopic, partitions[0], sarama.OffsetNewest)
	if nil != err {
		panic("failed to consume: " + err.Error())
	}

	ch := pcons.Messages()
	l := log.FromCtx(ctx)

	for i := 0; i < 5; i++ {
		go func() {
			for msg := range ch {
				var m Message
				if err = json.Unmarshal(msg.Value, &m); err != nil {
					l.Error("failed to unmarshal", log.WithData("error", err.Error()))
				}

				switch m.Type {
				case TimePerRequest:
					// logic here
				}
			}
		}()
	}

	select {
	case <-ctx.Done():
		s.consumer.Close()
	}
}
