package users_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/users"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUser(t *testing.T) {
	svc := setup(t)

	userID := format.NewUserID()

	_, err := svc.CreateUser(context.Background(), users.CreateUserRequest{
		UserID: userID,
	})
	assert.NoError(t, err)

	name := "name"
	bio := "bio"
	email := "me@email.com"

	reply, err := svc.UpdateUser(context.Background(), users.UpdateUserRequest{
		UserID: userID,
		Name:   &name,
		Email:  &email,
		Bio:    &bio,
	})
	assert.NoError(t, err)
	assert.Equal(t, userID, reply.UserID)
	assert.Equal(t, name, reply.Name)
	assert.Equal(t, email, reply.Email)
	assert.Equal(t, bio, reply.Bio)

	same, err := svc.CreateUser(context.Background(), users.CreateUserRequest{
		UserID: userID,
	})
	assert.NoError(t, err)
	assert.Equal(t, reply, same)
}
