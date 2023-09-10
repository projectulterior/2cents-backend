package follow_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/follow"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/stretchr/testify/assert"
)

func TestGetFollows(t *testing.T) {
	svc := setup(t)

	const (
		NUM_OF_FOLLOWS = 10
		BATCH_SIZE     = NUM_OF_FOLLOWS / 3
	)

	followerID := format.NewUserID()

	for i := 0; i < NUM_OF_FOLLOWS; i++ {
		_, err := svc.CreateFollow(context.Background(), follow.CreateFollowRequest{
			FollowerID: followerID,
			FolloweeID: format.NewUserIDFromIdentifier(fmt.Sprintf("%d", i)),
		})
		assert.NoError(t, err)

		time.Sleep(time.Millisecond)
	}

	i := NUM_OF_FOLLOWS - 1

	var cursor string
	for i >= 0 {
		follows, err := svc.GetFollows(context.Background(), &follow.GetFollowsRequest{
			Cursor: cursor,
			Limit:  BATCH_SIZE,
		})
		assert.NoError(t, err)

		for _, follow := range follows.Follows {
			expectedID := format.NewUserIDFromIdentifier(fmt.Sprintf("%d", i))
			assert.Equal(t, expectedID, follow.FolloweeID)
			i -= 1
		}

		cursor = follows.Next
	}
}

func TestGetUserFollows(t *testing.T) {
	svc := setup(t)

	const (
		NUM_OF_FOLLOWS = 10
		BATCH_SIZE     = NUM_OF_FOLLOWS / 3
	)

	followerID := format.NewUserID()

	for i := 0; i < NUM_OF_FOLLOWS; i++ {
		_, err := svc.CreateFollow(context.Background(), follow.CreateFollowRequest{
			FollowerID: followerID,
			FolloweeID: format.NewUserIDFromIdentifier(fmt.Sprintf("%d", i)),
		})
		assert.NoError(t, err)

		time.Sleep(time.Millisecond)

	}

	for i := 0; i < NUM_OF_FOLLOWS; i++ {
		_, err := svc.CreateFollow(context.Background(), follow.CreateFollowRequest{
			FollowerID: format.NewUserID(),
			FolloweeID: format.NewUserID(),
		})
		assert.NoError(t, err)

		time.Sleep(time.Millisecond)

	}

	i := NUM_OF_FOLLOWS - 1

	var cursor string
	for i >= 0 {
		follows, err := svc.GetFollows(context.Background(), &follow.GetFollowsRequest{
			Cursor:     cursor,
			Limit:      BATCH_SIZE,
			FollowerID: &followerID,
		})
		assert.NoError(t, err)

		for _, follow := range follows.Follows {
			expectedID := format.NewUserIDFromIdentifier(fmt.Sprintf("%d", i))
			assert.Equal(t, expectedID, follow.FolloweeID)
			i -= 1
		}

		cursor = follows.Next
	}
}
