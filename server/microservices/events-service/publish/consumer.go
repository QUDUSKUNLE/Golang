package publish

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func (broker *KafkaBroker) Subscribe(topic string, handler func(event []byte)) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
	})
	for {
		msg, _ := reader.ReadMessage(context.Background())
		handler(msg.Value)
	}
}
