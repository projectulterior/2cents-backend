package messaging_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/messaging"
	"github.com/stretchr/testify/assert"
)

func TestGetChannel(t *testing.T) {
	svc := setup(t)

	senderID := format.NewUserID()
	receiverID := format.NewUserID()

	reply, err := svc.CreateChannel(context.Background(), messaging.CreateChannelRequest{
		MemberIDs: []format.UserID{
			senderID,
			receiverID,
		},
	})

	assert.NoError(t, err)
	assert.Contains(t, reply.MemberIDs, senderID)
	assert.Contains(t, reply.MemberIDs, receiverID)
	assert.False(t, reply.CreatedAt.IsZero())

	get, err := svc.GetChannel(context.Background(), messaging.GetChannelRequest{
		ChannelID: reply.ChannelID,
	})
	assert.NoError(t, err)
	assert.Equal(t, reply.ChannelID, get.ChannelID)
	assert.Equal(t, reply.MemberIDs, get.MemberIDs)
	assert.NotEmpty(t, get.CreatedAt)
}
