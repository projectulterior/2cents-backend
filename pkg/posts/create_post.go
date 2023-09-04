package posts

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreatePostRequest struct {
	AuthorID    format.UserID
	Visibility  format.Visibility
	Content     string
	ContentType format.ContentType
}

type CreatePostResponse = Post

func (s *Service) CreatePost(ctx context.Context, req CreatePostRequest) (*CreatePostResponse, error) {
	now := time.Now()

	post := Post{
		PostID:      format.NewPostID(),
		AuthorID:    req.AuthorID,
		Visibility:  req.Visibility,
		Content:     req.Content,
		ContentType: req.ContentType,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	_, err := s.Collection(POSTS_COLLECTION).
		InsertOne(ctx, post)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &post, nil
}
