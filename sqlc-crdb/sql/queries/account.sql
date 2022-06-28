-- name: FindAllAccounts :many
SELECT * FROM accounts;

-- name: FindAccountsByIDs :many
SELECT * FROM accounts WHERE id = ANY(@ids::int[]);

-- name: FindAccountsWithStatuses :many
select * from accounts join statuses on accounts.id = statuses.account_id;

-- name: CreateAccount :exec
INSERT INTO accounts (username) VALUES (@username::text);