package storage

import (
	"context"
	"huma-sandbox/internal/domain"

	".gen/postgres/public/model"
	".gen/postgres/public/table"
	"braces.dev/errtrace"
	qb "github.com/go-jet/jet/v2/postgres"
	"github.com/oklog/ulid/v2"
)

func (s *Storage) FindUserList(ctx context.Context) ([]domain.User, error) {
	var dest []model.Users
	if err := qb.SELECT(table.Users.AllColumns).FROM(table.Users).QueryContext(ctx, s.db, &dest); err != nil {
		return nil, errtrace.Wrap(err)
	}

	users := make([]domain.User, 0, len(dest))
	for _, d := range dest {
		users = append(users, domain.FromUserModelToDomain(d))
	}

	return users, nil
}

func (s *Storage) FindUser(ctx context.Context, id string) (*domain.User, error) {
	var dest model.Users
	if err := qb.SELECT(table.Users.AllColumns).
		FROM(table.Users).
		WHERE(table.Users.ID.EQ(qb.String(id))).
		QueryContext(ctx, s.db, &dest); err != nil {
		return nil, errtrace.Wrap(err)
	}

	user := domain.FromUserModelToDomain(dest)

	return &user, nil
}

func (s *Storage) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	input := model.Users{
		ID:       ulid.Make().String(),
		Name:     user.Name.Value(),
		Nickname: user.Nickname.Value(),
	}

	var dest model.Users
	if err := table.Users.
		INSERT(table.Users.AllColumns).
		MODEL(input).
		RETURNING(table.Users.AllColumns).
		QueryContext(ctx, s.db, &dest); err != nil {
		return nil, errtrace.Wrap(err)
	}

	user = domain.FromUserModelToDomain(dest)

	return &user, nil
}

func (s *Storage) UpdateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	input := model.Users{
		ID:       user.ID.Value().String(),
		Name:     user.Name.Value(),
		Nickname: user.Nickname.Value(),
	}

	var dest model.Users
	if err := table.Users.
		UPDATE(table.Users.EXCLUDED.CreatedAt, table.Users.EXCLUDED.UpdatedAt).
		SET(input).
		WHERE(table.Users.ID.EQ(qb.String(user.ID.Value().String()))).
		RETURNING(table.Users.AllColumns).
		QueryContext(ctx, s.db, &dest); err != nil {
		return nil, errtrace.Wrap(err)
	}

	user = domain.FromUserModelToDomain(dest)

	return &user, nil
}
