package pubsub

import "context"

type publisher[M Message] struct {
	ch chan<- M
}

func (p *publisher[M]) Publish(ctx context.Context, msg M) error {
	p.ch <- msg
	return nil
}

func NewPublisher[M Message](ch chan<- M) Publisher[M] {
	return &publisher[M]{
		ch: ch,
	}
}
