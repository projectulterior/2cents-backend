package likes

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
)

const (
	LIKES_COLLECTION = "likes"
)

type Like struct {
	LikeID    format.LikeID `bson:"_id"`
	PostID    format.PostID `bson:"post_id"`
	LikerID   format.UserID `bson:"liker_id"`
	CreatedAt time.Time     `bson:"created_at"`
}

func (s *Service) Setup(ctx context.Context) error {
	return nil
}
