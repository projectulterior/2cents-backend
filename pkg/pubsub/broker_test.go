package pubsub_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/pubsub"
	"github.com/stretchr/testify/assert"
)

func TestBroker(t *testing.T) {
	broker := pubsub.NewBroker()

	route := pubsub.Route("basic")
	message := "hello"

	publisher := broker.Publisher(route)
	listener := broker.Listener(route)

	err := publisher.Publish(context.Background(), message)
	assert.NoError(t, err)

	msg, err := listener.Next(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, message, msg)
}
