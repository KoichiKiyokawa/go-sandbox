package main

import (
	"context"
	"encoding/json"
	"fmt"
	repository "without-di/db"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println(run())
}

func run() string {
	ctx := context.Background()
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	userRepo := repository.NewUserRepository(db)
	user, _ := userRepo.FindById(ctx, 1)
	res, _ := json.Marshal(user)
	return string(res)
}
