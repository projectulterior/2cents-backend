package pubsub

import (
	"context"
	"fmt"
)

type listener struct {
	ch <-chan Message
}

func (l *listener) Next(ctx context.Context) (Message, error) {
	msg, ok := <-l.ch
	if !ok {
		return nil, fmt.Errorf("unexpected closed channel")
	}

	return msg, nil
}
