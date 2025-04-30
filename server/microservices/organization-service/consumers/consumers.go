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
//   go consumers.ConsumeCreatedUserEvent(dbQueries, "localhost:9092", "user-events", "group-id")
package consumers

import (
	"context"
	"encoding/json"
	"log"

	organizationService "github.com/QUDUSKUNLE/microservices/organization-service/adapters/organizationcase"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	"github.com/segmentio/kafka-go"
)

func ConsumeCreatedUserEvent(d *db.Queries, brokerAddress, topic, groupID string) {
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
		switch topic {
		case USER_EVENTS:
			var user dto.UserCreatedEvent
			if err = json.Unmarshal(msg.Value, &user); err != nil {
				log.Printf("Error unmarshaling event: %v", err)
				continue
			}
			log.Printf("Processing OrganizationCreatedEvent: UserID=%s, Email=%s", user.UserID, user.Email)
			organizationService := organizationService.InitOrganizationServer(d)
			// Ensure this is initialized with a concrete implementation
			organization, err := organizationService.CreateOrganization(context.Background(), dto.OrganizationDto{UserID: user.UserID})
			if err != nil {
				log.Printf("Error creating organization: %v", err)
				continue
			}
			log.Printf("Organization created successfully: %v", organization.ID)
		default:
			log.Printf("Unknown topic: %s", topic)
		}
	}
}
