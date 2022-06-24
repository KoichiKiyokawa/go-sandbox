-- name: FindAllAccounts :many
SELECT * FROM accounts;

-- name: FindAccountsByIDs :many
SELECT * FROM accounts WHERE id = ANY(@ids::int[]);

-- name: CreateAccount :exec
INSERT INTO accounts (username) VALUES (@username::text);