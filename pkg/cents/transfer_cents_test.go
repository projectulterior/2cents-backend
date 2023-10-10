package cents_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/cents"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/stretchr/testify/assert"
)

func TestTransferCents(t *testing.T) {
	svc := setup(t)

	senderID := format.NewUserID()
	receiverID := format.NewUserID()
	amount := 10
	toSend := 1

	reply, err := svc.UpdateCents(context.Background(), cents.UpdateCentsRequest{
		UserID: senderID,
		Amount: amount,
	})
	assert.NoError(t, err)
	assert.Equal(t, senderID, reply.UserID)
	assert.Equal(t, amount, reply.Total)
	assert.Equal(t, amount, reply.Deposited)
	assert.Zero(t, reply.Received)
	assert.Zero(t, reply.Sent)

	reply, err = svc.TransferCents(context.Background(), cents.TransferCentsRequest{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Amount:     toSend,
	})
	assert.NoError(t, err)
	assert.Equal(t, senderID, reply.UserID)
	assert.Equal(t, amount-toSend, reply.Total)
	assert.Equal(t, amount, reply.Deposited)
	assert.Zero(t, reply.Received)
	assert.Equal(t, toSend, reply.Sent)

	receiver, err := svc.GetCents(context.Background(), cents.GetCentsRequest{
		UserID: receiverID,
	})
	assert.NoError(t, err)
	assert.Equal(t, receiverID, receiver.UserID)
	assert.Equal(t, toSend, receiver.Total)
	assert.Equal(t, toSend, receiver.Received)
	assert.Zero(t, receiver.Deposited)
	assert.Zero(t, receiver.Sent)
}
