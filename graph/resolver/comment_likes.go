package resolver

import "context"

type CommentLikesGetter interface {
	CommentLikes(context.Context) ([]*CommentLike, error)
	Next(context.Context) (*string, error)
}

type CommentLikes struct {
	CommentLikesGetter
}

func NewCommentLikes(getter CommentLikesGetter) *CommentLikes {
	return &CommentLikes{CommentLikesGetter: getter}
}
