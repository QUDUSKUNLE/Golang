package consumers

import (
	"context"
	"encoding/json"
	"log"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	"github.com/segmentio/kafka-go"
)


func ConsumeCreatedUserEvent(brokerAddress, topic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic: topic,
	})

	defer reader.Close()

	for {
		// Read messages from the topic
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		// Process the event
		var event dto.UserCreatedEvent
		if err = json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("Error unmarshaling event: %v", err)
			continue
		}
		log.Printf("Processing UserCreatedEvent: UserID=%s, Email=%s",event.UserID, event.Email)
	}
}
