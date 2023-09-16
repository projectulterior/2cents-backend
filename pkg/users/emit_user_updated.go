package users

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/pubsub"
)

const USER_UPDATED_EVENT pubsub.Route = "event.user.updated"

type UserUpdatedEvent struct {
	User      User
	Timestamp time.Time
}

func (s *Service) EmitUserUpdated(ctx context.Context, event UserUpdatedEvent) error {
	return s.UserUpdated.Publish(ctx, event)
}
