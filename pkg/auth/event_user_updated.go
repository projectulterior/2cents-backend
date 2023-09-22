package auth

import "time"

type UserUpdatedEvent struct {
	User      User
	Timestamp time.Time
}

func (e UserUpdatedEvent) Route() string {
	return "auth:event.user.updated"
}
