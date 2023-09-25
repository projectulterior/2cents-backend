package auth

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ExportUsers(ctx context.Context) error {
	now := time.Now()

	cursor, err := s.Collection(USERS_COLLECTION).
		Find(ctx, bson.M{})
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	var users []User
	err = cursor.All(ctx, &users)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	for _, user := range users {
		s.UserUpdated.Publish(ctx, UserUpdatedEvent{
			User:      user,
			Timestamp: now,
		})
	}

	return nil
}
