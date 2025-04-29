// EventPorts defines the interface for event-driven communication in the system.
// It provides methods for publishing events to a specific topic and subscribing
// to events from a specific topic.
//
// Publish sends an event to the specified topic. It returns an error if the
// publishing fails.
//
// Subscribe registers a handler function to process events from the specified
// topic. The handler function receives the event data as a byte slice.
package domain

import "context"

type EventPorts interface {
	Publish(context context.Context, topic string, event interface{}) error
	// Subscribe(topic string, handler func(event []byte))
}
