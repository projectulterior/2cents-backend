package resolver

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/comment_likes"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type CommentLike struct {
	svc *services.Services

	commentLikeID format.CommentLikeID
	getter[*comment_likes.GetCommentLikeResponse, func(context.Context) (*comment_likes.GetCommentLikeResponse, error)]
}

func NewCommentLikeByID(svc *services.Services, commentLikeID format.CommentLikeID) *CommentLike {
	return &CommentLike{
		svc:           svc,
		commentLikeID: commentLikeID,
		getter: NewGetter(
			func(ctx context.Context) (*comment_likes.CommentLike, error) {
				return svc.CommentLikes.GetCommentLike(ctx, comment_likes.GetCommentLikeRequest{
					CommentLikeID: commentLikeID,
				})
			},
		),
	}
}

func NewCommentLikeWithData(svc *services.Services, data *comment_likes.CommentLike) *CommentLike {
	return &CommentLike{
		svc:           svc,
		commentLikeID: data.CommentLikeID,
		getter: NewGetter(
			func(ctx context.Context) (*comment_likes.CommentLike, error) {
				return data, nil
			},
		),
	}
}

func (c *CommentLike) ID(ctx context.Context) (string, error) {
	return c.commentLikeID.String(), nil
}

func (c *CommentLike) Comment(ctx context.Context) (*Comment, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return NewCommentByID(c.svc, reply.CommentID), nil
}

func (c *CommentLike) Liker(ctx context.Context) (*User, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return NewUserByID(c.svc, reply.LikerID), nil
}

func (c *CommentLike) CreatedAt(ctx context.Context) (*time.Time, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.CreatedAt, nil
}
