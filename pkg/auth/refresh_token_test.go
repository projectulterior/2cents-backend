package auth_test

import (
	"context"
	"testing"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/auth"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRefreshToken(t *testing.T) {
	svc := setup(t)

	duration := time.Millisecond

	svc.AuthTokenTTL = duration

	username := "username"
	password := "password"

	reply, err := svc.CreateToken(context.Background(), auth.CreateTokenRequest{
		Username: username,
		Password: password,
	})
	assert.Nil(t, err)

	time.Sleep(duration)

	_, err = svc.VerifyToken(context.Background(), auth.VerifyTokenRequest{
		Token: reply.Auth,
	})
	assert.Equal(t, codes.PermissionDenied, status.Code(err))

	svc.AuthTokenTTL = time.Minute

	reply, err = svc.RefreshToken(context.Background(), auth.RefreshTokenRequest{
		Token: reply.Refresh,
	})
	assert.Nil(t, err)

	_, err = svc.VerifyToken(context.Background(), auth.VerifyTokenRequest{
		Token: reply.Auth,
	})
	assert.Nil(t, err)
}
