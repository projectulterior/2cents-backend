package auth

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdatePasswordRequest struct {
	UserID      format.UserID
	OldPassword string
	NewPassword string
}

func (s *Service) UpdatePassword(ctx context.Context, req UpdatePasswordRequest) error {
	oldPassword, err := salt(req.OldPassword)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	newPassword, err := salt(req.NewPassword)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	var user User
	err = s.Collection(USERS_COLLECTION).
		FindOneAndUpdate(ctx,
			bson.M{
				"_id":      req.UserID.String(),
				"password": oldPassword,
			},
			bson.M{
				"$set": bson.M{
					"password": newPassword,
				},
			},
			options.FindOneAndUpdate().
				SetReturnDocument(options.After),
		).Decode(&user)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return status.Error(codes.Internal, err.Error())
		}

		return status.Error(codes.NotFound, err.Error())
	}

	return nil
}
