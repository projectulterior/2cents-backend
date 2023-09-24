package auth_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/auth"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateToken(t *testing.T) {
	svc := setup(t)

	username := "username"
	password := "password"

	resp, err := svc.CreateToken(context.Background(), auth.CreateTokenRequest{
		Username: username,
		Password: password,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Auth)
	assert.NotEmpty(t, resp.Refresh)

	t.Run("second time", func(t *testing.T) {
		resp, err := svc.CreateToken(context.Background(), auth.CreateTokenRequest{
			Username: username,
			Password: password,
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.Auth)
		assert.NotEmpty(t, resp.Refresh)
	})

	t.Run("wrong password", func(t *testing.T) {
		_, err = svc.CreateToken(context.Background(), auth.CreateTokenRequest{
			Username: username,
			Password: password + "1",
		})
		assert.Equal(t, codes.PermissionDenied, status.Code(err))
	})
}
