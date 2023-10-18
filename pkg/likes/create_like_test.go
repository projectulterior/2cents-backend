package likes_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/cents"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/likes"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/stretchr/testify/assert"
)

func TestCreateLike(t *testing.T) {
	svc := setup(t)

	likerID := format.NewUserID()
	posterID := format.NewUserID()

	_, err := svc.Cents.UpdateCents(context.Background(), cents.UpdateCentsRequest{
		UserID: likerID,
		Amount: 1,
	})
	assert.NoError(t, err)

	_, err = svc.Cents.UpdateCents(context.Background(), cents.UpdateCentsRequest{
		UserID: posterID,
		Amount: 2,
	})
	assert.NoError(t, err)

	post, err := svc.Posts.CreatePost(context.Background(), posts.CreatePostRequest{
		AuthorID:    posterID,
		Visibility:  format.PUBLIC,
		Content:     "yer mom",
		ContentType: format.TEXT,
	})
	assert.NoError(t, err)

	reply, err := svc.CreateLike(context.Background(), likes.CreateLikeRequest{
		PostID:  post.PostID,
		LikerID: likerID,
	})
	assert.NoError(t, err)
	assert.Equal(t, post.PostID, reply.PostID)
	assert.Equal(t, likerID, reply.LikerID)
	assert.False(t, reply.CreatedAt.IsZero())

	// TODO: make the error pass
	// same, err := svc.CreateLike(context.Background(), likes.CreateLikeRequest{
	// 	PostID:  post.PostID,
	// 	LikerID: likerID,
	// })
	// assert.NoError(t, err)
	// assert.Equal(t, reply.LikeID, same.LikeID)
}
