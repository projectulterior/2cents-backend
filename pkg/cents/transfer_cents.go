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

type TransferCentsRequest struct {
	UserID format.UserID
	Amount int
}

type TransferCentsResponse = Cents

func (s *Service) TransferCents(ctx context.Context, req TransferCentsRequest) (*TransferCentsResponse, error) {
	inc := bson.M{}

	if req.Amount != 0 {
		inc["total"] = req.Amount
		if req.Amount < 0 {
			inc["given"] = (-req.Amount)
		}
		if req.Amount > 0 {
			inc["earned_cents"] = req.Amount
		}
	}

	var cents Cents

	err := s.Collection(CENTS_COLLECTION).
		FindOneAndUpdate(ctx,
			bson.M{"_id": req.UserID.String()},
			bson.M{
				"$inc": inc,
				"$currentDate": bson.M{
					"updated_at": true,
				},
			},
			options.FindOneAndUpdate().
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
