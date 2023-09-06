package resolver

import (
	"context"
	"time"

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

func NewCommentWithData(svc *services.Services, data *comments.Comment) *Comment {
	return &Comment{
		svc:       svc,
		commentID: data.CommentID,
		getter: NewGetter(
			func(ctx context.Context) (*comments.Comment, error) {
				return data, nil
			},
		),
	}
}

func (c *Comment) ID(ctx context.Context) (string, error) {
	return c.commentID.String(), nil
}

func (c *Comment) Post(ctx context.Context) (*Post, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return NewPostByID(c.svc, reply.PostID), nil
}

func (c *Comment) Content(ctx context.Context) (*string, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Content, nil
}

func (c *Comment) Author(ctx context.Context) (*User, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return NewUserByID(c.svc, reply.AuthorID), nil
}

func (c *Comment) CreatedAt(ctx context.Context) (*time.Time, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.CreatedAt, nil
}
