-- name: GetProduct :one
SELECT * FROM Products pr
WHERE pr.id = @id;

-- name: GetProducts :many
SELECT * FROM Products;

-- name: UpdateProduct :exec
UPDATE Products
SET quantity = @quantity
WHERE id = @id;

-- name: GetProductPrice :one
SELECT pr.price FROM Products pr
WHERE pr.id = @id;