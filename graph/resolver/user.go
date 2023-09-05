package resolver

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/services"
	"github.com/projectulterior/2cents-backend/pkg/users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	svc *services.Services

	userID format.UserID
	getter[*users.User, func(context.Context) (*users.User, error)]
}

func NewUserByID(svc *services.Services, userID format.UserID) *User {
	return &User{
		svc:    svc,
		userID: userID,
		getter: NewGetter(
			func(ctx context.Context) (*users.User, error) {
				return svc.Users.GetUser(ctx, users.GetUserRequest{
					UserID: userID,
				})
			},
		),
	}
}

func NewMyUser(svc *services.Services, userID format.UserID) *User {
	return &User{
		svc:    svc,
		userID: userID,
		getter: NewGetter(
			func(ctx context.Context) (*users.User, error) {
				user, err := svc.Users.GetUser(ctx, users.GetUserRequest{
					UserID: userID,
				})
				if status.Code(err) != codes.NotFound {
					return user, err
				}

				return svc.Users.CreateUser(ctx, users.CreateUserRequest{UserID: userID})
			},
		),
	}
}

func NewUserWithData(svc *services.Services, data *users.User) *User {
	return &User{
		svc:    svc,
		userID: data.UserID,
		getter: NewGetter(
			func(ctx context.Context) (*users.User, error) {
				return data, nil
			},
		),
	}
}

func (u *User) ID(ctx context.Context) (string, error) {
	return u.userID.String(), nil
}

func (u *User) Name(ctx context.Context) (string, error) {
	reply, err := u.getter.Call(ctx)
	if err != nil {
		return "", err
	}

	return reply.Name, nil
}

func (u *User) Bio(ctx context.Context) (string, error) {
	reply, err := u.getter.Call(ctx)
	if err != nil {
		return "", err
	}

	return reply.Bio, nil
}
