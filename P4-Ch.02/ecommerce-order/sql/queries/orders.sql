-- name: CreateOrder :one
INSERT INTO orders (id, created_at, modified_at, user_id, product_id) 
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetOrders :many
SELECT id, product_id FROM orders WHERE user_id=$1 LIMIT $2 OFFSET $3;

-- name: GetOrderById :one
SELECT id, created_at, modified_at, product_id FROM orders WHERE id=$1 AND user_id=$2;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE id=$1 AND user_id=$2;