package likes_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/likes"
	"github.com/stretchr/testify/assert"
)

func TestGetLike(t *testing.T) {
	svc := setup(t)

	postID := format.NewPostID()
	likerID := format.NewUserID()

	reply, err := svc.CreateLike(context.Background(), likes.CreateLikeRequest{
		PostID:  postID,
		LikerID: likerID,
	})
	assert.NoError(t, err)
	assert.Equal(t, likerID, reply.LikerID)
	assert.NotEmpty(t, reply.LikeID)
	assert.Equal(t, postID, reply.PostID)

	_, err = svc.GetLike(context.Background(), likes.GetLikeRequest{
		LikeID: reply.LikeID,
	})
	assert.NoError(t, err)
	// TODO: finish test
	// assert.Equal(t, reply, user)
}
