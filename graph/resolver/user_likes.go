package resolver

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/likes"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type UserLikes struct {
	svc *services.Services
	getter[*likes.GetLikesResponse, func(context.Context) (*likes.GetLikesResponse, error)]
}

func NewUserLikes(svc *services.Services, userID format.UserID, page Pagination) *UserLikes {
	return &UserLikes{
		svc: svc,
		getter: NewGetter(
			func(ctx context.Context) (*likes.GetLikesResponse, error) {
				return svc.Likes.GetLikes(ctx, &likes.GetLikesRequest{
					LikerID: &userID,
					Cursor:  page.Cursor,
					Limit:   page.Limit,
				})
			},
		),
	}
}

func (l *UserLikes) Likes(ctx context.Context) ([]*Like, error) {
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

func (l *UserLikes) Next(ctx context.Context) (*string, error) {
	reply, err := l.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Next, nil
}
