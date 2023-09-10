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

	for i := 0; i < NUM_OF_CHANNELS; i++ {
		_, err := svc.CreateChannel(context.Background(), messaging.CreateChannelRequest{
			MemberIDs: []format.UserID{format.NewUserIDFromIdentifier(fmt.Sprintf("%d", i))},
		})
		assert.NoError(t, err)

		time.Sleep(time.Millisecond)
	}

	i := NUM_OF_CHANNELS - 1

	var cursor string
	for i >= 0 {
		channels, err := svc.GetChannels(context.Background(), &messaging.GetChannelsRequest{
			Cursor: cursor,
			Limit:  BATCH_SIZE,
		})
		assert.NoError(t, err)

		for _, channel := range channels.Channels {
			expected := format.NewUserIDFromIdentifier(fmt.Sprintf("%d", i))
			assert.Equal(t, expected, channel.MemberIDs[0])
			i -= 1
		}

		cursor = channels.Next
	}
}
