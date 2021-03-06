// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: account.sql

package generated

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createAccount = `-- name: CreateAccount :exec
INSERT INTO accounts (username) VALUES ($1::text)
`

func (q *Queries) CreateAccount(ctx context.Context, username string) error {
	_, err := q.db.ExecContext(ctx, createAccount, username)
	return err
}

const findAccountsByIDs = `-- name: FindAccountsByIDs :many
SELECT id, username, creat_at FROM accounts WHERE id = ANY($1::int[])
`

func (q *Queries) FindAccountsByIDs(ctx context.Context, ids []int32) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, findAccountsByIDs, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(&i.ID, &i.Username, &i.CreatAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAccountsWithStatuses = `-- name: FindAccountsWithStatuses :many
select accounts.id, username, creat_at, statuses.id, body, created_at, account_id from accounts join statuses on accounts.id = statuses.account_id
`

type FindAccountsWithStatusesRow struct {
	ID        int32
	Username  sql.NullString
	CreatAt   sql.NullTime
	ID_2      int32
	Body      sql.NullString
	CreatedAt sql.NullTime
	AccountID sql.NullInt32
}

func (q *Queries) FindAccountsWithStatuses(ctx context.Context) ([]FindAccountsWithStatusesRow, error) {
	rows, err := q.db.QueryContext(ctx, findAccountsWithStatuses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindAccountsWithStatusesRow
	for rows.Next() {
		var i FindAccountsWithStatusesRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.CreatAt,
			&i.ID_2,
			&i.Body,
			&i.CreatedAt,
			&i.AccountID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllAccounts = `-- name: FindAllAccounts :many
SELECT id, username, creat_at FROM accounts
`

func (q *Queries) FindAllAccounts(ctx context.Context) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, findAllAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(&i.ID, &i.Username, &i.CreatAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
