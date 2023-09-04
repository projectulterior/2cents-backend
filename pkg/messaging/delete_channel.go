package messaging

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteChannelRequest struct {
	ChannelID format.ChannelID
}

type DeleteChannelResponse struct {
	ChannelID format.ChannelID
}

func (s *Service) DeleteChannel(ctx context.Context, req DeleteChannelRequest) (*DeleteChannelResponse, error) {
	_, err := s.Collection(CHANNELS_COLLECTION).
		DeleteOne(ctx, bson.M{"_id": req.ChannelID.String()})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &DeleteChannelResponse{
		ChannelID: req.ChannelID,
	}, nil
}
