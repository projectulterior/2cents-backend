package main

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/auth"
)

type Services struct {
	Auth *auth.Service
}

func services(ctx context.Context) (*Services, error) {
	authService := &auth.Service{}
	if err := authService.Setup(ctx); err != nil {
		return nil, err
	}

	return &Services{
		Auth: authService,
	}, nil
}
