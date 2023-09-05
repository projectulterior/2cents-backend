package posts

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
)

const (
	POSTS_COLLECTION = "posts"
)

type Post struct {
	PostID      format.PostID      `bson:"_id"`
	AuthorID    format.UserID      `bson:"author_id"`
	Visibility  format.Visibility  `bson:"visibility"`
	Content     string             `bson:"content"`
	ContentType format.ContentType `bson:"content_type"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func (s *Service) Setup(ctx context.Context) error {
	return nil
}
