package domain

type userID struct {
	value string
}

func (u userID) String() string {
	return u.value
}

type User struct {
	id       userID
	name     string
	nickname *string
}

func (u User) AutoGenerateNicknameFromOriginalName() User {
	gen := u.name + "さん"
	u.nickname = &gen

	return u
}

// getters
func (u User) GetID() userID {
	return u.id
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetNickname() *string {
	return u.nickname
}
