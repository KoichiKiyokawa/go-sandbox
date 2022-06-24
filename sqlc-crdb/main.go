//go:generate go run github.com/kyleconroy/sqlc/cmd/sqlc generate
package main

import (
	"context"
	"database/sql"
	"log"

	"sqlc-crdb/generated"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://root@4000b43cf5ba:26257?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	queries := generated.New(db)
	res, _ := queries.FindAllAccounts(context.Background())
	log.Printf("res: %#+v\n", res)
}
