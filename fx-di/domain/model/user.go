package model

type User struct {
	ID   int `bun:"id,pk"`
	Name string
}
