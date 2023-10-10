package cents

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetCentsRequest struct {
	UserID format.UserID
}

type GetCentsResponse = Cents

func (s *Service) GetCents(ctx context.Context, req GetCentsRequest) (*GetCentsResponse, error) {
	return s.getCents(ctx, req.UserID)
}

func (s *Service) getCents(ctx context.Context, userID format.UserID) (*Cents, error) {
	now := time.Now()

	var cents Cents
	err := s.Collection(CENTS_COLLECTION).
		FindOne(ctx, bson.M{"_id": userID.String()}).
		Decode(&cents)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return &Cents{
			UserID:    userID,
			Total:     0,
			Deposited: 0,
			Received:  0,
			Sent:      0,
			UpdatedAt: now,
		}, nil
	}

	return &cents, nil
}
