package posts

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetPostRequest struct {
	PostID format.PostID
}

type GetPostResponse struct {
	Post Post
}

func (s *Service) GetPost(ctx context.Context, req GetPostRequest) (*GetPostResponse, error) {
	var post Post

	err := s.Collection(POSTS_COLLECTION).
		FindOne(ctx, bson.M{"_id": req.PostID.String()}).
		Decode(&post)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &GetPostResponse{
		Post: post,
	}, nil
}
