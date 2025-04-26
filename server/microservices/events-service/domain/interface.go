package domain

type EventPorts interface {
	Publish(topic string, event interface{}) error
	Subscribe(topic string, handler func(event []byte))
}
