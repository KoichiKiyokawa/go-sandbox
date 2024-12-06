package domain

type userID struct {
	value string
}

func (u userID) String() string {
	return u.value
}

type User struct {
	value userValue
}

type userValue struct {
	ID       userID
	Name     string
	Nickname *string
}

func (u User) Value() userValue {
	return u.value
}

func (u *User) AutoGenerateNicknameFromOriginalName() {
	gen := u.value.Name + "さん"
	u.value.Nickname = &gen
}
