package dao

import (
	"context"
	"fx-di/app/domain/repository"
	"fx-di/ent"
	"fx-di/ent/post"
	"fx-di/ent/user"
)

type userRepository struct {
	db *ent.Client
}

func NewUserRepository(db *ent.Client) repository.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindOne(ctx context.Context, id int) (*ent.User, error) {
	u, err := r.db.User.Query().Where(user.IDEQ(id)).First(ctx)
	if err != nil {
		return nil, err
	}

	return convertUser(u), nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]*ent.User, error) {
	users, err := r.db.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*ent.User, len(users))
	for i, u := range users {
		result[i] = convertUser(u)
	}
	return result, nil
}

func (r *userRepository) FindOneByPostID(ctx context.Context, postID int) (*ent.User, error) {
	u, err := r.db.User.Query().Where(user.HasPostsWith(post.IDEQ(postID))).First(ctx)
	if err != nil {
		return nil, err
	}

	return convertUser(u), nil
}

func convertUser(u *ent.User) *ent.User {
	return &ent.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
