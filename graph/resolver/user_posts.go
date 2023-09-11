package resolver

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type UserPosts struct {
	svc *services.Services
	getter[*posts.GetPostsResponse, func(context.Context) (*posts.GetPostsResponse, error)]
}

func NewUserPosts(svc *services.Services, userID format.UserID, page Pagination) *UserPosts {
	return &UserPosts{
		svc: svc,
		getter: NewGetter(
			func(ctx context.Context) (*posts.GetPostsResponse, error) {
				return svc.Posts.GetPosts(ctx, &posts.GetPostsRequest{
					AuthorID: &userID,
					Cursor:   page.Cursor,
					Limit:    page.Limit,
				})
			},
		),
	}
}

func (p *UserPosts) Posts(ctx context.Context) ([]*Post, error) {
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

func (p *UserPosts) Next(ctx context.Context) (*string, error) {
	reply, err := p.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Next, nil
}
