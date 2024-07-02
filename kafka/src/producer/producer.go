package producer

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	bootstrapServer = "bootstrap.servers"
	kafkaBroker = "localhost:9092"
	topic = "Kafka"
)

type Message struct {
	Key string `json:"key"`
	Value string `json:"value"`
}


func Producer() {
	kaf, err := kafka.NewProducer(&kafka.ConfigMap{bootstrapServer: kafkaBroker})
	if err != nil {
		fmt.Printf("Failed to create a producer %s\n", err.Error())
		return
	}
	// Close producer after use
	defer kaf.Close()

	// Define message to be sent
	message := Message{Key: "example_key", Value: "Hello, Kafka!"}

	serializedMessage, err := SerializeMessage(message)
	if err != nil {
		fmt.Printf("Failed to serialize message: %s\n", err)
		return
	}
	if err := ProduceMessage(kaf, topic, serializedMessage); err != nil {
		fmt.Printf("Failed to produce message: %s\n", err)
		return
	}
}


func SerializeMessage(message Message) ([]byte, error) {
	serialized, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize message: %w", err)
	}
	return serialized, nil
}

func ProduceMessage(producer *kafka.Producer, topic string, message []byte) error {
	// Create a new message producer
	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value: message,
	}

	// Produce the kafka message
	deliveryChan := make(chan kafka.Event)
	if err := producer.Produce(kafkaMessage, deliveryChan); err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}

	er := <-deliveryChan
	m := er.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return fmt.Errorf("delivery failed: %s", m.TopicPartition.Error)
	}
	close(deliveryChan)
	return nil
}
