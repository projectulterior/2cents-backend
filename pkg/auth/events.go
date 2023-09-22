package auth

import (
	"time"

	"github.com/projectulterior/2cents-backend/pkg/pubsub"
)

const (
	USER_UPDATED_EVENT pubsub.Route = "event.user.updated"
)

type UserUpdatedEvent struct {
	User      User
	Timestamp time.Time
}
