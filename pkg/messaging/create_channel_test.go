package messaging_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/messaging"
	"github.com/stretchr/testify/assert"
)

func TestCreateChannel(t *testing.T) {
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
}
