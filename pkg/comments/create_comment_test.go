package comments_test

import (
	"context"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/comments"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/stretchr/testify/assert"
)

func TestCreateComment(t *testing.T) {
	svc := setup(t)

	postid := format.NewPostID()
	content := "twocents comment"
	authorid := format.NewUserID()

	reply, err := svc.CreateComment(context.Background(), comments.CreateCommentRequest{
		PostID:   postid,
		Content:  content,
		AuthorID: authorid,
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, reply.CommentID)
	assert.Equal(t, postid, reply.PostID)
	assert.Equal(t, content, reply.Content)
	assert.Equal(t, authorid, reply.AuthorID)
	assert.False(t, reply.CreatedAt.IsZero())
	assert.Equal(t, reply.CreatedAt, reply.UpdatedAt)
}
