package auth

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateTokenResponse = TokenResponse

func (s *Service) CreateToken(ctx context.Context, req CreateTokenRequest) (*CreateTokenResponse, error) {
	if !verifyUsername(req.Username) {
		return nil, status.Error(codes.InvalidArgument, "invalid username")
	}

	password, err := salt(req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	userID := format.NewUserID()

	var user User
	err = s.Collection(USERS_COLLECTION).
		FindOneAndUpdate(ctx,
			bson.M{"username": req.Username, "password": password},
			bson.M{
				"$setOnInsert": bson.M{
					"_id":        userID,
					"created_at": time.Now(),
				},
			},
			options.FindOneAndUpdate().
				SetReturnDocument(options.After).
				SetUpsert(true),
		).Decode(&user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, status.Error(codes.PermissionDenied, "wrong password")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	auth, refresh, err := s.createToken(ctx, user.UserID)
	if err != nil {
		return nil, err
	}

	if user.UserID == userID {
		// new user
		s.UserUpdated.Publish(ctx, UserUpdatedEvent{
			User:      user,
			Timestamp: user.CreatedAt,
		})
	}

	return &CreateTokenResponse{
		Auth:    auth,
		Refresh: refresh,
	}, nil
}

func verifyUsername(username string) bool {
	return len(username) > MIN_USERNAME_LENGTH
}

func salt(str string) (string, error) {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}

func (s *Service) createToken(ctx context.Context, userID format.UserID) (string, string, error) {
	now := time.Now()
	tokenID := format.NewTokenID()

	_, err := s.Collection(TOKENS_COLLECTION).
		InsertOne(ctx, Token{
			TokenID:     tokenID,
			UserID:      userID,
			CreatedAt:   now,
			RefreshedAt: now,
			ExpiredAt:   now.Add(s.AuthTokenTTL),
		})
	if err != nil {
		return "", "", status.Error(codes.Internal, err.Error())
	}

	return generateTokens(s.Secret, s.AuthTokenTTL, s.RefreshTokenTTL, tokenID, userID)
}
