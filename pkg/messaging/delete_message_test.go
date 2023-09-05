package messaging_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/messaging"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeleteMessage(t *testing.T) {
	svc := setup(t)

	senderID := format.NewUserID()

	channel, err := svc.CreateChannel(context.Background(), messaging.CreateChannelRequest{
		MemberIDs: []format.UserID{
			senderID,
			format.NewUserID(),
		},
	})

	assert.NoError(t, err)

	content := "message"

	message, err := svc.CreateMessage(context.Background(), messaging.CreateMessageRequest{
		ChannelID:   channel.ChannelID,
		SenderID:    senderID,
		Content:     content,
		ContentType: format.TEXT,
	})
	assert.NoError(t, err)

	_, err = svc.GetMessage(context.Background(), messaging.GetMessageRequest{
		MessageID: message.MessageID,
	})
	assert.NoError(t, err)

	deleted, err := svc.DeleteMessage(context.Background(), messaging.DeleteMessageRequest{
		MessageID: message.MessageID,
		SenderID:  message.SenderID,
	})
	assert.NoError(t, err)
	assert.Equal(t, message.MessageID, deleted.MessageID)

	_, err = svc.GetMessage(context.Background(), messaging.GetMessageRequest{
		MessageID: message.MessageID,
	})
	assert.Equal(t, codes.NotFound, status.Code(err))
}
