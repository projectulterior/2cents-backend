package pubsub

import "context"

type Route string

type Message any

type Broker interface {
	Publisher(Route) Publisher
	Listener(Route) Listener
	Shutdown(context.Context) error
}

type Publisher interface {
	Publish(context.Context, Message) error
}

type Listener interface {
	Next(context.Context) (Message, error)
	Close(context.Context)
}
