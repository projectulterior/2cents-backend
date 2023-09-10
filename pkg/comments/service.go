package comments

import (
	"github.com/projectulterior/2cents-backend/pkg/posts"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Service struct {
	*posts.Service

	*mongo.Database
	*zap.Logger
}
