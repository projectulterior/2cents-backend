package resolver

import "context"

type FollowGetter interface {
	Follows(context.Context) ([]*Follow, error)
	Next(context.Context) (*string, error)
}

type Follows struct {
	FollowGetter
}

func NewFollows(getter FollowGetter) *Follows {
	return &Follows{FollowGetter: getter}
}
