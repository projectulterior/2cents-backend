package posts_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/stretchr/testify/assert"
)

func TestGetPosts(t *testing.T) {
	svc := setup(t)

	const (
		NUM_OF_POSTS = 10
		BATCH_SIZE   = NUM_OF_POSTS / 3
	)

	authorID := format.NewUserID()

	for i := 0; i < NUM_OF_POSTS; i++ {
		_, err := svc.CreatePost(context.Background(), posts.CreatePostRequest{
			AuthorID:    authorID,
			Visibility:  format.PUBLIC,
			Content:     fmt.Sprintf("%d", i),
			ContentType: format.TEXT,
		})
		assert.NoError(t, err)

		time.Sleep(time.Millisecond)
	}

	i := NUM_OF_POSTS - 1

	var cursor string
	for i >= 0 {
		posts, err := svc.GetPosts(context.Background(), &posts.GetPostsRequest{
			Cursor: cursor,
			Limit:  BATCH_SIZE,
		})
		assert.NoError(t, err)

		for _, post := range posts.Posts {
			expectedContent := fmt.Sprintf("%d", i)
			assert.Equal(t, expectedContent, post.Content)
			i -= 1
		}

		cursor = posts.Next
	}
}
