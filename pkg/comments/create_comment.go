package comments

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateCommentRequest struct {
	PostID   format.PostID
	Content  string
	AuthorID format.UserID
}

type CreateCommentResponse = Comment

func (s *Service) CreateComment(ctx context.Context, req CreateCommentRequest) (*CreateCommentResponse, error) {
	now := time.Now()

	comment := Comment{
		CommentID: format.NewCommentID(),
		PostID:    req.PostID,
		Content:   req.Content,
		AuthorID:  req.AuthorID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err := s.Collection(COMMENT_COLLECTION).
		InsertOne(ctx, comment)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &comment, nil
}
