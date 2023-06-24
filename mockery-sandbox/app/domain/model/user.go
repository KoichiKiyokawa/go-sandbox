package model

type User struct {
	id   int
	name string
}

func (u *User) ID() int {
	return u.id
}

func (u *User) Name() string {
	return u.name
}
