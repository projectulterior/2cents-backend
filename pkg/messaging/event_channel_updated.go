package messaging

import "time"

type ChannelUpdatedEvent struct {
	Channel   Channel
	Timestamp time.Time
}

func (e ChannelUpdatedEvent) Route() string {
	return "messging:event.channel.updated"
}
