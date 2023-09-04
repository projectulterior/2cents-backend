package likes

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
)

type CreateLikeRequest struct {
	PostID  format.PostID
	LikerID format.UserID
}

func (s *Service) CreateLike(ctx context.Context)
