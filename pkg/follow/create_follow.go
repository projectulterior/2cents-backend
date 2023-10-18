package follow

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/cents"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateFollowRequest struct {
	FollowerID format.UserID
	FolloweeID format.UserID
}

type CreateFollowResponse = Follow

func (s *Service) CreateFollow(ctx context.Context, req CreateFollowRequest) (*CreateFollowResponse, error) {
	now := time.Now()

	follow := Follow{
		FollowID:   format.NewFollowID(req.FollowerID, req.FolloweeID),
		FollowerID: req.FollowerID,
		FolloweeID: req.FolloweeID,
		CreatedAt:  now,
	}

	session, err := s.Database.Client().StartSession()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {

		_, err = s.Cents.TransferCents(ctx, cents.TransferCentsRequest{
			SenderID:   req.FollowerID,
			ReceiverID: req.FolloweeID,
			Amount:     1,
		})
		if err != nil {
			return nil, err
		}

		_, err = s.Collection(FOLLOW_COLLECTION).
			InsertOne(ctx, follow)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	return &follow, nil
}
