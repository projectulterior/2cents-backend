package comments_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/comments"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/stretchr/testify/assert"
)

func TestUpdateComment(t *testing.T) {
	svc := setup(t)

	postid := format.NewPostID()
	content := "twocents comment"
	newcontent1 := "new commment"
	newcontent2 := "sbk"
	authorid := format.NewUserID()

	reply, err := svc.CreateComment(context.Background(), comments.CreateCommentRequest{
		PostID:   postid,
		Content:  content,
		AuthorID: authorid,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, reply.CommentID)
	assert.Equal(t, postid, reply.PostID)
	assert.Equal(t, content, reply.Content)
	assert.Equal(t, authorid, reply.AuthorID)
	assert.False(t, reply.CreatedAt.IsZero())
	assert.Equal(t, reply.CreatedAt, reply.UpdatedAt)

	updated1, err := svc.UpdateComment(context.Background(), comments.UpdateCommentRequest{
		CommentID: reply.CommentID,
		AuthorID:  reply.AuthorID,
		Content:   newcontent1,
	})
	fmt.Println(err)
	assert.NoError(t, err)
	assert.Equal(t, newcontent1, updated1.Content)

	get1, err := svc.GetComment(context.Background(), comments.GetCommentRequest{
		CommentID: reply.CommentID,
	})
	assert.NoError(t, err)
	assert.Equal(t, newcontent1, get1.Content)

	updated2, err := svc.UpdateComment(context.Background(), comments.UpdateCommentRequest{
		CommentID: reply.CommentID,
		AuthorID:  reply.AuthorID,
		Content:   newcontent2,
	})
	assert.NoError(t, err)
	assert.Equal(t, newcontent2, updated2.Content)

	get2, err := svc.GetComment(context.Background(), comments.GetCommentRequest{
		CommentID: reply.CommentID,
	})
	assert.NoError(t, err)
	assert.Equal(t, newcontent2, get2.Content)
}
