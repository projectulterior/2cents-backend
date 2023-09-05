package services

import (
	"github.com/projectulterior/2cents-backend/pkg/auth"
	"github.com/projectulterior/2cents-backend/pkg/users"
)

type Services struct {
	Auth  *auth.Service
	Users *users.Service
}
