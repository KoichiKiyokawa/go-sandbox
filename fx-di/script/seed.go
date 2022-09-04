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
		_, err := db.User.Create().SetName(name).SetEmail(fmt.Sprintf("%s@example.com", name)).Save(ctx)
		if err != nil {
			panic(err)
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
