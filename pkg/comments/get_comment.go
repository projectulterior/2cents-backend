package comments

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetCommentRequest struct {
	CommentID format.CommentID
}

type GetCommentResponse = Comment

func (s *Service) GetComment(ctx context.Context, req GetCommentRequest) (*GetCommentResponse, error) {
	var comment Comment

	err := s.Collection(COMMENT_COLLECTION).
		FindOne(ctx, bson.M{"_id": req.CommentID.String()}).
		Decode(&comment)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &comment, nil
}
