package main

import (
	"context"
	"fmt"
	"gqlgen-ent/ent"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ctx := context.Background()

	for i := 0; i < 10; i++ {
		client.User.Create().SetName(fmt.Sprintf("user%d", i)).ExecX(ctx)
		client.Post.Create().SetAuthorID()
	}
}
