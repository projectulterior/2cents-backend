package search_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/users"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	svc := setup(t)

	brianID := format.NewUserIDFromIdentifier("brian")
	// calebID := format.NewUserIDFromIdentifier("caleb")
	// davidID := format.NewUserIDFromIdentifier("david")

	err := svc.ProcessUserUpdated(context.Background(), users.UserUpdatedEvent{
		User: users.User{
			UserID: brianID,
			Name:   "brian",
		},
	})
	assert.NoError(t, err)
}
