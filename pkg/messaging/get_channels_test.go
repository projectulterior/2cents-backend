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

func TestGetChannels(t *testing.T) {
	svc := setup(t)

	const (
		NUM_OF_CHANNELS = 10
		BATCH_SIZE      = NUM_OF_CHANNELS / 3
	)

	userID := format.NewUserID()

	for i := 0; i < NUM_OF_CHANNELS; i++ {
		channel, err := svc.CreateChannel(context.Background(), messaging.CreateChannelRequest{
			MemberIDs: []format.UserID{userID, format.NewUserIDFromIdentifier(fmt.Sprintf("%d", i))},
		})
		assert.NoError(t, err)

		_, err = svc.CreateMessage(context.Background(), messaging.CreateMessageRequest{
			ChannelID:   channel.ChannelID,
			SenderID:    userID,
			Content:     "hello",
			ContentType: format.TEXT,
		})
		assert.NoError(t, err)

		time.Sleep(time.Millisecond)
	}

	i := NUM_OF_CHANNELS - 1

	var cursor string
	for i >= 0 {
		channels, err := svc.GetChannels(context.Background(), messaging.GetChannelsRequest{
			MemberID: userID,
			Cursor:   cursor,
			Limit:    BATCH_SIZE,
		})
		assert.NoError(t, err)
		assert.True(t, len(channels.Channels) > 0)

		for _, channel := range channels.Channels {
			expected := format.NewUserIDFromIdentifier(fmt.Sprintf("%d", i))
			assert.Equal(t, expected, channel.MemberIDs[1])
			i -= 1
		}

		cursor = channels.Next
	}
}
