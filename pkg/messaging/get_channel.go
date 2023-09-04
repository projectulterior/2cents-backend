package messaging

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetChannelRequest struct {
	ChannelID format.ChannelID
}

type GetChannelResponse = Channel

func (s *Service) GetChannel(ctx context.Context, req GetChannelRequest) (*GetChannelResponse, error) {
	var channel Channel
	err := s.Collection(CHANNELS_COLLECTION).
		FindOne(ctx, bson.M{"_id": req.ChannelID.String()}).
		Decode(&channel)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &channel, nil
}
