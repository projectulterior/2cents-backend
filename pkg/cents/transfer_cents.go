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
	SenderID   format.UserID
	ReceiverID format.UserID
	Amount     int
}

type TransferCentsResponse = Cents // sender's cents

func (s *Service) TransferCents(ctx context.Context, req TransferCentsRequest) (*TransferCentsResponse, error) {
	if req.Amount < 1 {
		return nil, status.Error(codes.InvalidArgument, "amount must be greater than 0")
	}

	session, err := s.Database.Client().StartSession()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer session.EndSession(ctx)

	var cents Cents

	// TODO(cp): move to create_posts/likes/etc level to hold the transaction context throughout the transfer process
	_, err = session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
		/*
			- Check if sender has sufficient funds
				- _id = senderID, cents.Total >= amount
			- Update Sender's document by:
				total = total - amount
				sent += amount
				updatedAt = now
			- Update Receiver's document by:
				total += amount
				received += amount
				updatedAt = now
		*/

		err := s.Collection(CENTS_COLLECTION).
			FindOneAndUpdate(ctx,
				bson.M{
					"_id":   req.SenderID.String(),
					"total": bson.M{"$gte": req.Amount},
				},
				bson.M{
					"$inc": bson.M{
						"total": (-req.Amount),
						"sent":  req.Amount,
					},
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

			// sender doc not found
			return nil, status.Error(codes.FailedPrecondition, "insufficient funds")
		}

		err = s.Collection(CENTS_COLLECTION).FindOneAndUpdate(ctx,
			bson.M{
				"_id": req.ReceiverID.String(),
			},
			bson.M{
				"$inc": bson.M{
					"total":    req.Amount,
					"received": req.Amount,
				},
				"$currentDate": bson.M{
					"updated_at": true,
				},
			},
			options.FindOneAndUpdate().
				SetReturnDocument(options.After).
				SetUpsert(true),
		).Err()

		return nil, err
	})
	if err != nil {
		return nil, err
	}

	return &cents, nil
}
