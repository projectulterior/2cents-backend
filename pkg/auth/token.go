package auth

import (
	"fmt"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"

	"github.com/golang-jwt/jwt/v4"
)

type AuthType string

const (
	AUTH_TOKEN_TYPE    = "auth"
	REFRESH_TOKEN_TYPE = "refresh"

	AUTH_TOKEN_TTL    = 10 * time.Minute
	REFRESH_TOKEN_TTL = 24 * time.Hour
)

func generateTokens(secret string, tokenID format.TokenID, userID format.UserID) (string, string, error) {
	token, err := generateToken(secret, jwt.MapClaims{
		"token_id":   tokenID.String(),
		"user_id":    userID.String(),
		"token_type": AUTH_TOKEN_TYPE,
		"exp":        time.Now().Add(AUTH_TOKEN_TTL).Unix(),
	})
	if err != nil {
		return "", "", err
	}

	refresh, err := generateToken(secret, jwt.MapClaims{
		"token_id":   tokenID.String(),
		"user_id":    userID.String(),
		"token_type": REFRESH_TOKEN_TYPE,
		"exp":        time.Now().Add(REFRESH_TOKEN_TTL).Unix(),
	})
	if err != nil {
		return "", "", err
	}

	return token, refresh, nil
}

func generateToken(secret string, claims jwt.MapClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}

func verifyToken(secret, tokenString string) (jwt.MapClaims, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
