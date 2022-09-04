//go:build ignore
// +build ignore

package main

import (
	"context"
	"fx-di/ent"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:dev.db?_fk=1")
	if err != nil {
		log.Fatalf("failed open database: %v", err)

	}
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
