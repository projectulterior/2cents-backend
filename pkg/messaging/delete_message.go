package messaging

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteMessageRequest struct {
	MessageID format.MessageID
	SenderID  format.UserID
}

type DeleteMessageResponse struct {
	MessageID format.MessageID
}

func (s *Service) DeleteMessage(ctx context.Context, req DeleteMessageRequest) (*DeleteMessageResponse, error) {
	_, err := s.Collection(MESSAGES_COLLECTION).
		DeleteOne(ctx, bson.M{
			"_id":       req.MessageID.String(),
			"sender_id": req.SenderID.String(),
		})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &DeleteMessageResponse{
		MessageID: req.MessageID,
	}, nil
}
