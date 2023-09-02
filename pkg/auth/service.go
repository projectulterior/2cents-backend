package auth

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Service struct {
	Secret string

	*mongo.Database
	*zap.Logger
}
