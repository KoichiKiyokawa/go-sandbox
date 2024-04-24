-- name: GetUserByID :one
select * from users where id = ?;

-- name: GetUserList :many
select * from users limit @limit offset @offset;

-- name: CreateUser :one
insert into users (name, email) values (@name, @email) returning *;

-- name: ChangeBalance :one
update users set balance = balance + @amount where id = @id returning *;
