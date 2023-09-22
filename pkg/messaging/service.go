package messaging

import (
	"github.com/projectulterior/2cents-backend/pkg/pubsub"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Service struct {
	ChannelUpdated pubsub.Publisher[ChannelUpdatedEvent]

	*mongo.Database
	*zap.Logger
}
