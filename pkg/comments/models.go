package comments

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
)

const (
	COMMENT_COLLECTION = "comments"
)

type Comment struct {
	CommentID format.CommentID `bson:"_id"`
	PostID    format.PostID    `bson:"post_id"`
	Content   string           `bson:"content"`
	AuthorID  format.UserID    `bson:"author_id"`
	CreatedAt time.Time        `bson:"created_at"`
	UpdatedAt time.Time        `bson:"updated_at"`
}

func (s *Service) Setup(ctx context.Context) error {
	return nil
}
