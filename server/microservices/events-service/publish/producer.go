package publish

import (
	"context"
	"encoding/json"
	"github.com/QUDUSKUNLE/microservices/events-service/domain"
	"github.com/segmentio/kafka-go"
)

// KafkaBroker represents a Kafka writer wrapper
type KafkaBroker struct {
	writer *kafka.Writer
}

// Define Kafka Publisher
func (broker *KafkaBroker) Publish(topic string, event interface{}) error {
	message, _ := json.Marshal(event)
	return broker.writer.WriteMessages(context.Background(), kafka.Message{
		Topic: topic, Value: message,
	})
}

func NewBroker(writer *kafka.Writer) domain.EventPorts {
	return &KafkaBroker{writer: writer}
}
