package auth

import "time"

type UserUpdatedEvent struct {
	User      User      `json:"user"`
	Timestamp time.Time `json:"timestamp"`
}

func (e UserUpdatedEvent) Route() string {
	return "auth:event.user.updated"
}
