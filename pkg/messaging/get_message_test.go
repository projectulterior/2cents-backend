package messaging_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/messaging"
	"github.com/stretchr/testify/assert"
)

func TestGetMessage(t *testing.T) {
	svc := setup(t)

	senderID := format.NewUserID()
	receiverID := format.NewUserID()

	channel, err := svc.CreateChannel(context.Background(), messaging.CreateChannelRequest{
		MemberIDs: []format.UserID{
			senderID,
			receiverID,
		},
	})

	assert.NoError(t, err)
	assert.Contains(t, channel.MemberIDs, senderID)
	assert.Contains(t, channel.MemberIDs, receiverID)
	assert.False(t, channel.CreatedAt.IsZero())

	content := "message"

	message, err := svc.CreateMessage(context.Background(), messaging.CreateMessageRequest{
		ChannelID:   channel.ChannelID,
		SenderID:    senderID,
		Content:     content,
		ContentType: format.TEXT,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, message.MessageID)
	assert.NotEmpty(t, message.SenderID)
	assert.False(t, message.CreatedAt.IsZero())

}
