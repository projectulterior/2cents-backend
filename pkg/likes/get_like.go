package likes

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetLikeRequest struct {
	LikeID format.LikeID
}

type GetLikeResponse = Like

func (s *Service) GetLike(ctx context.Context, req GetLikeRequest) (*GetLikeResponse, error) {
	return s.getLike(ctx, req.LikeID)
}

func (s *Service) getLike(ctx context.Context, likeID format.LikeID) (*Like, error) {
	var like Like
	err := s.Collection(LIKES_COLLECTION).
		FindOne(ctx, bson.M{"_id": likeID.String()}).
		Decode(&like)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &like, nil
}
