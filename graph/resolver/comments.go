package resolver

import "context"

type CommentsGetter interface {
	Comments(context.Context) ([]*Comment, error)
	Next(context.Context) (*string, error)
}

type Comments struct {
	CommentsGetter
}

func NewComments(getter CommentsGetter) *Comments {
	return &Comments{CommentsGetter: getter}
}
