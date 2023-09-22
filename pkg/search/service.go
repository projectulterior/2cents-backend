package search

import (
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
)

type Service struct {
	UsersIndex string
	PostsIndex string

	*elasticsearch.TypedClient
	*zap.Logger
}
