package messaging

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	var channel Channel
	err := s.Collection(CHANNELS_COLLECTION).
		FindOne(ctx,
			bson.M{
				"_id":        req.ChannelID.String(),
				"member_ids": req.SenderID.String(),
			},
		).Decode(&channel)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, status.Error(codes.NotFound, err.Error())
	}

	now := time.Now()

	message := Message{
		MessageID:   format.NewMessageID(),
		ChannelID:   req.ChannelID,
		SenderID:    req.SenderID,
		Content:     req.Content,
		ContentType: req.ContentType,
		CreatedAt:   now,
	}

	_, err = s.Collection(MESSAGES_COLLECTION).
		InsertOne(ctx, message)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	s.ChannelUpdated.Publish(ctx, ChannelUpdatedEvent{
		Channel:   channel,
		Timestamp: now,
	})

	return &message, nil
}
