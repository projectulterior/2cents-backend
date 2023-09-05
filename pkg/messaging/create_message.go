package messaging

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateMessageRequest struct {
	ChannelID   format.ChannelID
	SenderID    format.UserID
	Content     string
	ContentType format.ContentType
}

type CreateMessageResponse = Message

func (s *Service) CreateMessage(ctx context.Context, req CreateMessageRequest) (*CreateMessageResponse, error) {
	message := Message{
		MessageID:   format.NewMessageID(),
		ChannelID:   req.ChannelID,
		SenderID:    req.SenderID,
		Content:     req.Content,
		ContentType: req.ContentType,
		CreatedAt:   time.Now(),
	}

	_, err := s.Collection(MESSAGES_COLLECTION).
		InsertOne(ctx, message)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &message, nil
}
