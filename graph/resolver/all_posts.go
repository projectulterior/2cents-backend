package resolver

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type AllPosts struct {
	svc *services.Services
	getter[*posts.GetPostsResponse, func(context.Context) (*posts.GetPostsResponse, error)]
}

func NewAllPosts(svc *services.Services, page Pagination) *AllPosts {
	return &AllPosts{
		svc: svc,
		getter: NewGetter(
			func(ctx context.Context) (*posts.GetPostsResponse, error) {
				return svc.Posts.GetPosts(ctx, &posts.GetPostsRequest{
					Cursor: page.Cursor,
					Limit:  page.Limit,
				})
			},
		),
	}
}

func (p *AllPosts) Posts(ctx context.Context) ([]*Post, error) {
	reply, err := p.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	var toRet []*Post
	for _, post := range reply.Posts {
		toRet = append(toRet, NewPostWithData(p.svc, post))
	}

	return toRet, nil
}

func (p *AllPosts) Next(ctx context.Context) (*string, error) {
	reply, err := p.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Next, nil
}
