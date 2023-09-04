package follow

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteFollowRequest struct {
	FollowID format.FollowID
}

type DeleteFollowResponse struct {
	FollowID format.FollowID
}

func (s *Service) DeleteFollow(ctx context.Context, req DeleteFollowRequest) (*DeleteFollowResponse, error) {
	_, err := s.Collection(FOLLOW_COLLECTION).
		DeleteOne(ctx, bson.M{"_id": req.FollowID.String()})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &DeleteFollowResponse{
		FollowID: req.FollowID,
	}, nil
}
