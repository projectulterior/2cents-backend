package main

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/auth"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Services struct {
	Auth *auth.Service
}

// setup services
func services(ctx context.Context, cfg Config, m *mongo.Client, log *zap.Logger) (*Services, error) {
	authService := &auth.Service{
		Secret:          cfg.Secret,
		AuthTokenTTL:    cfg.AuthTokenTTL,
		RefreshTokenTTL: cfg.RefreshTokenTTL,
		Database:        m.Database("auth"),
		Logger:          log,
	}
	if err := authService.Setup(ctx); err != nil {
		return nil, err
	}

	return &Services{
		Auth: authService,
	}, nil
}
