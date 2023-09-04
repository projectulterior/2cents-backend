package messaging

import (
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
)

const (
	CHANNELS_COLLECTION = "channels"
	MESSAGES_COLLECTION = "messages"
)

type Channel struct {
	ChannelID format.ChannelID `bson:"_id"`
	MemberIDs []format.UserID  `bson:"member_ids"`
	CreatedAt time.Time        `bson:"created_at"`
}