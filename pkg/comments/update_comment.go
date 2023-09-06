package comments

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateCommentRequest struct {
	CommentID format.CommentID
	AuthorID  format.UserID
	Content   string
}

type UpdateCommentResponse = Comment

func (s *Service) UpdateComment(ctx context.Context, req UpdateCommentRequest) (*UpdateCommentResponse, error) {
	var comment Comment

	err := s.Collection(COMMENTS_COLLECTION).
		FindOneAndUpdate(ctx,
			bson.M{
				"_id":       req.CommentID.String(),
				"author_id": req.AuthorID.String(),
			},
			bson.M{
				"$set": bson.M{"content": req.Content},
				"$currentDate": bson.M{
					"updated_at": true,
				},
			},
			options.FindOneAndUpdate().
				SetReturnDocument(options.After),
		).Decode(&comment)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &comment, nil
}
