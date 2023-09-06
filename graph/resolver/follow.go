package resolver

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/follow"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type Follow struct {
	svc *services.Services

	followID format.FollowID
	getter[*follow.Follow, func(context.Context) (*follow.Follow, error)]
}

func NewFollowByID(svc *services.Services, followID format.FollowID) *Follow {
	return &Follow{
		svc:      svc,
		followID: followID,
		getter: NewGetter(
			func(ctx context.Context) (*follow.Follow, error) {
				return svc.Follows.GetFollow(ctx, follow.GetFollowRequest{})
			},
		),
	}
}

func NewFollowWithData(svc *services.Services, data *follow.Follow) *Follow {
	return &Follow{
		svc:      svc,
		followID: data.FollowID,
		getter: NewGetter(
			func(ctx context.Context) (*follow.Follow, error) {
				return data, nil
			},
		),
	}
}

func (f *Follow) ID(ctx context.Context) (string, error) {
	return f.followID.String(), nil
}

func (f *Follow) Follower(ctx context.Context) (*User, error) {
	reply, err := f.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return NewUserByID(f.svc, reply.FollowerID), nil
}

func (f *Follow) Followee(ctx context.Context) (*User, error) {
	reply, err := f.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return NewUserByID(f.svc, reply.FolloweeID), nil
}

func (f *Follow) CreatedAt(ctx context.Context) (*time.Time, error) {
	reply, err := f.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.CreatedAt, nil
}
