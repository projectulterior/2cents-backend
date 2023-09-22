package search_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/projectulterior/2cents-backend/pkg/auth"
	"github.com/projectulterior/2cents-backend/pkg/logger"
	"github.com/projectulterior/2cents-backend/pkg/pubsub"
	"github.com/projectulterior/2cents-backend/pkg/search"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	Service string `envconfig:"SERVICE" default:"search"`
	Elastic string `envconfig:"ELASTIC" default:"http://localhost:9200"`
}

var client *elasticsearch.TypedClient
var userUpdated pubsub.Exchange[auth.UserUpdatedEvent]
var log *zap.Logger

func TestMain(m *testing.M) {
	ctx := context.Background()

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(fmt.Errorf("failed to process configs: %s", err))
	}

	client, err = elasticsearch.NewTypedClient(elasticsearch.Config{
		RetryOnStatus: []int{502, 503, 504, 429},
		MaxRetries:    5,
		Addresses:     []string{cfg.Elastic},
	})
	if err != nil {
		panic(err)
	}

	log, err = logger.InitLogger(cfg.Service)
	if err != nil {
		panic(err)
	}

	userUpdated = pubsub.NewExchange[auth.UserUpdatedEvent]()
	defer userUpdated.Shutdown(ctx)

	os.Exit(m.Run())
}

func setup(t *testing.T) *search.Service {
	name := fmt.Sprintf("%s-%s", strings.ToLower(t.Name()), time.Now().Format("01-02--15;04;05"))

	return &search.Service{
		UsersIndex:  "users-" + name,
		PostsIndex:  "posts-" + name,
		TypedClient: client,
		Logger:      log,
	}
}
