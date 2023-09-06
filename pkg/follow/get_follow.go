package follow

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetFollowRequest struct {
	FollowID format.FollowID
}

type GetFollowResponse = Follow

func (s *Service) GetFollow(ctx context.Context, req GetFollowRequest) (*GetFollowResponse, error) {
	return s.getFollow(ctx, req.FollowID)
}

func (s *Service) getFollow(ctx context.Context, followID format.FollowID) (*GetFollowResponse, error) {
	var follow Follow
	err := s.Collection(FOLLOW_COLLECTION).
		FindOne(ctx, bson.M{"_id": followID.String()}).
		Decode(&follow)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.NotFound, err.Error())

	}
	return &follow, nil
}
