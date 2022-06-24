//go:build ignore
// +build ignore

package seed

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sqlc-crdb/generated"

	"golang.org/x/sync/errgroup"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	q := generated.New(db)
	const size = 10
	eg := errgroup.Group{}
	for i := 0; i < size; i++ {
		eg.Go(func() error {
			return q.CreateAccount(context.Background(), fmt.Sprintf("user%d", i))
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("done!")
}
