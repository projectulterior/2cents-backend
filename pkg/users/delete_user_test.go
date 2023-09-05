package users_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/users"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeleteUser(t *testing.T) {
	svc := setup(t)

	userID := format.NewUserID()

	reply, err := svc.CreateUser(context.Background(), users.CreateUserRequest{
		UserID: userID,
	})
	assert.NoError(t, err)
	assert.Equal(t, userID, reply.UserID)

	delete, err := svc.DeleteUser(context.Background(), users.DeleteUserRequest{
		UserID: userID,
	})
	assert.NoError(t, err)
	assert.Equal(t, userID, delete.UserID)

	_, err = svc.GetUser(context.Background(), users.GetUserRequest{
		UserID: userID,
	})
	assert.Equal(t, codes.NotFound, status.Code(err))
}
