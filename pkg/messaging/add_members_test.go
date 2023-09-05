package messaging_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/messaging"
	"github.com/stretchr/testify/assert"
)

func TestAddMembers(t *testing.T) {
	svc := setup(t)

	senderID := format.NewUserID()
	receiverID := format.NewUserID()

	channel, err := svc.CreateChannel(context.Background(), messaging.CreateChannelRequest{
		MemberIDs: []format.UserID{
			senderID,
			receiverID,
		},
	})
	assert.Nil(t, err)

	memberID := format.NewUserID()

	reply, err := svc.AddMembers(context.Background(), messaging.AddMembersRequest{
		ChannelID: channel.ChannelID,
		MemberID:  senderID,
		MemberIDs: []format.UserID{memberID},
	})
	assert.NoError(t, err)
	assert.Equal(t, channel.ChannelID, reply.ChannelID)
	assert.Contains(t, reply.MemberIDs, senderID)
	assert.Contains(t, reply.MemberIDs, receiverID)
	assert.Contains(t, reply.MemberIDs, memberID)
	assert.NotEmpty(t, reply.UpdatedAt)
}
