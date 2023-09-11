package comment_likes

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteCommentLikeRequest struct {
	CommentLikeID format.CommentLikeID
}

type DeleteCommentLikeResponse struct {
	CommentLikeID format.CommentLikeID
}

func (s *Service) DeleteCommentLike(ctx context.Context, req DeleteCommentLikeRequest) (*DeleteCommentLikeResponse, error) {
	_, err := s.Collection(COMMENT_LIKES_COLLECTION).
		DeleteOne(ctx, bson.M{"_id": req.CommentLikeID.String()})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &DeleteCommentLikeResponse{
		CommentLikeID: req.CommentLikeID,
	}, nil
}
