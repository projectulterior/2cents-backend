package resolver

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/follow"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type UserFollowers struct {
	svc *services.Services
	getter[*follow.GetFollowsResponse, func(context.Context) (*follow.GetFollowsResponse, error)]
}

func NewUserFollowers(svc *services.Services, userID format.UserID, page Pagination) *UserFollowers {
	return &UserFollowers{
		svc: svc,
		getter: NewGetter(
			func(ctx context.Context) (*follow.GetFollowsResponse, error) {
				return svc.Follows.GetFollows(ctx, &follow.GetFollowsRequest{
					FolloweeID: &userID,
					Cursor:     page.Cursor,
					Limit:      page.Limit,
				})
			},
		),
	}
}

func (f *UserFollowers) Follows(ctx context.Context) ([]*Follow, error) {
	reply, err := f.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	var toRet []*Follow
	for _, follow := range reply.Follows {
		toRet = append(toRet, NewFollowWithData(f.svc, follow))
	}

	return toRet, nil
}

func (f *UserFollowers) Next(ctx context.Context) (*string, error) {
	reply, err := f.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Next, nil
}
