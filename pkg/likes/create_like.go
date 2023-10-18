package likes

import (
	"context"
	"fmt"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/cents"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/posts"
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

	session, err := s.Database.Client().StartSession()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {

		post, err := s.Posts.GetPost(ctx, posts.GetPostRequest{
			PostID: req.PostID,
		})
		if err != nil {
			return nil, err
		}
		fmt.Print("here")

		_, err = s.Collection(LIKES_COLLECTION).
			InsertOne(ctx, like)
		if err != nil {
			if !mongo.IsDuplicateKeyError(err) {
				return nil, status.Error(codes.Internal, err.Error())
			}
			fmt.Print("hello")

			// duplicate like
			return s.getLike(ctx, like.LikeID)
		}

		_, err = s.Cents.TransferCents(ctx, cents.TransferCentsRequest{
			SenderID:   req.LikerID,
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

	return &like, nil
}
