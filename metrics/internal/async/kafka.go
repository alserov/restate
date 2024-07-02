package async

import (
	"context"
	"fmt"
	kafkago "github.com/segmentio/kafka-go"
	"os"
)

var _ Consumer = &kafka{}

type kafka struct {
	consumer *kafkago.Reader

	topic string
}

func newKafka(addr string, topic string) *kafka {
	cfg := kafkago.ReaderConfig{
		Brokers: []string{addr},
		Topic:   topic,
		GroupID: "system_metrics_group",
	}

	cons := kafkago.NewReader(cfg)

	return &kafka{
		consumer: cons,
		topic:    topic,
	}
}

func (k *kafka) Consume(ctx context.Context) <-chan []byte {
	ch := make(chan []byte, 1)
	go func() {
		defer close(ch)
		defer k.consumer.Close()

		for {
			msg, err := k.consumer.ReadMessage(ctx)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading message: %v\n", err)
				break
			}

			ch <- msg.Value
		}
	}()

	return ch
}
