package comments

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteCommentRequest struct {
	CommentID format.CommentID
}

type DeleteCommentResponse struct {
	CommentID format.CommentID
}

func (s *Service) DeleteComment(ctx context.Context, req DeleteCommentRequest) (*DeleteCommentResponse, error) {
	_, err := s.Collection(COMMENT_COLLECTION).
		DeleteOne(ctx, bson.M{"_id": req.CommentID.String()})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &DeleteCommentResponse{
		CommentID: req.CommentID,
	}, nil
}
