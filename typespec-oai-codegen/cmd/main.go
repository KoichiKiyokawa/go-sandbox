package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"typespec-oai-codegen/generated"
	"typespec-oai-codegen/generated/db"
	"typespec-oai-codegen/handler"

	_ "github.com/mattn/go-sqlite3"

	echomiddleware "github.com/labstack/echo/v4/middleware"
	middleware "github.com/oapi-codegen/echo-middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	swagger, err := generated.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	e := echo.New()
	e.Use(echomiddleware.Logger())
	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	e.Use(middleware.OapiRequestValidator(swagger))
	e.File("/swagger.yaml", "tsp-output/@typespec/openapi3/openapi.yaml")

	sqlite3, err := sql.Open("sqlite3", "dev.db")
	if err != nil {
		panic(err)
	}

	queries := db.New(sqlite3)
	h := handler.NewHandler(*queries)

	si := generated.NewStrictHandler(h, nil)
	// We now register our petStore above as the handler for the interface
	generated.RegisterHandlers(e, si)

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(net.JoinHostPort("127.0.0.1", getPort())))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
