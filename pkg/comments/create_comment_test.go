package comments_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/cents"
	"github.com/projectulterior/2cents-backend/pkg/comments"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/stretchr/testify/assert"
)

func TestCreateComment(t *testing.T) {
	svc := setup(t)

	content := "no way!"
	commenterid := format.NewUserID()
	posterid := format.NewUserID()

	_, err := svc.Cents.UpdateCents(context.Background(), cents.UpdateCentsRequest{
		UserID: posterid,
		Amount: 2,
	})
	assert.NoError(t, err)

	_, err = svc.Cents.UpdateCents(context.Background(), cents.UpdateCentsRequest{
		UserID: commenterid,
		Amount: 1,
	})
	assert.NoError(t, err)

	post, err := svc.Posts.CreatePost(context.Background(), posts.CreatePostRequest{
		AuthorID:    posterid,
		Visibility:  format.FRIENDS,
		Content:     "im gay",
		ContentType: format.TEXT,
	})
	assert.NoError(t, err)

	reply, err := svc.CreateComment(context.Background(), comments.CreateCommentRequest{
		PostID:   post.PostID,
		Content:  content,
		AuthorID: commenterid,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, reply.CommentID)
	assert.Equal(t, post.PostID, reply.PostID)
	assert.Equal(t, content, reply.Content)
	assert.Equal(t, commenterid, reply.AuthorID)
	assert.False(t, reply.CreatedAt.IsZero())
	assert.Equal(t, reply.CreatedAt, reply.UpdatedAt)
}
