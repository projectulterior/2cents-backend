package messaging

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetChannelByMembersRequest struct {
	MemberIDs []format.UserID
}

type GetChannelByMembersResponse = Channel

func (s *Service) GetChannelByMembers(ctx context.Context, req GetChannelByMembersRequest) (*GetChannelByMembersResponse, error) {
	var channel Channel
	err := s.Collection(CHANNELS_COLLECTION).
		FindOne(ctx, bson.M{
			"member_ids": bson.M{
				"$all": req.MemberIDs,
			},
		}).
		Decode(&channel)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &channel, nil
}
