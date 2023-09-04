package users

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateUserRequest struct {
	UserID format.UserID
}
type CreateUserResponse = User

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error) {
	now := time.Now()

	var user User
	err := s.Collection(USERS_COLLECTION).
		FindOneAndUpdate(ctx,
			bson.M{"_id": req.UserID.String()},
			bson.M{
				"$set": bson.M{},
				"$setOnInsert": bson.M{
					"created_at": now,
					"updated_at": now,
				},
			},
			options.FindOneAndUpdate().
				SetReturnDocument(options.After).
				SetUpsert(true),
		).Decode(&user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &user, nil
}
