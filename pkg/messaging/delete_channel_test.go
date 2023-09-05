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

func TestDeleteChannel(t *testing.T) {
	svc := setup(t)

	reply, err := svc.CreateChannel(context.Background(), messaging.CreateChannelRequest{
		MemberIDs: []format.UserID{
			format.NewUserID(),
			format.NewUserID(),
		},
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, reply.ChannelID)
	assert.NotEmpty(t, reply.MemberIDs)
	assert.False(t, reply.CreatedAt.IsZero())

	deleted, err := svc.DeleteChannel(context.Background(), messaging.DeleteChannelRequest{
		ChannelID: reply.ChannelID,
	})

	assert.Nil(t, err)
	assert.Equal(t, reply.ChannelID, deleted.ChannelID)

	_, err = svc.GetChannel(context.Background(), messaging.GetChannelRequest{
		ChannelID: reply.ChannelID,
	})
	assert.Equal(t, codes.NotFound, status.Code(err))
}
