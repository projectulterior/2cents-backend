package likes_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/likes"
	"github.com/stretchr/testify/assert"
)

func TestCreateLike(t *testing.T) {
	svc := setup(t)

	likerID := format.NewUserID()
	postID := format.NewPostID()

	likeID := format.NewLikeID(postID, likerID)

	reply, err := svc.CreateLike(context.Background(), likes.CreateLikeRequest{
		PostID:  postID,
		LikerID: likerID,
	})
	assert.Nil(t, err)
	assert.Equal(t, likeID, reply.LikeID)
	assert.Equal(t, postID, reply.PostID)
	assert.Equal(t, likerID, reply.LikerID)
	assert.False(t, reply.CreatedAt.IsZero())

	same, err := svc.CreateLike(context.Background(), likes.CreateLikeRequest{
		PostID:  postID,
		LikerID: likerID,
	})
	assert.Nil(t, err)
	assert.Equal(t, reply.LikeID, same.LikeID)
}
