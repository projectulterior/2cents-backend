package likes

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteLikeRequest struct {
	LikeID format.LikeID
}

type DeleteLikeResponse struct {
	LikeID format.LikeID
}

func (s *Service) DeleteLike(ctx context.Context, req DeleteLikeRequest) (*DeleteLikeResponse, error) {
	_, err := s.Collection(LIKES_COLLECTION).
		DeleteOne(ctx, bson.M{"_id": req.LikeID.String()})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &DeleteLikeResponse{
		LikeID: req.LikeID,
	}, nil
}
