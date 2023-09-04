package posts

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdatePostRequest struct {
	PostID      format.PostID
	AuthorID    format.UserID
	Visibility  *format.Visibility
	Content     *string
	ContentType *format.ContentType
}

type UpdatePostResponse struct {
	Post Post
}

func (s *Service) UpdatePost(ctx context.Context, req UpdatePostRequest) (*UpdatePostResponse, error) {
	set := bson.M{}

	if req.Visibility != nil {
		set["visibility"] = *req.Visibility
	}

	if req.Content != nil {
		set["content"] = *req.Content
	}

	if req.ContentType != nil {
		set["content_type"] = *req.ContentType
	}

	var post Post
	err := s.Collection(POSTS_COLLECTION).
		FindOneAndUpdate(ctx,
			bson.M{
				"_id":       req.PostID.String(),
				"author_id": req.AuthorID.String(),
			},
			bson.M{
				"$set": set,
				"$currentDate": bson.M{
					"updated_at": true,
				},
			},
			options.FindOneAndUpdate().
				SetReturnDocument(options.After),
		).Decode(&post)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &UpdatePostResponse{
		Post: post,
	}, nil
}
