package resolver

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/likes"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type PostLikes struct {
	svc *services.Services
	getter[*likes.GetLikesResponse, func(context.Context) (*likes.GetLikesResponse, error)]
}

func NewPostLikes(svc *services.Services, postID format.PostID, page Pagination) *PostLikes {
	return &PostLikes{
		svc: svc,
		getter: NewGetter(
			func(ctx context.Context) (*likes.GetLikesResponse, error) {
				return svc.Likes.GetLikes(ctx, &likes.GetLikesRequest{
					PostID: &postID,
					Cursor: page.Cursor,
					Limit:  page.Limit,
				})
			},
		),
	}
}

func (l *PostLikes) Likes(ctx context.Context) ([]*Like, error) {
	reply, err := l.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	var toRet []*Like
	for _, like := range reply.Likes {
		toRet = append(toRet, NewLikeWithData(l.svc, like))
	}

	return toRet, nil
}

func (l *PostLikes) Next(ctx context.Context) (*string, error) {
	reply, err := l.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Next, nil
}
