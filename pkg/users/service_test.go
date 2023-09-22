package users_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/logger"
	"github.com/projectulterior/2cents-backend/pkg/pubsub"
	"github.com/projectulterior/2cents-backend/pkg/users"

	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Config struct {
	Service string `envconfig:"SERVICE" default:"users"`
	Mongo   string `envconfig:"MONGO" default:"mongodb://localhost:27017/?replicaSet=rs0&tlsInsecure=true&directConnection=true"`
}

var client *mongo.Client
var userUpdated pubsub.Exchange[users.UserUpdatedEvent]
var userDeleted pubsub.Exchange[users.UserDeletedEvent]
var log *zap.Logger

func TestMain(m *testing.M) {
	ctx := context.Background()

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(fmt.Errorf("failed to process configs: %s", err))
	}

	fmt.Println("mongo:", cfg.Mongo)

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(cfg.Mongo))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	log, err = logger.InitLogger(cfg.Service)
	if err != nil {
		panic(err)
	}

	userUpdated = pubsub.NewExchange[users.UserUpdatedEvent]()
	defer userUpdated.Shutdown(ctx)

	userDeleted = pubsub.NewExchange[users.UserDeletedEvent]()
	defer userDeleted.Shutdown(ctx)

	os.Exit(m.Run())
}

func setup(t *testing.T) *users.Service {
	name := fmt.Sprintf("%s-%s", t.Name(), time.Now().Format("01-02--15:04:05"))

	return &users.Service{
		UserUpdated: userUpdated.Publisher(),
		UserDeleted: userDeleted.Publisher(),

		Database: client.Database(name),
		Logger:   log,
	}
}
