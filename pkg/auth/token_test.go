package auth

import (
	"fmt"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	secret := "test"
	claims := jwt.MapClaims{
		"user_id":    format.NewUserID().String(),
		"token_type": AUTH_TOKEN_TYPE,
	}
	token, err := generateToken(secret, claims)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)

	assert.NotEmpty(t, token)

	verifiedClaims, err := verifyToken(secret, token)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, claims, verifiedClaims)
}
