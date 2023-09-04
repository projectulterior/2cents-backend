package messaging

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
)

type AddMembersRequest struct {
	ChannelID format.ChannelID
	MemberIDs []format.UserID
}

type AddMembersResponse = Channel

func (s *Service) AddMembers(ctx context.Context, req AddMembersRequest) (*AddMembersResponse, error) {
	return nil, nil
}
