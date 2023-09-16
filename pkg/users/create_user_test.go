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

	updated := broker.Listener(users.USER_UPDATED_EVENT)

	userID := format.NewUserID()

	reply, err := svc.CreateUser(context.Background(), users.CreateUserRequest{
		UserID: userID,
	})
	assert.NoError(t, err)
	assert.Equal(t, userID, reply.UserID)

	event, err := updated.Next(context.Background())
	assert.NoError(t, err)

	user, ok := event.(users.UserUpdatedEvent)
	assert.True(t, ok)
	assert.Equal(t, *reply, user.User)
	assert.NotEmpty(t, user.Timestamp)
}
