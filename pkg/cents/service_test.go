package cents_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/projectulterior/2cents-backend/pkg/cents"
	"github.com/projectulterior/2cents-backend/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Config struct {
	Service string `envconfig:"SERVICE" default:"cents"`
	Mongo   string `envconfig:"MONGO" default:"mongodb://localhost:27017/?replicaSet=rs0&tlsInsecure=true&directConnection=true"`
}

var client *mongo.Client
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

	os.Exit(m.Run())
}

func setup(t *testing.T) *cents.Service {
	name := fmt.Sprintf("%s-%s", t.Name(), time.Now().Format("01-02--15:04:05"))

	return &cents.Service{
		Database: client.Database(name),
		Logger:   log,
	}
}
