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

	assert.Nil(t, err)
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

	assert.Nil(t, err)
	assert.Equal(t, authorid, updatedFriends.Post.AuthorID)
	assert.Equal(t, visibilityFriends, updatedFriends.Post.Visibility)
	assert.NotEqual(t, updatedFriends.Post.UpdatedAt, updatedFriends.Post.CreatedAt)
	assert.Equal(t, content, updatedFriends.Post.Content)
	assert.Equal(t, contentType, updatedFriends.Post.ContentType)
	assert.NotEmpty(t, updatedFriends.Post.PostID)
	assert.False(t, updatedFriends.Post.CreatedAt.IsZero())

	updatedPrivate, err := svc.UpdatePost(context.Background(), posts.UpdatePostRequest{
		PostID:     reply1.PostID,
		AuthorID:   reply1.AuthorID,
		Visibility: &visibilityPrivate,
	})

	assert.Nil(t, err)
	assert.Equal(t, authorid, updatedPrivate.Post.AuthorID)
	assert.Equal(t, visibilityPrivate, updatedPrivate.Post.Visibility)
	assert.Equal(t, content, updatedPrivate.Post.Content)
	assert.Equal(t, contentType, updatedPrivate.Post.ContentType)
	assert.NotEmpty(t, updatedPrivate.Post.PostID)
	assert.False(t, updatedPrivate.Post.CreatedAt.IsZero())

	updatedContent, err := svc.UpdatePost(context.Background(), posts.UpdatePostRequest{
		PostID:   reply1.PostID,
		AuthorID: reply1.AuthorID,
		Content:  &newContent,
	})

	assert.Nil(t, err)
	assert.Equal(t, authorid, updatedContent.Post.AuthorID)
	assert.Equal(t, visibilityPrivate, updatedContent.Post.Visibility)
	assert.Equal(t, newContent, updatedContent.Post.Content)
	assert.Equal(t, contentType, updatedContent.Post.ContentType)
	assert.NotEmpty(t, updatedContent.Post.PostID)
	assert.False(t, updatedContent.Post.CreatedAt.IsZero())
}
