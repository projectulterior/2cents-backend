package resolver

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/comment_likes"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type UserCommentLikes struct {
	svc *services.Services
	getter[*comment_likes.GetCommentLikesResponse, func(context.Context) (*comment_likes.GetCommentLikesResponse, error)]
}

func NewUserCommentLikes(svc *services.Services, likerID format.UserID, page Pagination) *UserCommentLikes {
	return &UserCommentLikes{
		svc: svc,
		getter: NewGetter(
			func(ctx context.Context) (*comment_likes.GetCommentLikesResponse, error) {
				return svc.CommentLikes.GetCommentLikes(ctx, &comment_likes.GetCommentLikesRequest{
					LikerID: &likerID,
					Cursor:  page.Cursor,
					Limit:   page.Limit,
				})
			},
		),
	}
}

func (u *UserCommentLikes) CommentLikes(ctx context.Context) ([]*CommentLike, error) {
	reply, err := u.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	var toRet []*CommentLike
	for _, commentLike := range reply.CommentLikes {
		toRet = append(toRet, NewCommentLikeWithData(u.svc, commentLike))
	}

	return toRet, nil
}

func (u *UserCommentLikes) Next(ctx context.Context) (*string, error) {
	reply, err := u.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Next, nil
}
