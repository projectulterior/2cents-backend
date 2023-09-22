package auth

import (
	"time"

	"github.com/projectulterior/2cents-backend/pkg/pubsub"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Service struct {
	Secret          string
	AuthTokenTTL    time.Duration
	RefreshTokenTTL time.Duration

	UserUpdated pubsub.Publisher[UserUpdatedEvent]

	*mongo.Database
	*zap.Logger
}
