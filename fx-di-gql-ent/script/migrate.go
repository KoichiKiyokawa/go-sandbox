//go:build ignore
// +build ignore

package main

import (
	"context"
	"fx-di/ent"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	client := newDB()
	if os.Getenv("RESET") != "" {
		reset(ctx, client)
	}

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
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

func reset(ctx context.Context, client *ent.Client) {
	client.Post.Delete().ExecX(ctx)
	client.User.Delete().ExecX(ctx)
}
