package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID          uint
	Name        string
	Posts       []Post
	CreditCards []CreditCard
}

type Post struct {
	ID     uint
	Title  string
	UserID uint
}

type CreditCard struct {
	ID     uint
	Number string
	UserID uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &Post{}, &CreditCard{})

	// Create
	db.Create([]User{
		{Name: "user1", Posts: []Post{{Title: "post1"}}, CreditCards: []CreditCard{{Number: "123456"}}},
		{Name: "user2", Posts: []Post{{Title: "post2"}}, CreditCards: []CreditCard{{Number: "123456"}}},
	})

	var users []User
	_ = db.Debug().Model(&User{}).Preload("Posts").Find(&users).Error
	log.Printf("users: %#+v\n", users)

	var user User
	_ = db.Debug().Model(&User{}).Preload("Posts").First(&user).Error
	log.Printf("user: %#+v\n", user)
}
