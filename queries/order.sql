-- name: CreateOrder :exec
INSERT INTO orders(user_id, baskets_id, address, address_location_x, address_location_y, pick_up_location_x, pick_up_location_y, metadata, total_price)
VALUES (@user_id, @baskets_id, @address, @address_location_x, @address_location_y, @pick_up_location_x, @pick_up_location_y, @metadata, @total_price);

-- name: CancelOrder :exec
DELETE FROM orders
WHERE id = @id;


-- name: GetOrderOwner :one
SELECT user_id
FROM orders
WHERE id = @id;