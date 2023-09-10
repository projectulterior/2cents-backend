package likes_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/likes"
	"github.com/stretchr/testify/assert"
)

func TestGetLikes(t *testing.T) {
	svc := setup(t)

	const (
		NUM_OF_POSTS = 10
		BATCH_SIZE   = NUM_OF_POSTS / 3
	)

	postID := format.NewPostID()

	for i := 0; i < NUM_OF_POSTS; i++ {
		_, err := svc.CreateLike(context.Background(), likes.CreateLikeRequest{
			PostID:  postID,
			LikerID: format.NewUserIDFromIdentifier(fmt.Sprintf("%d", i)),
		})
		assert.NoError(t, err)

		time.Sleep(time.Millisecond)
	}

	i := NUM_OF_POSTS - 1

	var cursor string
	for i >= 0 {
		likes, err := svc.GetLikes(context.Background(), &likes.GetLikesRequest{
			Cursor: cursor,
			Limit:  BATCH_SIZE,
		})
		assert.NoError(t, err)

		for _, like := range likes.Likes {
			expectedID := format.NewUserIDFromIdentifier(fmt.Sprintf("%d", i))
			assert.Equal(t, expectedID, like.LikerID)
			i -= 1
		}

		cursor = likes.Next
	}
}

func TestGetUserLikes(t *testing.T) {
	svc := setup(t)

	const (
		NUM_OF_LIKES = 10
		BATCH_SIZE   = NUM_OF_LIKES / 3
	)

	likerID := format.NewUserID()

	for i := 0; i < NUM_OF_LIKES; i++ {
		_, err := svc.CreateLike(context.Background(), likes.CreateLikeRequest{
			PostID:  format.NewPostIDFromIdentifier(fmt.Sprintf("%d", i)),
			LikerID: likerID,
		})
		assert.NoError(t, err)

		time.Sleep(time.Millisecond)

	}

	for i := 0; i < NUM_OF_LIKES; i++ {
		_, err := svc.CreateLike(context.Background(), likes.CreateLikeRequest{
			PostID:  format.NewPostID(),
			LikerID: format.NewUserID(),
		})
		assert.NoError(t, err)

		time.Sleep(time.Millisecond)

	}

	i := NUM_OF_LIKES - 1

	var cursor string
	for i >= 0 {
		likes, err := svc.GetLikes(context.Background(), &likes.GetLikesRequest{
			Cursor:  cursor,
			Limit:   BATCH_SIZE,
			LikerID: &likerID,
		})
		assert.NoError(t, err)

		for _, like := range likes.Likes {
			expectedID := format.NewPostIDFromIdentifier(fmt.Sprintf("%d", i))
			assert.Equal(t, expectedID, like.PostID)
			i -= 1
		}

		cursor = likes.Next
	}
}

func TestGetPostLikes(t *testing.T) {
	svc := setup(t)

	const (
		NUM_OF_LIKES = 10
		BATCH_SIZE   = NUM_OF_LIKES / 3
	)

	postID := format.NewPostID()

	for i := 0; i < NUM_OF_LIKES; i++ {
		_, err := svc.CreateLike(context.Background(), likes.CreateLikeRequest{
			PostID:  postID,
			LikerID: format.NewUserIDFromIdentifier(fmt.Sprintf("%d", i)),
		})
		assert.NoError(t, err)

		time.Sleep(time.Millisecond)

	}

	for i := 0; i < NUM_OF_LIKES; i++ {
		_, err := svc.CreateLike(context.Background(), likes.CreateLikeRequest{
			PostID:  format.NewPostID(),
			LikerID: format.NewUserID(),
		})
		assert.NoError(t, err)

		time.Sleep(time.Millisecond)

	}

	i := NUM_OF_LIKES - 1

	var cursor string
	for i >= 0 {
		likes, err := svc.GetLikes(context.Background(), &likes.GetLikesRequest{
			Cursor: cursor,
			Limit:  BATCH_SIZE,
			PostID: &postID,
		})
		assert.NoError(t, err)

		for _, like := range likes.Likes {
			expectedID := format.NewUserIDFromIdentifier(fmt.Sprintf("%d", i))
			assert.Equal(t, expectedID, like.LikerID)
			i -= 1
		}

		cursor = likes.Next
	}
}
