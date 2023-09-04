package comments

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Service struct {
	*mongo.Database
	*zap.Logger
}
