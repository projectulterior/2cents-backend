package messaging

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetMessageRequest struct {
	MessageID format.MessageID
}

type GetMessageResponse = Message

func (s *Service) GetMessage(ctx context.Context, req GetMessageRequest) (*GetMessageResponse, error) {
	var message Message
	err := s.Collection(MESSAGES_COLLECTION).
		FindOne(ctx, bson.M{"_id": req.MessageID.String()}).
		Decode(&message)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &message, nil
}
