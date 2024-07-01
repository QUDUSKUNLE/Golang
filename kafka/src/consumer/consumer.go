package consumer

import (
	"fmt"
	"os"
	"os/signal"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Consumer() {
	// Create a new Kafka consumer
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers":   "localhost:9092",
			"group.id":           "test-group",
			"auto.offset.reset":  "earliest",
	})
	if err != nil {
			fmt.Printf("Failed to create consumer: %s\n", err)
			return
	}
	defer c.Close()

	// Subscribe to the Kafka topic
	err = c.SubscribeTopics([]string{"Kafka"}, nil)
	if err != nil {
			fmt.Printf("Failed to subscribe to topic: %s\n", err)
			return
	}

	// Setup a channel to handle OS signals for graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	// Start consuming messages
	run := true
	for (run) {
			select {
			case sig := <-sigchan:
					fmt.Printf("Received signal %v: terminating\n", sig)
					run = false
			default:
					// Poll for Kafka messages
					ev := c.Poll(100)
					if ev == nil {
							continue
					}

					switch e := ev.(type) {
					case *kafka.Message:
							// Process the consumed message
							fmt.Printf("Received message from topic %s: %s\n", *e.TopicPartition.Topic, string(e.Value))
					case kafka.Error:
							// Handle Kafka errors
							fmt.Printf("Error: %v\n", e)
					}
			}
	}
}
