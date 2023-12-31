package cents

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateCentsRequest struct {
	UserID format.UserID
	Amount int
}

type UpdateCentsResponse = Cents

func (s *Service) UpdateCents(ctx context.Context, req UpdateCentsRequest) (*UpdateCentsResponse, error) {
	inc := bson.M{}

	if req.Amount == 0 {
		return s.getCents(ctx, req.UserID)
	}

	inc["total"] = req.Amount
	if req.Amount > 0 {
		inc["deposited"] = req.Amount
	}

	var cents Cents
	err := s.Collection(CENTS_COLLECTION).
		FindOneAndUpdate(ctx,
			bson.M{"_id": req.UserID.String()},
			bson.M{
				"$inc": inc,
				"$setOnInsert": bson.M{
					"received": 0,
					"sent":     0,
				},
				"$currentDate": bson.M{
					"updated_at": true,
				},
			},
			options.FindOneAndUpdate().
				SetUpsert(true).
				SetReturnDocument(options.After),
		).Decode(&cents)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &cents, nil
}
