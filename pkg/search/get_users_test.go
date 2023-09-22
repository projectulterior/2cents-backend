package search_test

import (
	"context"
	"testing"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/auth"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/search"
	"github.com/projectulterior/2cents-backend/pkg/users"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	svc := setup(t)

	brianID := format.NewUserIDFromIdentifier("brian")
	calebID := format.NewUserIDFromIdentifier("caleb")
	davidID := format.NewUserIDFromIdentifier("david")

	userIDs := []format.UserID{brianID, calebID, davidID}

	for _, userID := range userIDs {
		err := svc.ProcessUserUpdated(context.Background(), users.UserUpdatedEvent{
			User: users.User{
				UserID: userID,
				Name:   userID.Identifier(),
			},
		})
		assert.NoError(t, err)

		err = svc.ProcessUsernameUpdated(context.Background(), auth.UserUpdatedEvent{
			User: auth.User{
				UserID:   userID,
				Username: userID.Identifier(),
			},
		})
		assert.NoError(t, err)
	}

	time.Sleep(time.Second)

	reply, err := svc.GetUsers(context.Background(), search.GetUsersRequest{
		Query: brianID.Identifier(),
		Limit: 1,
	})
	assert.NoError(t, err)
	assert.Len(t, reply.Users, 1)
	assert.Equal(t, brianID, reply.Users[0].UserID)
}
