package messaging_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/messaging"
	"github.com/stretchr/testify/assert"
)

func TestGetMessages(t *testing.T) {
	svc := setup(t)

	const (
		NUM_OF_MESSAGES = 10
		BATCH_SIZE      = NUM_OF_MESSAGES / 3
	)

	authorID := format.NewUserID()

	channel, err := svc.CreateChannel(context.Background(), messaging.CreateChannelRequest{
		MemberIDs: []format.UserID{authorID},
	})
	assert.NoError(t, err)

	for i := 0; i < NUM_OF_MESSAGES; i++ {
		_, err := svc.CreateMessage(context.Background(), messaging.CreateMessageRequest{
			ChannelID:   channel.ChannelID,
			SenderID:    authorID,
			Content:     fmt.Sprintf("%d", i),
			ContentType: format.TEXT,
		})
		assert.NoError(t, err)

		time.Sleep(time.Millisecond)
	}

	i := NUM_OF_MESSAGES - 1

	var cursor string
	for i >= 0 {
		messages, err := svc.GetMessages(context.Background(), &messaging.GetMessagesRequest{
			ChannelID: channel.ChannelID,
			Cursor:    cursor,
			Limit:     BATCH_SIZE,
		})
		assert.NoError(t, err)

		for _, message := range messages.Messages {
			expected := fmt.Sprintf("%d", i)
			assert.Equal(t, expected, message.Content)
			i -= 1
		}

		cursor = messages.Next
	}
}
