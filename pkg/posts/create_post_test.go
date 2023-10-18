package posts_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/cents"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	svc := setup(t)

	createPost(t, svc, format.NewUserID())
}

func createPost(t *testing.T, svc *posts.Service, authorID format.UserID) *posts.Post {
	_, err := svc.Cents.UpdateCents(context.Background(), cents.UpdateCentsRequest{
		UserID: authorID,
		Amount: 2,
	})
	assert.NoError(t, err)

	visibilityPublic := format.PUBLIC
	content := "hello"
	contentType := format.TEXT

	post, err := svc.CreatePost(context.Background(), posts.CreatePostRequest{
		AuthorID:    authorID,
		Visibility:  visibilityPublic,
		Content:     content,
		ContentType: contentType,
	})

	assert.NoError(t, err)
	assert.Equal(t, authorID, post.AuthorID)
	assert.Equal(t, visibilityPublic, post.Visibility)
	assert.Equal(t, content, post.Content)
	assert.Equal(t, contentType, post.ContentType)
	assert.NotEmpty(t, post.PostID)
	assert.False(t, post.CreatedAt.IsZero())
	assert.Equal(t, post.CreatedAt, post.UpdatedAt)

	return post
}
