package resolver

import "context"

type UsersGetter interface {
	Users(context.Context) ([]*User, error)
	Next(context.Context) (*string, error)
}

type Users struct {
	UsersGetter
}

func NewUsers(getter UsersGetter) *Users {
	return &Users{UsersGetter: getter}
}
