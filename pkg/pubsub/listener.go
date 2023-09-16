package pubsub

import (
	"context"
	"fmt"
)

type listener struct {
	router *router
	ch     <-chan Message
}

func (l *listener) Next(ctx context.Context) (Message, error) {
	select {
	case <-ctx.Done():
		return nil, context.Canceled
	case msg, ok := <-l.ch:
		if !ok {
			return nil, fmt.Errorf("unexpected closed channel")
		}
		return msg, nil
	}
}

func (l *listener) Close(ctx context.Context) {
	l.router.removeListener(l)
}
