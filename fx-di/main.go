package main

import (
	"fx-di/ent"
	"fx-di/generated"
	"fx-di/infra/dao"
	"fx-di/resolver"
	"fx-di/service"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	app := fx.New(
		fx.Provide(
			func() *zap.Logger { return logger },
			newDB,
			dao.NewUserRepository,
			dao.NewPostRepository,
			service.NewUserService,
			service.NewPostService,
			resolver.NewResolver,
		),
		fx.Invoke(register),
		fx.WithLogger(
			func() fxevent.Logger {
				return &fxevent.ZapLogger{Logger: logger}
			},
		),
	)

	app.Run()
}

func newDB() *ent.Client {
	client, err := ent.Open("sqlite3", "file:dev.db?_fk=1")
	if err != nil {
		panic(err)
	}

	if os.Getenv("DB_DEBUG") != "" {
		return client.Debug()
	}
	return client
}

const defaultPort = "8080"

func register(resolver *resolver.Resolver) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
