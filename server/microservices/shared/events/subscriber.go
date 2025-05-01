// ConsumeCreatedUserEvent listens to a Kafka topic for user creation events and processes them.
//
// Parameters:
//   - d (*db.Queries): A database queries object for interacting with the database.
//   - brokerAddress (string): The address of the Kafka broker.
//   - topic (string): The Kafka topic to consume messages from.
//   - groupID (string): The consumer group ID for Kafka.
//
// Behavior:
//   - Initializes a Kafka reader with the provided broker address, topic, and group ID.
//   - Continuously reads messages from the specified Kafka topic.
//   - Processes messages based on the topic type. Specifically, it handles user creation events
//     by unmarshaling the event data, logging the event, and creating an associated organization
//     in the database using the organization service.
//   - Logs errors encountered during message reading, unmarshaling, or organization creation.
//
// Notes:
//   - The function assumes the existence of a constant `USER_EVENTS` for identifying user-related events.
//   - The organization service must be properly initialized with a concrete implementation of `db.Queries`.
//   - The function runs indefinitely and should be executed in a separate goroutine or managed appropriately.
//
// Example Usage:
//
//	go consumers.ConsumeCreatedUserEvent(dbQueries, "localhost:9092", "user-events", "group-id")
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
