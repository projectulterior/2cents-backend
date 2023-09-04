package users

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateUserRequest struct {
	UserID format.UserID
	Name   *string
	Email  *string
	Bio    *string
}
type UpdateUserResponse = User

func (s *Service) UpdateUser(ctx context.Context, req UpdateUserRequest) (*UpdateUserResponse, error) {
	set := bson.M{}

	if req.Name != nil {
		set["name"] = *req.Name
	}

	if req.Email != nil {
		set["email"] = *req.Email
	}

	if req.Bio != nil {
		set["bio"] = *req.Bio
	}

	var user User
	err := s.Collection(USERS_COLLECTION).
		FindOneAndUpdate(ctx,
			bson.M{"_id": req.UserID.String()},
			bson.M{
				"$set": set,
				"$currentDate": bson.M{
					"updated_at": true,
				},
			},
			options.FindOneAndUpdate().
				SetReturnDocument(options.After),
		).Decode(&user)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &user, nil
}
