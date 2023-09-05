package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/projectulterior/2cents-backend/cmd/daemon/handler"
	"github.com/projectulterior/2cents-backend/cmd/daemon/httputil"
	"github.com/projectulterior/2cents-backend/cmd/daemon/middleware"
	"github.com/projectulterior/2cents-backend/graph"
	"github.com/projectulterior/2cents-backend/pkg/auth"
	"github.com/projectulterior/2cents-backend/pkg/logger"
	"github.com/projectulterior/2cents-backend/pkg/os/process"
	http_server "github.com/projectulterior/2cents-backend/pkg/server/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/kelseyhightower/envconfig"
	"go.elastic.co/apm/module/apmhttp/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Config struct {
	Host            string        `envconfig:"HOST" default:"http://localhost:8080"`
	Port            int           `envconfig:"PORT" default:"8080"`
	Secret          string        `envconfig:"SECRET" default:"secret"`
	AuthTokenTTL    time.Duration `envconfig:"AUTH_TOKEN_TTL" default:"10m"`
	RefreshTokenTTL time.Duration `envconfig:"REFRESH_TOKEN_TTL" default:"1h"`
	Mongo           string        `evnconfig:"MONGO"`
}

func main() {
	ctx := process.Context()

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(fmt.Errorf("failed to process configs: %s", err))
	}

	// logs
	log, err := logger.InitLogger("")
	if err != nil {
		panic(err)
	}

	// mongo
	m, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Mongo))
	if err != nil {
		panic(err)
	}
	defer m.Disconnect(ctx)

	// services
	svc, err := initServices(ctx, cfg, m, log)
	if err != nil {
		panic(err)
	}

	c := graph.Config{
		Resolvers: &graph.Resolver{
			Services: svc,
		},
	}

	ready := httputil.NewReady(
		httputil.TextHandler(http.StatusServiceUnavailable, "application/json", `"NOT READY"`),
	)

	mux := httputil.NewBaseMux(
		ready.Handler(httputil.TextHandler(http.StatusOK, "application/json", `"READY"`)),
	)

	mux.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	mux.HandleFunc("/explorer", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r,
			fmt.Sprintf("https://sandbox.apollo.dev/?endpoint=%s", cfg.Host+"/graphql"),
			http.StatusSeeOther,
		)
	})

	srv := graphql_handler.New(graph.NewExecutableSchema(c))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			if initPayload.Authorization() == "" {
				// NOTE: on initialization, the first call to
				// InitFunc does not contain anything in `initPayload`
				return ctx, nil
			}

			reply, err := svc.Auth.VerifyToken(ctx, auth.VerifyTokenRequest{
				Token: initPayload.Authorization(),
			})
			st, ok := status.FromError(err)
			if !ok {
				log.Info("error in verifying")
				return nil, fmt.Errorf("error in decoding error")
			}

			switch st.Code() {
			case codes.OK:
			case codes.PermissionDenied:
				log.Info("permission denied here")
				return nil, fmt.Errorf("permission denied: %s", st.Message())
			default:
				log.Info("unknown err", zap.String("message", st.Message()))
				return nil, fmt.Errorf("unkwown error: %s", st.Message())
			}

			return context.WithValue(ctx, middleware.AUTH_USER_CONTEXT_KEY, reply.UserID), nil
		},
	})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	mux.Handle("/graphql",
		middleware.NewAuth(svc.Auth,
			srv,
		),
	)

	mux.HandleFunc("/auth/create_token", handler.HandleCreateToken(svc.Auth))
	mux.HandleFunc("/auth/refresh_token", handler.HandleRefreshToken(svc.Auth))

	handler := middleware.NewRecover(
		middleware.NewLogger(
			middleware.NewCors(
				mux,
			),
		),
	)

	server := http_server.Server{
		Port:    cfg.Port,
		Handler: apmhttp.Wrap(handler),
	}
	server.Run()
	defer server.Shutdown()

	ready.Ready()
	log.Info("server started", zap.String("explorer", fmt.Sprintf("%s/explorer", cfg.Host)))

	<-ctx.Done()

	ready.Unready()
	log.Info("shutting down server")
}
