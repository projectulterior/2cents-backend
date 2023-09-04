package users

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetUserRequest struct {
	UserID format.UserID
}
type GetUserResponse = User

func (s *Service) GetUser(ctx context.Context, req GetUserRequest) (*GetUserResponse, error) {
	var user User
	err := s.Collection(USERS_COLLECTION).
		FindOne(ctx, bson.M{"_id": req.UserID.String()}).
		Decode(&user)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &user, nil
}
