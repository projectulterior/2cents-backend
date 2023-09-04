package likes_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/likes"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeleteLike(t *testing.T) {
	svc := setup(t)

	likerID := format.NewUserID()

	reply, err := svc.CreateLike(context.Background(), likes.CreateLikeRequest{
		LikerID: likerID,
	})
	assert.Nil(t, err)
	assert.Equal(t, likerID, reply.LikerID)

	delete, err := svc.DeleteLike(context.Background(), likes.DeleteLikeRequest{
		LikeID: reply.LikeID,
	})
	assert.Nil(t, err)
	assert.Equal(t, reply.LikeID, delete.LikeID)

	_, err = svc.GetLike(context.Background(), likes.GetLikeRequest{
		LikeID: reply.LikeID,
	})

	assert.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))

	delete, err = svc.DeleteLike(context.Background(), likes.DeleteLikeRequest{
		LikeID: reply.LikeID,
	})
	assert.Nil(t, err)
	assert.Equal(t, reply.LikeID, delete.LikeID)

	_, err = svc.GetLike(context.Background(), likes.GetLikeRequest{
		LikeID: reply.LikeID,
	})

	assert.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))
}
