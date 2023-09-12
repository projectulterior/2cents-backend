package comment_likes

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
)

const (
	COMMENT_LIKES_COLLECTION = "comment_likes"
)

type CommentLike struct {
	CommentLikeID format.CommentLikeID `bson:"_id"`
	CommentID     format.CommentID     `bson:"comment_id"`
	LikerID       format.UserID        `bson:"liker_id"`
	CreatedAt     time.Time            `bson:"created_at"`
}

func (s *Service) Setup(ctx context.Context) error {
	return nil
}
