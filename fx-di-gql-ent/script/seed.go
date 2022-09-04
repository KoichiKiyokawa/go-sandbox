//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"fx-di/ent"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()
	db := newDB(ctx)

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("user%d", i)
		u := db.User.Create().SetName(name).SetEmail(fmt.Sprintf("%s@example.com", name)).SaveX(ctx)
		for j := 0; j < 100; j++ {
			_ = db.Post.Create().SetTitle(fmt.Sprintf("%s-post%d-title", name, j)).SetContent(fmt.Sprintf("%s-post%d-content", name, j)).SetAuthor(u).SaveX(ctx)
		}
	}
}

func newDB(ctx context.Context) *ent.Client {
	client, err := ent.Open("sqlite3", "file:dev.db?_fk=1")
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err != nil {
		panic(err)
	}

	return client
}
