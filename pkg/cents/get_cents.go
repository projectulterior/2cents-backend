package cents

import (
	"context"

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
	var cents Cents

	err := s.Collection(CENTS_COLLECTION).
		FindOne(ctx, bson.M{"_id": req.UserID.String()}).
		Decode(&cents)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &cents, nil
}
