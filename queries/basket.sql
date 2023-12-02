-- name: AddBasketItem :one
INSERT INTO basket(product_id, quantity)
VALUES (@product_id, @quantity)
RETURNING id;

-- name: DeleteBasketItem :exec
DELETE FROM basket
WHERE id = @id;