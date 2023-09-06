package posts

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeletePostRequest struct {
	PostID   format.PostID
	AuthorID format.UserID
}

type DeletePostResponse struct {
	PostID format.PostID
}

func (s *Service) DeletePost(ctx context.Context, req DeletePostRequest) (*DeletePostResponse, error) {
	_, err := s.Collection(POSTS_COLLECTION).
		DeleteOne(ctx, bson.M{
			"_id":       req.PostID.String(),
			"author_id": req.AuthorID.String(),
		})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &DeletePostResponse{
		PostID: req.PostID,
	}, nil
}
