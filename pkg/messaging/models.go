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

type Message struct {
	MessageID   format.MessageID   `bson:"_id"`
	SenderID    format.UserID      `bson:"sender_id"`
	ChannelID   format.ChannelID   `bson:"channel_id"`
	Content     string             `bson:"content"`
	ContentType format.ContentType `bson:"content_type"`
	CreatedAt   time.Time          `bson:"created_at"`
}
