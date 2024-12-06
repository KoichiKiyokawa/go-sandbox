package main

import (
	"context"
	"database/sql"
	"fmt"
	"huma-sandbox/internal/logger"
	"huma-sandbox/internal/schema"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/humacli"
)

type Options struct {
	Port int `default:"8888" help:"Port to listen on" short:"p"`
}

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		mux := http.NewServeMux()

		db := &sql.DB{}

		api := humago.New(mux, huma.DefaultConfig("My API", "1.0.0"))
		api.UseMiddleware(logger.LogMiddleware)
		schema.RegisterUserHandlers(api, db)

		//nolint:exhaustruct
		server := http.Server{
			Addr:              fmt.Sprintf("127.0.0.1:%d", options.Port),
			Handler:           mux,
			ReadHeaderTimeout: 5 * time.Second,
		}

		// Tell the CLI how to start your router.
		hooks.OnStart(func() {
			must(server.ListenAndServe())
		})

		// Tell the CLI how to stop your server.
		hooks.OnStop(func() {
			// Give the server 5 seconds to gracefully shut down, then give up.
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			must(server.Shutdown(ctx))
		})
	})
	cli.Run()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
