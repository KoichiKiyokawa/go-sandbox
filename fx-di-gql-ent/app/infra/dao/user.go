package dao

import (
	"context"
	"fx-di/app/domain/repository"
	"fx-di/ent"
	"fx-di/ent/user"
)

type userRepository struct {
	db *ent.Client
}

func NewUserRepository(db *ent.Client) repository.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindOne(ctx context.Context, id int) (*ent.User, error) {
	return r.db.User.Query().Where(user.IDEQ(id)).First(ctx)
}

func (r *userRepository) FindAll(ctx context.Context) ([]*ent.User, error) {
	return r.db.User.Query().All(ctx)
}

func convertUser(u *ent.User) *ent.User {
	return &ent.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
