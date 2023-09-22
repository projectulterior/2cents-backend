package users

import (
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
)

type UserDeletedEvent struct {
	UserID    format.UserID
	Timestamp time.Time
}

func (e UserDeletedEvent) Route() string {
	return "users:event.user.deleted"
}
