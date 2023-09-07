package main

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/auth"
	"github.com/projectulterior/2cents-backend/pkg/comments"
	"github.com/projectulterior/2cents-backend/pkg/follow"
	"github.com/projectulterior/2cents-backend/pkg/likes"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/projectulterior/2cents-backend/pkg/services"
	"github.com/projectulterior/2cents-backend/pkg/users"

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

	usersService := &users.Service{
		Database: m.Database("users"),
		Logger:   log,
	}
	if err := usersService.Setup(ctx); err != nil {
		return nil, err
	}

	postsService := &posts.Service{
		Database: m.Database("posts"),
		Logger:   log,
	}
	if err := postsService.Setup(ctx); err != nil {
		return nil, err
	}

	likesService := &likes.Service{
		Database: m.Database("likes"),
		Logger:   log,
	}
	if err := likesService.Setup(ctx); err != nil {
		return nil, err
	}

	followsService := &follow.Service{
		Database: m.Database("follows"),
		Logger:   log,
	}
	if err := followsService.Setup(ctx); err != nil {
		return nil, err
	}
	commentsService := &comments.Service{
		Database: m.Database("comments"),
		Logger:   log,
		Service:  postsService,
	}
	if err := followsService.Setup(ctx); err != nil {
		return nil, err
	}

	return &services.Services{
		Auth:     authService,
		Users:    usersService,
		Posts:    postsService,
		Likes:    likesService,
		Follows:  followsService,
		Comments: commentsService,
	}, nil
}
