package messaging

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateChannelRequest struct {
	MemberIDs []format.UserID
}

type CreateChannelResponse = Channel

func (s *Service) CreateChannel(ctx context.Context, req CreateChannelRequest) (*CreateChannelResponse, error) {
	now := time.Now()

	channel := Channel{
		ChannelID: format.NewChannelID(),
		CreatedAt: now,
		UpdatedAt: now,
		MemberIDs: req.MemberIDs,
	}

	_, err := s.Collection(CHANNELS_COLLECTION).
		InsertOne(ctx, channel)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &channel, nil
}
