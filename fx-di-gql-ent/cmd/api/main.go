package main

import (
	"fmt"
	"fx-di/ent"
	"fx-di/generated"
	"fx-di/infra/dao"
	"fx-di/resolver"
	"fx-di/service"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {

	app := fx.New(
		fx.Provide(
			// main
			newLogger,
			newDB,

			// dao
			dao.NewUserRepository,
			dao.NewPostRepository,

			// service
			service.NewUserService,
			service.NewPostService,

			// resolver
			resolver.NewResolver,
		),
		fx.Invoke(register),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		},
		),
	)

	app.Run()
}

func newLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return logger
}

func newDB() *ent.Client {
	client, err := ent.Open("postgres", os.Getenv("DB_URL"))

	if err != nil {
		panic(err)
	}

	if os.Getenv("DB_DEBUG") != "" {
		return client.Debug()
	}
	return client
}

const defaultPort = "8080"

func register(resolver *resolver.Resolver, logger *zap.Logger, client *ent.Client) {
	defer client.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Info(fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", port))
	logger.Error("serve error", zap.Error(http.ListenAndServe("127.0.0.1:"+port, nil)))
}
