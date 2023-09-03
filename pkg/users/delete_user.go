package users

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteUserRequest struct {
	UserID format.UserID
}
type DeleteUserResponse struct {
	UserID format.UserID
}

func (s *Service) DeleteUser(ctx context.Context, req DeleteUserRequest) (*DeleteUserResponse, error) {
	_, err := s.Collection(USERS_COLLECTION).
		DeleteOne(ctx, bson.M{"_id": req.UserID.String()})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &DeleteUserResponse{
		UserID: req.UserID,
	}, nil
}
