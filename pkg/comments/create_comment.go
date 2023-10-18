package comments

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/cents"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateCommentRequest struct {
	PostID   format.PostID
	Content  string
	AuthorID format.UserID
}

type CreateCommentResponse = Comment

func (s *Service) CreateComment(ctx context.Context, req CreateCommentRequest) (*CreateCommentResponse, error) {
	now := time.Now()

	comment := Comment{
		CommentID: format.NewCommentID(),
		PostID:    req.PostID,
		Content:   req.Content,
		AuthorID:  req.AuthorID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	session, err := s.Database.Client().StartSession()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	_, err = session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {

		post, err := s.Posts.GetPost(ctx, posts.GetPostRequest{
			PostID: req.PostID,
		})
		if err != nil {
			return nil, err
		}

		_, err = s.Collection(COMMENTS_COLLECTION).
			InsertOne(ctx, comment)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		_, err = s.Cents.TransferCents(ctx, cents.TransferCentsRequest{
			SenderID:   req.AuthorID,
			ReceiverID: post.AuthorID,
			Amount:     1,
		})
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
