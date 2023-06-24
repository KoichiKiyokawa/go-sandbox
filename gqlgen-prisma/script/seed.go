package main

import (
	"context"
	"fmt"
	"gqlgen-prisma/db"
)

func main() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	for i := 1; i <= 10; i++ {
		user, err := client.User.CreateOne(
			db.User.Name.Set(fmt.Sprintf("user%d", i)),
			db.User.Email.Set(fmt.Sprintf("user%d@example.com", i)),
		).Exec(ctx)
		if err != nil {
			panic(err)
		}

		for j := 1; j <= 100; j++ {
			_, err := client.Post.CreateOne(
				db.Post.Title.Set(fmt.Sprintf("post%d", j)),
				db.Post.Content.Set(fmt.Sprintf("content%d", j)),
				db.Post.Author.Link(db.User.ID.Equals(user.ID)),
			).Exec(ctx)
			if err != nil {
				panic(err)
			}
		}
	}
}
