package resolver

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type Post struct {
	svc *services.Services

	postID format.PostID
	getter[*posts.Post, func(context.Context) (*posts.Post, error)]
}

func NewPostByID(svc *services.Services, postID format.PostID) *Post {
	return &Post{
		svc:    svc,
		postID: postID,
		getter: NewGetter(
			func(ctx context.Context) (*posts.Post, error) {
				return svc.Posts.GetPost(ctx, posts.GetPostRequest{
					PostID: postID,
				})
			},
		),
	}
}

func NewPostWithData(svc *services.Services, data *posts.Post) *Post {
	return &Post{
		svc:    svc,
		postID: data.PostID,
		getter: NewGetter(
			func(ctx context.Context) (*posts.Post, error) {
				return data, nil
			},
		),
	}
}

func (p *Post) ID(ctx context.Context) (string, error) {
	return p.postID.String(), nil
}

func (p *Post) Author(ctx context.Context) (*User, error) {
	reply, err := p.getter.Call(ctx)
	if err != nil {
		return nil, err
	}
	return NewUserByID(p.svc, reply.AuthorID), nil
}

func (p *Post) Visibility(ctx context.Context) (*format.Visibility, error) {
	reply, err := p.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Visibility, nil
}

func (p *Post) Content(ctx context.Context) (*string, error) {
	reply, err := p.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Content, nil
}

func (p *Post) ContentType(ctx context.Context) (*format.ContentType, error) {
	reply, err := p.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.ContentType, nil
}

func (p *Post) CreatedAt(ctx context.Context) (*time.Time, error) {
	reply, err := p.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.CreatedAt, nil
}

func (p *Post) UpdatedAt(ctx context.Context) (*time.Time, error) {
	reply, err := p.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.UpdatedAt, nil
}

func (p *Post) Likes(ctx context.Context, page Pagination) (*Likes, error) {
	return NewLikes(NewPostLikes(p.svc, p.postID, page)), nil
}

func (p *Post) Comments(ctx context.Context, page Pagination) (*Comments, error) {
	return NewComments(NewPostComments(p.svc, p.postID, page)), nil
}
