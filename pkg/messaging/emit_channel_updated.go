package messaging

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/pubsub"
)

const CHANNEL_UPDATED_EVENT pubsub.Route = "event.channel.updated"

type ChannelUpdatedEvent struct {
	Channel   Channel
	Timestamp time.Time
}

func (s *Service) EmitChannelUpdated(ctx context.Context, event ChannelUpdatedEvent) error {
	return s.ChannelUpdated.Publish(ctx, event)
}
