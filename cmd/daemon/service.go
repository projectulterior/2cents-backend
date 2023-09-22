package main

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/projectulterior/2cents-backend/pkg/auth"
	"github.com/projectulterior/2cents-backend/pkg/comment_likes"
	"github.com/projectulterior/2cents-backend/pkg/comments"
	"github.com/projectulterior/2cents-backend/pkg/follow"
	"github.com/projectulterior/2cents-backend/pkg/likes"
	"github.com/projectulterior/2cents-backend/pkg/messaging"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/projectulterior/2cents-backend/pkg/pubsub/broker"
	"github.com/projectulterior/2cents-backend/pkg/search"
	"github.com/projectulterior/2cents-backend/pkg/services"
	"github.com/projectulterior/2cents-backend/pkg/users"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// setup services
func initServices(ctx context.Context, cfg Config, m *mongo.Client, es *elasticsearch.TypedClient, log *zap.Logger) (*services.Services, error) {
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
		UserUpdated: broker.Exchange(users.UserUpdatedEvent{}).Publisher(),
		UserDeleted: broker.Exchange(users.UserDeletedEvent{}).Publisher(),

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
		Service:  postsService,
		Database: m.Database("comments"),
		Logger:   log,
	}
	if err := commentsService.Setup(ctx); err != nil {
		return nil, err
	}

	comment_LikesService := &comment_likes.Service{
		Database: m.Database("comment_likes"),
		Logger:   log,
	}
	if err := comment_LikesService.Setup(ctx); err != nil {
		return nil, err
	}

	messagingService := &messaging.Service{
		ChannelUpdated: broker.Exchange(messaging.ChannelUpdatedEvent{}).Publisher(),

		Database: m.Database("messages"),
		Logger:   log,
	}
	if err := messagingService.Setup(ctx); err != nil {
		return nil, err
	}

	searchService := &search.Service{
		UsersIndex: "users",
		PostsIndex: "posts",

		TypedClient: es,
		Logger:      log,
	}
	if err := searchService.Setup(ctx); err != nil {
		return nil, err
	}
	go broker.Exchange(auth.UserUpdatedEvent{}).Subscribe(searchService.ProcessUsernameUpdated)
	go broker.Exchange(users.UserUpdatedEvent{}).Subscribe(searchService.ProcessUserUpdated)

	return &services.Services{
		Auth:         authService,
		Users:        usersService,
		Posts:        postsService,
		Comments:     commentsService,
		Likes:        likesService,
		Follows:      followsService,
		CommentLikes: comment_LikesService,
		Messaging:    messagingService,
	}, nil
}
