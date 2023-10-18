package likes

import (
	"github.com/projectulterior/2cents-backend/pkg/cents"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Service struct {
	Cents *cents.Service
	Posts *posts.Service

	*mongo.Database
	*zap.Logger
}
