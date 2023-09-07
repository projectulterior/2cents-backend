package comments

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteCommentRequest struct {
	CommentID format.CommentID
	DeleterID format.UserID
}

type DeleteCommentResponse struct {
	CommentID format.CommentID
}

func (s *Service) DeleteComment(ctx context.Context, req DeleteCommentRequest) (*DeleteCommentResponse, error) {
	// find commment & delete if deleteID == comment.author_id
	err := s.Collection(COMMENTS_COLLECTION).
		FindOneAndDelete(ctx, bson.M{
			"_id":       req.CommentID.String(),
			"author_id": req.DeleterID.String(),
		}).Err()
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		// here: comment with author_id not found

		// find the comment by id to get post id to
		// check if deleted_id == post.author_id
		var comment Comment
		err := s.Collection(COMMENTS_COLLECTION).
			FindOne(ctx, bson.M{"_id": req.CommentID.String()}).
			Decode(&comment)
		if err != nil {
			if err != mongo.ErrNoDocuments {
				return nil, status.Error(codes.Internal, err.Error())
			}

			// here: comment doesn't exist, can return as if success
			goto SUCCESS
		}

		// here: comment does exist, must check post for author_id
		post, err := s.GetPost(ctx, posts.GetPostRequest{
			PostID: comment.PostID,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		if req.DeleterID == post.AuthorID {
			err := s.Collection(COMMENTS_COLLECTION).
				FindOneAndDelete(ctx, bson.M{"_id": req.CommentID.String()}).
				Decode(&comment)
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
		}
	}

SUCCESS:
	return &DeleteCommentResponse{
		CommentID: req.CommentID,
	}, nil
}
