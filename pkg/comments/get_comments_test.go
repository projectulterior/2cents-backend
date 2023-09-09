package comments_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/comments"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/stretchr/testify/assert"
)

func TestGetComments(t *testing.T) {
	svc := setup(t)

	const (
		NUM_OF_COMMENTS = 10
		BATCH_SIZE      = NUM_OF_COMMENTS / 3
	)

	for i := 0; i < NUM_OF_COMMENTS; i++ {
		_, err := svc.CreateComment(context.Background(), comments.CreateCommentRequest{
			PostID:  format.NewPostID(),
			Content: string(format.NewCommentIDFromIdentifier(fmt.Sprintf("%d", i))),
		})
		assert.NoError(t, err)
	}

	i := NUM_OF_COMMENTS - 1

	var cursor string
	for i >= 0 {
		comments, err := svc.GetComments(context.Background(), &comments.GetCommentsRequest{
			Cursor: cursor,
			Limit:  BATCH_SIZE,
		})
		assert.NoError(t, err)

		for _, comment := range comments.Comments {
			expectedID := format.NewCommentIDFromIdentifier(fmt.Sprintf("%d", i))
			assert.Equal(t, expectedID, comment.CommentID)
			i -= 1
		}

		cursor = comments.Next
	}
}
