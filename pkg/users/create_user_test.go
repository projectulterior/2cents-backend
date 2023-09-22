package users_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/users"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	svc := setup(t)

	updated := userUpdated.Listener()

	userID := format.NewUserID()

	reply, err := svc.CreateUser(context.Background(), users.CreateUserRequest{
		UserID: userID,
	})
	assert.NoError(t, err)
	assert.Equal(t, userID, reply.UserID)

	event, err := updated.Next(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, *reply, event.User)
	assert.NotEmpty(t, event.Timestamp)
}
