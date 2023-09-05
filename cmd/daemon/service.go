package main

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/auth"
	"github.com/projectulterior/2cents-backend/pkg/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// setup services
func initServices(ctx context.Context, cfg Config, m *mongo.Client, log *zap.Logger) (*services.Services, error) {
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

	return &services.Services{
		Auth: authService,
	}, nil
}
