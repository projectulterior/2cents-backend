package follow_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/cents"
	"github.com/projectulterior/2cents-backend/pkg/follow"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/stretchr/testify/assert"
)

func TestCreateFollow(t *testing.T) {
	svc := setup(t)

	followerID := format.NewUserID()
	followeeID := format.NewUserID()

	followID := format.NewFollowID(followerID, followeeID)

	_, err := svc.Cents.UpdateCents(context.Background(), cents.UpdateCentsRequest{
		UserID: followerID,
		Amount: 2,
	})
	assert.NoError(t, err)

	reply, err := svc.CreateFollow(context.Background(), follow.CreateFollowRequest{
		FollowerID: followerID,
		FolloweeID: followeeID,
	})

	assert.NoError(t, err)
	assert.Equal(t, followID, reply.FollowID)
	assert.Equal(t, followerID, reply.FollowerID)
	assert.Equal(t, followeeID, reply.FolloweeID)
	assert.NotEmpty(t, reply.CreatedAt)

}
