package comment_likes

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateCommentLikeRequest struct {
	CommentID format.CommentID
	LikerID   format.UserID
}

type CreateCommentLikeResponse = CommentLike

func (s *Service) CreateCommentLike(ctx context.Context, req CreateCommentLikeRequest) (*CreateCommentLikeResponse, error) {
	commentLike := CommentLike{
		CommentLikeID: format.NewCommentLikeID(req.CommentID, req.LikerID),
		CommentID:     req.CommentID,
		LikerID:       req.LikerID,
		CreatedAt:     time.Now(),
	}

	_, err := s.Collection(COMMENT_LIKES_COLLECTION).
		InsertOne(ctx, commentLike)
	if err != nil {
		if !mongo.IsDuplicateKeyError(err) {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return s.getCommentLike(ctx, commentLike.CommentLikeID)
	}

	return &commentLike, nil
}
