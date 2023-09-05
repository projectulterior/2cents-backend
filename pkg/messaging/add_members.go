package messaging

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AddMembersRequest struct {
	ChannelID format.ChannelID
	MemberID  format.UserID
	MemberIDs []format.UserID
}

type AddMembersResponse = Channel

func (s *Service) AddMembers(ctx context.Context, req AddMembersRequest) (*AddMembersResponse, error) {
	var channel Channel
	err := s.Collection(CHANNELS_COLLECTION).
		FindOneAndUpdate(ctx,
			bson.M{
				"_id":        req.ChannelID.String(),
				"member_ids": req.MemberID.String(),
			},
			bson.M{
				"$addToSet": bson.M{
					"member_ids": bson.M{
						"$each": req.MemberIDs,
					},
				},
				"$currentDate": bson.M{
					"updated_at": true,
				},
			},
			options.FindOneAndUpdate().
				SetReturnDocument(options.After),
		).Decode(&channel)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &channel, nil
}
