package users

import (
	"context"
	"time"
)

type UserUpdatedEvent struct {
	User      User
	Timestamp time.Time
}

func (s *Service) EmitUserUpdated(ctx context.Context, event UserUpdatedEvent) error {
	return s.UserUpdated.Publish(ctx, event)
}
