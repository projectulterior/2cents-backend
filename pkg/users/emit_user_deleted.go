package users

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
)

type UserDeletedEvent struct {
	UserID    format.UserID
	Timestamp time.Time
}

func (s *Service) EmitUserDeleted(ctx context.Context, event UserDeletedEvent) error {
	return s.UserDeleted.Publish(ctx, event)
}
