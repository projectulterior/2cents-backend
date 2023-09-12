package comment_likes

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetCommentLikeRequest struct {
	CommentLikeID format.CommentLikeID
}

type GetCommentLikeResponse = CommentLike

func (s *Service) GetCommentLike(ctx context.Context, req GetCommentLikeRequest) (*GetCommentLikeResponse, error) {
	return s.getCommentLike(ctx, req.CommentLikeID)
}

func (s *Service) getCommentLike(ctx context.Context, commentLikeID format.CommentLikeID) (*CommentLike, error) {
	var commentLike CommentLike
	err := s.Collection(COMMENT_LIKES_COLLECTION).
		FindOne(ctx, bson.M{"_id": commentLikeID.String()}).
		Decode(&commentLike)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &commentLike, nil
}
