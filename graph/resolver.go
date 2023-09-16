package graph

import (
	"github.com/projectulterior/2cents-backend/pkg/pubsub"
	"github.com/projectulterior/2cents-backend/pkg/services"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	*services.Services
	pubsub.Broker

	*zap.Logger
}
