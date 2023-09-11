package resolver

import "context"

type PostGetter interface {
	Posts(ctx context.Context) ([]*Post, error)
	Next(ctx context.Context) (*string, error)
}

type Posts struct {
	PostGetter
}

func NewPosts(getter PostGetter) *Posts {
	return &Posts{PostGetter: getter}
}
