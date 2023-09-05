package posts_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/stretchr/testify/assert"
)

func TestUpdatePost(t *testing.T) {
	svc := setup(t)

	authorid := format.NewUserID()
	visibilityPublic := format.PUBLIC
	visibilityFriends := format.FRIENDS
	visibilityPrivate := format.PRIVATE
	content := "hello"
	newContent := "new comment"
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

	updatedFriends, err := svc.UpdatePost(context.Background(), posts.UpdatePostRequest{
		PostID:     reply1.PostID,
		AuthorID:   reply1.AuthorID,
		Visibility: &visibilityFriends,
	})

	assert.NoError(t, err)
	assert.Equal(t, authorid, updatedFriends.AuthorID)
	assert.Equal(t, visibilityFriends, updatedFriends.Visibility)
	assert.NotEqual(t, updatedFriends.UpdatedAt, updatedFriends.CreatedAt)
	assert.Equal(t, content, updatedFriends.Content)
	assert.Equal(t, contentType, updatedFriends.ContentType)
	assert.NotEmpty(t, updatedFriends.PostID)
	assert.False(t, updatedFriends.CreatedAt.IsZero())

	updatedPrivate, err := svc.UpdatePost(context.Background(), posts.UpdatePostRequest{
		PostID:     reply1.PostID,
		AuthorID:   reply1.AuthorID,
		Visibility: &visibilityPrivate,
	})

	assert.NoError(t, err)
	assert.Equal(t, authorid, updatedPrivate.AuthorID)
	assert.Equal(t, visibilityPrivate, updatedPrivate.Visibility)
	assert.Equal(t, content, updatedPrivate.Content)
	assert.Equal(t, contentType, updatedPrivate.ContentType)
	assert.NotEmpty(t, updatedPrivate.PostID)
	assert.False(t, updatedPrivate.CreatedAt.IsZero())

	updatedContent, err := svc.UpdatePost(context.Background(), posts.UpdatePostRequest{
		PostID:   reply1.PostID,
		AuthorID: reply1.AuthorID,
		Content:  &newContent,
	})

	assert.NoError(t, err)
	assert.Equal(t, authorid, updatedContent.AuthorID)
	assert.Equal(t, visibilityPrivate, updatedContent.Visibility)
	assert.Equal(t, newContent, updatedContent.Content)
	assert.Equal(t, contentType, updatedContent.ContentType)
	assert.NotEmpty(t, updatedContent.PostID)
	assert.False(t, updatedContent.CreatedAt.IsZero())
}
