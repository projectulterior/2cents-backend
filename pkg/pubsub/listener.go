package pubsub

import (
	"context"
	"fmt"
)

type listener[M Message] struct {
	ex *exchange[M]
	ch <-chan M
}

func (l *listener[M]) Next(ctx context.Context) (*M, error) {
	select {
	case <-ctx.Done():
		return nil, context.Canceled
	case msg, ok := <-l.ch:
		if !ok {
			return nil, fmt.Errorf("unexpected closed channel")
		}
		return &msg, nil
	}
}

func (l *listener[M]) Close(ctx context.Context) {
	l.ex.removeListener(l)
}
