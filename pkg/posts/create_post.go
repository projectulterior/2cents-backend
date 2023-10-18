package posts

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/cents"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreatePostRequest struct {
	AuthorID    format.UserID
	Visibility  format.Visibility
	Content     string
	ContentType format.ContentType
}

type CreatePostResponse = Post

func (s *Service) CreatePost(ctx context.Context, req CreatePostRequest) (*CreatePostResponse, error) {
	now := time.Now()

	post := Post{
		PostID:      format.NewPostID(),
		AuthorID:    req.AuthorID,
		Visibility:  req.Visibility,
		Content:     req.Content,
		ContentType: req.ContentType,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	session, err := s.Database.Client().StartSession()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {

		_, err = s.Cents.TransferCents(ctx, cents.TransferCentsRequest{
			SenderID:   req.AuthorID,
			ReceiverID: format.DEFAULT_ADMIN_ID,
			Amount:     2,
		})
		if err != nil {
			return nil, err
		}

		_, err = s.Collection(POSTS_COLLECTION).
			InsertOne(ctx, post)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	return &post, nil
}
