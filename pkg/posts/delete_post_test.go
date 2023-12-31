package posts_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeletePost(t *testing.T) {
	svc := setup(t)

	authorid := format.NewUserID()
	visibilityPublic := format.PUBLIC
	content := "hello"
	contentType := format.TEXT

	reply1, err := svc.CreatePost(context.Background(), posts.CreatePostRequest{
		AuthorID:    authorid,
		Visibility:  visibilityPublic,
		Content:     content,
		ContentType: contentType,
	})

	assert.NoError(t, err)
	assert.Equal(t, authorid, reply1.AuthorID)
	assert.Equal(t, visibilityPublic, reply1.Visibility)
	assert.Equal(t, content, reply1.Content)
	assert.Equal(t, contentType, reply1.ContentType)
	assert.NotEmpty(t, reply1.PostID)
	assert.False(t, reply1.CreatedAt.IsZero())
	assert.Equal(t, reply1.CreatedAt, reply1.UpdatedAt)

	deleted, err := svc.DeletePost(context.Background(), posts.DeletePostRequest{
		PostID:   reply1.PostID,
		AuthorID: authorid,
	})

	assert.NoError(t, err)
	assert.Equal(t, reply1.PostID, deleted.PostID)

	_, err = svc.GetPost(context.Background(), posts.GetPostRequest{
		PostID: reply1.PostID,
	})
	assert.Equal(t, codes.NotFound, status.Code(err))
}
