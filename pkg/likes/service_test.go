package likes_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/likes"
	"github.com/projectulterior/2cents-backend/pkg/logger"

	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Config struct {
	Service string `envconfig:"SERVICE" default:"likes"`
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

func setup(t *testing.T) *likes.Service {
	name := fmt.Sprintf("%s-%s", t.Name(), time.Now().Format("01-02--15:04:05"))

	return &likes.Service{
		Database: client.Database(name),
		Logger:   log,
	}
}
