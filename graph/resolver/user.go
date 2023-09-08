package resolver

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/auth"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/services"
	"github.com/projectulterior/2cents-backend/pkg/users"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	svc *services.Services

	userID format.UserID
	auth   getter[*auth.User, func(context.Context) (*auth.User, error)]
	getter[*users.User, func(context.Context) (*users.User, error)]
}

func newUser(svc *services.Services, userID format.UserID) *User {
	return &User{
		svc:    svc,
		userID: userID,
		auth: NewGetter(
			func(ctx context.Context) (*auth.User, error) {
				return svc.Auth.GetUser(ctx, auth.GetUserRequest{
					UserID: userID,
				})
			},
		),
	}
}

func NewUserByID(svc *services.Services, userID format.UserID) *User {
	user := newUser(svc, userID)

	user.getter = NewGetter(
		func(ctx context.Context) (*users.User, error) {
			return svc.Users.GetUser(ctx, users.GetUserRequest{
				UserID: userID,
			})
		},
	)

	return user
}

func NewMyUser(svc *services.Services, userID format.UserID) *User {
	user := newUser(svc, userID)

	user.getter = NewGetter(
		func(ctx context.Context) (*users.User, error) {
			user, err := svc.Users.GetUser(ctx, users.GetUserRequest{
				UserID: userID,
			})
			if status.Code(err) != codes.NotFound {
				return user, err
			}

			return svc.Users.CreateUser(ctx, users.CreateUserRequest{UserID: userID})
		},
	)

	return user
}

func NewUserWithData(svc *services.Services, data *users.User) *User {
	user := newUser(svc, data.UserID)

	user.getter = NewGetter(
		func(ctx context.Context) (*users.User, error) {
			return data, nil
		},
	)

	return user
}

func (u *User) ID(ctx context.Context) (string, error) {
	return u.userID.String(), nil
}

func (u *User) Username(ctx context.Context) (*string, error) {
	reply, err := u.auth.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Username, nil
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

func (u *User) Profile(ctx context.Context) (*string, error) {
	reply, err := u.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Profile, nil
}

func (u *User) Cover(ctx context.Context) (*string, error) {
	reply, err := u.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Cover, nil
}
