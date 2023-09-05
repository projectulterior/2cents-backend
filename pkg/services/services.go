package services

import (
	"github.com/projectulterior/2cents-backend/pkg/auth"
	"github.com/projectulterior/2cents-backend/pkg/comments"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/projectulterior/2cents-backend/pkg/users"
)

type Services struct {
	Auth     *auth.Service
	Users    *users.Service
	Posts    *posts.Service
	Comments *comments.Service
}
