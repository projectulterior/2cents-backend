package comments_test

import (
	"context"
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

	get1, err := svc.GetComment(context.Background(), comments.GetCommentRequest{
		CommentID: reply.CommentID,
	})
	assert.NoError(t, err)
	assert.Equal(t, content, get1.Content)

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
}

// TODO:
func TestDeleteComment_PostAuthor(t *testing.T) {
	svc := setup(t)

	postid := format.NewPostID()
	content := "twocents comment"
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

	get1, err := svc.GetComment(context.Background(), comments.GetCommentRequest{
		CommentID: reply.CommentID,
	})
	assert.NoError(t, err)
	assert.Equal(t, content, get1.Content)

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
}
