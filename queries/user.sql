-- name: GetUser :one
SELECT u.id, u.name, u.surname, u.phone, u.role, u.address, u.location_x, u.location_y FROM users u
WHERE u.id = @id;