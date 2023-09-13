package resolver

import "context"

type MessagesGetter interface {
	Messages(context.Context) ([]*Message, error)
	Next(context.Context) (*string, error)
}

type Messages struct {
	MessagesGetter
}

func NewMessages(getter MessagesGetter) *Messages {
	return &Messages{
		MessagesGetter: getter,
	}
}
