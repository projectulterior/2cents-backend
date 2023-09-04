package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Service struct {
	Secret          string
	AuthTokenTTL    time.Duration
	RefreshTokenTTL time.Duration

	*mongo.Database
	*zap.Logger
}
