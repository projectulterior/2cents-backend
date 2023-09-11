package resolver

import "context"

type PostsGetter interface {
	Posts(context.Context) ([]*Post, error)
	Next(context.Context) (*string, error)
}

type Posts struct {
	PostsGetter
}

func NewPosts(getter PostsGetter) *Posts {
	return &Posts{PostsGetter: getter}
}
