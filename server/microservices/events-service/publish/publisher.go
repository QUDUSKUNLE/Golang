package publish

import (
	"context"
	"encoding/json"
	"log"

	"github.com/QUDUSKUNLE/microservices/events-service/domain"
	"github.com/segmentio/kafka-go"
)

// KafkaBroker represents a Kafka writer wrapper
type KafkaPublisher struct {
	writer *kafka.Writer
}

// Define Kafka Publisher
func (broker *KafkaPublisher) Publish(ctx context.Context, topic string, event interface{}) error {
	message, _ := json.Marshal(event)
	if err := broker.writer.WriteMessages(ctx, kafka.Message{
		Topic: topic, Value: message,
	}); err != nil {
		log.Println("Error publishing message:", err)
		return err
	}
	log.Println("Message published successfully")
	return nil
}

func NewBroker(brokerAddress, topic string) domain.EventPorts {
	return &KafkaPublisher{writer: &kafka.Writer{
		Addr:     kafka.TCP(brokerAddress),
		Balancer: &kafka.LeastBytes{},
	}}
}
