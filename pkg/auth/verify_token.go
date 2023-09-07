package auth

import (
	"context"
	"fmt"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VerifyTokenRequest struct {
	Token string `json:"token"`
}
type VerifyTokenResponse struct {
	UserID format.UserID `json:"user_id"`
}

func (s *Service) VerifyToken(ctx context.Context, req VerifyTokenRequest) (*VerifyTokenResponse, error) {
	e := func(err error) error { return status.Error(codes.PermissionDenied, err.Error()) }

	claims, err := verifyToken(s.Secret, req.Token)
	if err != nil {
		return nil, e(err)
	}

	tokenType, ok := claims["token_type"].(string)
	if !ok {
		return nil, e(fmt.Errorf("invalid token"))
	}

	if tokenType != AUTH_TOKEN_TYPE {
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
		FindOne(ctx, bson.M{
			"_id":     tokenID.String(),
			"user_id": userID.String(),
		}).Decode(&token)
	if err != nil {
		return nil, e(err)
	}

	return &VerifyTokenResponse{
		UserID: userID,
	}, nil
}
