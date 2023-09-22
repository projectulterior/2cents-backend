package users

import "time"

type UserUpdatedEvent struct {
	User      User
	Timestamp time.Time
}

func (e UserUpdatedEvent) Route() string {
	return "users:event.user.updated"
}
