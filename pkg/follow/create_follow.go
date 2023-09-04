package follow

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
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

	_, err := s.Collection(FOLLOW_COLLECTION).
		InsertOne(ctx, follow)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &follow, nil
}
