package messaging

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
)

type CreateMessageRequest struct {
	ChannelID   format.ChannelID
	SenderID    format.UserID
	Content     string
	ContentType format.ContentType
}

type CreateMessageResponse = Message

func (s *Service) CreateMessage(ctx context.Context)
