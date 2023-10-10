package cents_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/cents"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCents(t *testing.T) {
	svc := setup(t)

	userID := format.NewUserID()
	amount := 5

	reply, err := svc.UpdateCents(context.Background(), cents.UpdateCentsRequest{
		UserID: userID,
		Amount: amount,
	})
	assert.NoError(t, err)
	assert.Equal(t, userID, reply.UserID)
	assert.Equal(t, amount, reply.Total)
	assert.Equal(t, amount, reply.Deposited)
	assert.Zero(t, reply.Received)
	assert.Zero(t, reply.Sent)

}
