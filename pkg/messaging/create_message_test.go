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

func TestCreateMessage(t *testing.T) {
	svc := setup(t)

	senderID := format.NewUserID()
	receiverID := format.NewUserID()

	channel, err := svc.CreateChannel(context.Background(), messaging.CreateChannelRequest{
		MemberIDs: []format.UserID{senderID, receiverID},
	})
	assert.Nil(t, err)

	content := "message"

	reply, err := svc.CreateMessage(context.Background(), messaging.CreateMessageRequest{
		ChannelID:   channel.ChannelID,
		SenderID:    senderID,
		Content:     content,
		ContentType: format.TEXT,
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, reply.MessageID)
	assert.NotEmpty(t, reply.SenderID)
	assert.False(t, reply.CreatedAt.IsZero())
}

func TestCreateMessage_NotFound(t *testing.T) {
	svc := setup(t)

	t.Run("channel not found", func(t *testing.T) {
		_, err := svc.CreateMessage(context.Background(), messaging.CreateMessageRequest{
			ChannelID:   format.NewChannelID(),
			SenderID:    format.NewUserID(),
			Content:     "content",
			ContentType: format.TEXT,
		})
		assert.Equal(t, codes.NotFound, status.Code(err))
	})

	t.Run("member not found", func(t *testing.T) {
		senderID := format.NewUserID()
		receiverID := format.NewUserID()

		channel, err := svc.CreateChannel(context.Background(), messaging.CreateChannelRequest{
			MemberIDs: []format.UserID{senderID, receiverID},
		})
		assert.Nil(t, err)

		_, err = svc.CreateMessage(context.Background(), messaging.CreateMessageRequest{
			ChannelID:   channel.ChannelID,
			SenderID:    format.NewUserID(),
			Content:     "message",
			ContentType: format.TEXT,
		})
		assert.Equal(t, codes.NotFound, status.Code(err))
	})
}
