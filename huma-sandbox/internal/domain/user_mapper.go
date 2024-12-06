package domain

import (
	"huma-sandbox/.gen/postgres/public/model"
)

func FromUserModelToDomain(userModel model.Users) User {
	return User{
		Name:     toReadonly(userModel.Name),
		ID:       toReadonly(userID{value: userModel.ID}),
		Nickname: toReadonly(userModel.Nickname),
	}
}
