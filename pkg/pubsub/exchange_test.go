package pubsub_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/pubsub"
	"github.com/stretchr/testify/assert"
)

type TestMessage struct {
	Content string
}

func (m TestMessage) Route() string {
	return "test"
}

func TestExchange(t *testing.T) {
	exchange := pubsub.NewExchange[TestMessage]()

	message := "hello"

	publisher := exchange.Publisher()
	listener := exchange.Listener()

	err := publisher.Publish(context.Background(), TestMessage{
		Content: message,
	})
	assert.NoError(t, err)

	msg, err := listener.Next(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, message, msg.Content)
}
