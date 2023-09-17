package resolver

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/comments"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type UserComments struct {
	svc *services.Services
	getter[*comments.GetCommentsResponse, func(context.Context) (*comments.GetCommentsResponse, error)]
}

func NewUserComments(svc *services.Services, userID format.UserID, page Pagination) *UserComments {
	return &UserComments{
		svc: svc,
		getter: NewGetter(
			func(ctx context.Context) (*comments.GetCommentsResponse, error) {
				return svc.Comments.GetComments(ctx, &comments.GetCommentsRequest{
					AuthorID: userID,
					Cursor:   page.Cursor,
					Limit:    page.Limit,
				})
			},
		),
	}
}

func (c *UserComments) Comments(ctx context.Context) ([]*Comment, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	var toRet []*Comment
	for _, comment := range reply.Comments {
		toRet = append(toRet, NewCommentWithData(c.svc, comment))
	}

	return toRet, nil
}

func (c *UserComments) Next(ctx context.Context) (*string, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Next, nil
}
