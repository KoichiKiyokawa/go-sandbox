package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"go-workspace-module-account/command/infra"
	uiapi "go-workspace-module-account/command/ui/api"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/samber/do"
)

type Options struct {
	Debug bool   `doc:"Enable debug logging"`
	Host  string `doc:"Hostname to listen on."`
	Port  int    `doc:"Port to listen on." short:"p" default:"8888"`
}

func main() {
	// Create a CLI app which takes a port option.
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Create a new router & API
		router := chi.NewMux()
		api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))
		db, err := sql.Open("sqlite3", "")
		if err != nil {
			panic(err)
		}

		// di
		injector := do.New()
		do.ProvideValue(injector, db)
		do.ProvideValue(injector, api)
		do.Provide(injector, infra.NewDao)
		do.Provide(injector, uiapi.NewHandler)
		handler := do.MustInvoke[*uiapi.Handler](injector)
		handler.RegisterAll()

		// Tell the CLI how to start your router.
		hooks.OnStart(func() {
			fmt.Printf("server hosted in http://localhost:8888")
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
		})
	})

	// Run the CLI. When passed no commands, it starts the server.
	cli.Run()
}
