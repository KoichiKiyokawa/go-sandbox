package domain

import (
	"time"

	"huma-sandbox/.gen/postgres/public/model"
)

func FromUserModelToDomain(userModel model.Users) User {
	return User{
		id:       userID{value: userModel.ID},
		name:     userModel.Name,
		nickname: userModel.Nickname,
	}
}

func FromUserDomainToModel(user User) model.Users {
	return model.Users{
		ID:        user.id.String(),
		Name:      user.name,
		Nickname:  user.nickname,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
