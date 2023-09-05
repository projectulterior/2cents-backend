package messaging

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateMessageRequest struct {
	MessageID   format.MessageID
	SenderID    format.UserID
	Content     *string
	ContentType *format.ContentType
}

type UpdateMessageResponse = Message

func (s *Service) UpdateMessage(ctx context.Context, req UpdateMessageRequest) (*UpdateMessageResponse, error) {
	var message Message

	err := s.Collection(MESSAGES_COLLECTION).
		FindOneAndUpdate(ctx,
			bson.M{
				"_id":       req.MessageID.String(),
				"sender_id": req.SenderID.String(),
			},
			bson.M{
				"$set": bson.M{
					"content":      req.Content,
					"content_type": req.ContentType,
				},
				"$currentDate": bson.M{
					"updated_at": true,
				},
			},
			options.FindOneAndUpdate().
				SetReturnDocument(options.After),
		).Decode(&message)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &message, nil
}
