package resolver

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/comment_likes"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type CommentsCommentLikes struct {
	svc *services.Services
	getter[*comment_likes.GetCommentLikesResponse, func(context.Context) (*comment_likes.GetCommentLikesResponse, error)]
}

func NewCommentsCommentLikes(svc *services.Services, commentID format.CommentID, page Pagination) *CommentsCommentLikes {
	return &CommentsCommentLikes{
		svc: svc,
		getter: NewGetter(
			func(ctx context.Context) (*comment_likes.GetCommentLikesResponse, error) {
				return svc.CommentLikes.GetCommentLikes(ctx, &comment_likes.GetCommentLikesRequest{
					CommentID: &commentID,
					Cursor:    page.Cursor,
					Limit:     page.Limit,
				})
			},
		),
	}
}

func (c *CommentsCommentLikes) CommentLikes(ctx context.Context) ([]*CommentLike, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	var toRet []*CommentLike
	for _, commentLike := range reply.CommentLikes {
		toRet = append(toRet, NewCommentLikeWithData(c.svc, commentLike))
	}

	return toRet, nil
}

func (c *CommentsCommentLikes) Next(ctx context.Context) (*string, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Next, nil
}
