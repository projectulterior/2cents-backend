package users_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/users"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	svc := setup(t)

	const (
		NUM_OF_USERS = 10
		BATCH_SIZE   = NUM_OF_USERS / 3
	)

	// iusr[i] Ex. iusr0, iusr1 ...
	for i := 0; i < NUM_OF_USERS; i++ {
		_, err := svc.CreateUser(context.Background(), users.CreateUserRequest{
			UserID: format.NewUserIDFromIdentifier(fmt.Sprintf("%d", i)),
		})
		assert.NoError(t, err)
	}

	i := NUM_OF_USERS - 1 // iusr9

	var cursor string
	for i >= 0 {
		users, err := svc.GetUsers(context.Background(), &users.GetUsersRequest{
			Cursor: cursor,
			Limit:  BATCH_SIZE,
		})
		assert.NoError(t, err)

		for _, user := range users.Users {
			expectedID := format.NewUserIDFromIdentifier(fmt.Sprintf("%d", i))
			assert.Equal(t, expectedID, user.UserID)
			i -= 1
		}

		cursor = users.Next
	}
}
