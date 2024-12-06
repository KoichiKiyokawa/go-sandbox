package domain

import (
	"huma-sandbox/.gen/postgres/public/model"
)

func FromUserModelToDomain(userModel model.Users) User {
	return User{
		value: userValue{
			ID:       userID{value: userModel.ID},
			Name:     userModel.Name,
			Nickname: userModel.Nickname,
		},
	}
}

func FromUserDomainToModel(user User) model.Users {
	return model.Users{
		ID:       user.value.ID.String(),
		Name:     user.value.Name,
		Nickname: user.value.Nickname,
	}
}
