-- name: GetUserByID :one
select * from users where id = ?;

-- name: GetUserList :many
select * from users limit @limit offset @offset;