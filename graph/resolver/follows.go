package resolver

import "context"

type FollowsGetter interface {
	Likes(context.Context) ([]*Follow, error)
	Next(context.Context) (*string, error)
}

type Follows struct {
	FollowsGetter
}

func NewFollows(getter FollowsGetter) *Follows {
	return &Follows{FollowsGetter: getter}
}
