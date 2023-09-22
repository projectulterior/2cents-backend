package users

import (
	"github.com/projectulterior/2cents-backend/pkg/pubsub"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Service struct {
	UserUpdated pubsub.Publisher[UserUpdatedEvent]
	UserDeleted pubsub.Publisher[UserDeletedEvent]

	*mongo.Database
	*zap.Logger
}
