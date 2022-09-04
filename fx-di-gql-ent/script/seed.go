//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"fx-di/ent"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	db := newDB()

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("user%d", i)
		u := db.User.Create().SetName(name).SetEmail(fmt.Sprintf("%s@example.com", name)).SaveX(ctx)
		for j := 0; j < 100; j++ {
			_ = db.Post.Create().SetTitle(fmt.Sprintf("%s-post%d-title", name, j)).SetContent(fmt.Sprintf("%s-post%d-content", name, j)).SetAuthor(u).SaveX(ctx)
		}
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
