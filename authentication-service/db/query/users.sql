-- name: GetAllUsers :many
select id, email, first_name, last_name, password, user_active, created_at, updated_at from users order by last_name;

-- name: GetByEmailUsers :one
select id, email, first_name, last_name, password, user_active, created_at, updated_at from users where email = ($1);

-- name: GetOneUsers :one
select id, email, first_name, last_name, password, user_active, created_at, updated_at from users where index = ($1);

-- name: UpdateUserByID :exec
update users set email = ($1), first_name = ($2), last_name = ($3), user_active = ($4), updated_at = ($5) where index = ($6);

-- name: DeleteUserByID :exec
delete from users where index = ($1);

-- name: CreateUserID :one
insert into users (email, first_name, last_name, password, user_active, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7) returning index;

-- name: ResetPassword :exec
update users set password = ($1) where index = ($2);