package comments_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/comments"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeleteComment(t *testing.T) {
	svc := setup(t)

	postid := format.NewPostID()
	content := "twocents comment"
	newcontent1 := "new commment"
	authorID := format.NewUserID()

	reply, err := svc.CreateComment(context.Background(), comments.CreateCommentRequest{
		PostID:   postid,
		Content:  content,
		AuthorID: authorID,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, reply.CommentID)
	assert.Equal(t, postid, reply.PostID)
	assert.Equal(t, content, reply.Content)
	assert.Equal(t, authorID, reply.AuthorID)
	assert.False(t, reply.CreatedAt.IsZero())
	assert.Equal(t, reply.CreatedAt, reply.UpdatedAt)

	deleted, err := svc.DeleteComment(context.Background(), comments.DeleteCommentRequest{
		CommentID: reply.CommentID,
		DeleterID: authorID,
	})
	assert.NoError(t, err)
	assert.Equal(t, reply.CommentID, deleted.CommentID)

	_, err = svc.GetComment(context.Background(), comments.GetCommentRequest{
		CommentID: reply.CommentID,
	})
	assert.Equal(t, codes.NotFound, status.Code(err))

	reply1, err := svc.CreateComment(context.Background(), comments.CreateCommentRequest{
		PostID:   postid,
		Content:  content,
		AuthorID: authorID,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, reply1.CommentID)
	assert.Equal(t, postid, reply1.PostID)
	assert.Equal(t, content, reply1.Content)
	assert.Equal(t, authorID, reply1.AuthorID)
	assert.False(t, reply1.CreatedAt.IsZero())
	updated1, err := svc.UpdateComment(context.Background(), comments.UpdateCommentRequest{
		CommentID: reply1.CommentID,
		AuthorID:  reply1.AuthorID,
		Content:   newcontent1,
	})
	fmt.Println(err)
	assert.NoError(t, err)
	assert.Equal(t, newcontent1, updated1.Content)

	get1, err := svc.GetComment(context.Background(), comments.GetCommentRequest{
		CommentID: reply1.CommentID,
	})
	assert.NoError(t, err)
	assert.Equal(t, newcontent1, get1.Content)

	deleted1, err := svc.DeleteComment(context.Background(), comments.DeleteCommentRequest{
		CommentID: reply1.CommentID,
	})
	assert.NoError(t, err)
	assert.Equal(t, reply1.CommentID, deleted1.CommentID)

	_, err = svc.GetComment(context.Background(), comments.GetCommentRequest{
		CommentID: reply1.CommentID,
	})
	assert.Equal(t, codes.NotFound, status.Code(err))
}
