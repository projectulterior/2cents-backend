package resolver

import "context"

type LikesGetter interface {
	Likes(context.Context) ([]*Like, error)
	Next(context.Context) (*string, error)
}

type Likes struct {
	LikesGetter
}

func NewLikes(getter LikesGetter) *Likes {
	return &Likes{LikesGetter: getter}
}
