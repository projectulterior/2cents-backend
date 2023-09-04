package likes

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateLikeRequest struct {
	PostID  format.PostID
	LikerID format.UserID
}

type CreateLikeResponse = Like

func (s *Service) CreateLike(ctx context.Context, req CreateLikeRequest) (*CreateLikeResponse, error) {
	like := Like{
		LikeID:    format.NewLikeID(req.PostID, req.LikerID),
		PostID:    req.PostID,
		LikerID:   req.LikerID,
		CreatedAt: time.Now(),
	}

	_, err := s.Collection(LIKES_COLLECTION).
		InsertOne(ctx, like)
	if err != nil {
		if !mongo.IsDuplicateKeyError(err) {
			return nil, status.Error(codes.Internal, err.Error())
		}

		// duplicate like
		return s.getLike(ctx, like.LikeID)
	}

	return &like, nil
}
