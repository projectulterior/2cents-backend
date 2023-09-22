package auth

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	var user User
	err := s.Collection(USERS_COLLECTION).
		FindOneAndUpdate(ctx,
			bson.M{"username": req.Username},
			bson.M{
				"$setOnInsert": bson.M{
					"_id":        format.NewUserID(),
					"password":   req.Password,
					"created_at": time.Now(),
				},
			},
			options.FindOneAndUpdate().
				SetReturnDocument(options.After).
				SetUpsert(true),
		).Decode(&user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if req.Password != user.Password {
		return nil, status.Error(codes.PermissionDenied, "wrong password")
	}

	auth, refresh, err := s.createToken(ctx, user.UserID)
	if err != nil {
		return nil, err
	}

	return &CreateTokenResponse{
		Auth:    auth,
		Refresh: refresh,
	}, nil
}

func verifyUsername(username string) bool {
	return len(username) > MIN_USERNAME_LENGTH
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
