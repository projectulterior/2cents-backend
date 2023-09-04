package auth

import (
	"context"
	"fmt"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RefreshTokenRequest struct {
	Token string `json:"token"`
}
type RefreshTokenResponse = TokenResponse

func (s *Service) RefreshToken(ctx context.Context, req RefreshTokenRequest) (*RefreshTokenResponse, error) {
	e := func(err error) error { return status.Error(codes.PermissionDenied, err.Error()) }

	claims, err := verifyToken(s.Secret, req.Token)
	if err != nil {
		return nil, e(err)
	}

	tokenType, ok := claims["token_type"].(string)
	if !ok {
		return nil, e(fmt.Errorf("invalid token"))
	}

	if tokenType != REFRESH_TOKEN_TYPE {
		return nil, e(fmt.Errorf("invalid token_type"))
	}

	userID, err := format.ParseUserID(claims["user_id"].(string))
	if err != nil {
		return nil, e(err)
	}

	tokenID, err := format.ParseTokenID(claims["token_id"].(string))
	if err != nil {
		return nil, e(err)
	}

	// verify with database
	var token Token
	err = s.Collection(TOKENS_COLLECTION).
		FindOneAndUpdate(ctx,
			bson.M{
				"_id":     tokenID.String(),
				"user_id": userID.String(),
			},
			bson.M{
				"$inc": bson.M{
					"count": 1,
				},
				"$currentDate": bson.M{
					"refreshed_at": true,
				},
			},
			options.FindOneAndUpdate().
				SetReturnDocument(options.After),
		).Decode(&token)
	if err != nil {
		return nil, e(err)
	}

	auth, refresh, err := generateTokens(s.Secret, s.AuthTokenTTL, s.RefreshTokenTTL, tokenID, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &RefreshTokenResponse{
		Auth:    auth,
		Refresh: refresh,
	}, nil
}
