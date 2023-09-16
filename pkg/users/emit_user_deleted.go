package users

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/pubsub"
)

const USER_DELETED_EVENT pubsub.Route = "event.user.deleted"

type UserDeletedEvent struct {
	UserID    format.UserID
	Timestamp time.Time
}

func (s *Service) EmitUserDeleted(ctx context.Context, event UserDeletedEvent) error {
	return s.UserDeleted.Publish(ctx, event)
}
