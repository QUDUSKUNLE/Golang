// ConsumeCreatedUserEvent listens to a Kafka topic for user creation events
package events

import (
	"context"
	"log"
	"github.com/segmentio/kafka-go"
)

func EventsSubscriber(brokerAddress, topic, groupID string, processEvent func(context.Context, []byte) error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{brokerAddress},
		Topic:          topic,
		GroupID:        groupID,
		CommitInterval: 0,
		StartOffset:    kafka.FirstOffset,
	})

	for {
		// Read messages from the topic
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		// Process the event
		if err := processEvent(context.Background(), msg.Value); err != nil {
			log.Printf("Error processing event: %v", err)
			continue
		}
	}
}
