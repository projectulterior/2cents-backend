package resolver

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/search"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type SearchUsers struct {
	svc *services.Services
	getter[*search.GetUsersResponse, func(context.Context) (*search.GetUsersResponse, error)]
}

func NewSearchUsers(svc *services.Services, query string, page Pagination) *SearchUsers {
	return &SearchUsers{
		svc: svc,
		getter: NewGetter(
			func(ctx context.Context) (*search.GetUsersResponse, error) {
				return svc.Search.GetUsers(ctx, search.GetUsersRequest{
					Query:  query,
					Cursor: page.Cursor,
					Limit:  page.Limit,
				})
			},
		),
	}
}

func (u *SearchUsers) Users(ctx context.Context) ([]*User, error) {
	reply, err := u.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	var toRet []*User
	for _, user := range reply.Users {
		toRet = append(toRet, NewUserByID(u.svc, user.UserID))
	}

	return toRet, nil
}

func (u *SearchUsers) Next(ctx context.Context) (*string, error) {
	reply, err := u.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Next, nil
}
