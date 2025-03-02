package storage

import (
	"context"

	"huma-sandbox/.gen/postgres/public/model"
	"huma-sandbox/.gen/postgres/public/table"
	"huma-sandbox/internal/domain"

	"braces.dev/errtrace"
	qb "github.com/go-jet/jet/v2/postgres"
)

type FindUserListFilter struct {
	Name *string
}

func (s *Storage) FindUserList(ctx context.Context, filter FindUserListFilter) ([]domain.User, error) {
	query := qb.SELECT(table.Users.AllColumns).FROM(table.Users)

	if filter.Name != nil {
		query = query.WHERE(table.Users.Name.EQ(qb.String(*filter.Name)))
	}

	var dest []model.Users
	if err := query.QueryContext(ctx, s.db, &dest); err != nil {
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
	input := domain.FromUserDomainToModel(user)

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
	input := domain.FromUserDomainToModel(user)

	var dest model.Users
	if err := table.Users.
		UPDATE(table.Users.EXCLUDED.CreatedAt, table.Users.EXCLUDED.UpdatedAt).
		SET(input).
		WHERE(table.Users.ID.EQ(qb.String(input.ID))).
		RETURNING(table.Users.AllColumns).
		QueryContext(ctx, s.db, &dest); err != nil {
		return nil, errtrace.Wrap(err)
	}

	user = domain.FromUserModelToDomain(dest)

	return &user, nil
}
