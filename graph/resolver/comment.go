package resolver

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/comments"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type Comment struct {
	svc *services.Services

	commentID format.CommentID
	getter[*comments.Comment, func(context.Context) (*comments.Comment, error)]
}

func NewCommentByID(svc *services.Services, commentID format.CommentID) *Comment {
	return &Comment{
		svc:       svc,
		commentID: commentID,
		getter: NewGetter(
			func(ctx context.Context) (*comments.Comment, error) {
				return svc.Comments.GetComment(ctx, comments.GetCommentRequest{
					CommentID: commentID,
				})
			},
		),
	}
}

func (c *Comment) ID(ctx context.Context) (string, error) {
	return c.commentID.String(), nil
}

// func (c *Comment) PostID(ctx context.Context) (string, error) {
// 	reply, err := c.getter.Call(ctx)
// 	if err != nil {
// 		return "", err
// 	}

// 	return reply.PostID, nil
// }
