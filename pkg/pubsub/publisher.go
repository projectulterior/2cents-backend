package pubsub

import "context"

type publisher struct {
	ch chan<- Message
}

func (p *publisher) Publish(ctx context.Context, msg Message) error {
	p.ch <- msg
	return nil
}
