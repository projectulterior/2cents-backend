package auth

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	USERS_COLLECTION  = "users"
	TOKENS_COLLECTION = "tokens"

	MIN_USERNAME_LENGTH = 5
)

type User struct {
	Username  string        `bson:"_id"`
	Password  string        `bson:"password"`
	UserID    format.UserID `bson:"user_id"`
	CreatedAt time.Time     `bson:"created_at"`
}

type Token struct {
	TokenID     format.TokenID `bson:"_id"`
	UserID      format.UserID  `bson:"user_id"`
	Count       int            `bson:"count"`
	CreatedAt   time.Time      `bson:"created_at"`
	RefreshedAt time.Time      `bson:"refreshed_at"`
	ExpiredAt   time.Time      `bson:"expired_at"`
}

func (s *Service) Setup(ctx context.Context) error {
	_, err := s.Collection(USERS_COLLECTION).
		Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "user_id", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		return err
	}

	return nil
}
