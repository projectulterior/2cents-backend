package posts

import (
	"github.com/projectulterior/2cents-backend/pkg/cents"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Service struct {
	Cents *cents.Service

	*mongo.Database
	*zap.Logger
}
