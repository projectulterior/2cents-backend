package pubsub

import "context"

type Message interface {
	Route() string
}

type Publisher[M Message] interface {
	Publish(context.Context, M) error
}

type Listener[M Message] interface {
	Next(context.Context) (*M, error)
	Close(context.Context)
}

type Exchange[M Message] interface {
	Publisher() Publisher[M]
	Listener() Listener[M]
	Subscribe(func(context.Context, M) error)
	Shutdown(context.Context) error
}
