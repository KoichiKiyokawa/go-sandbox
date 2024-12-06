package domain

type userID struct {
	value string
}

func (u userID) String() string {
	return u.value
}

type User struct {
	ID       readonly[userID]
	Name     readonly[string]
	Nickname readonly[*string]
}

func (u *User) AutoGenerateNicknameFromOriginalName() {
	gen := u.Name.Value() + "さん"
	u.Nickname = toReadonly(&gen)
}
