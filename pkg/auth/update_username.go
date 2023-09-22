package auth

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
)

type UpdateUsernameRequest struct {
	UserID   format.UserID
	Username string
}

type UpserUsernameResponse = User

func (s *Service) UpdateUsername(ctx context.Context, req UpdateUsernameRequest) (*UpserUsernameResponse, error) {
	return nil, nil
}
