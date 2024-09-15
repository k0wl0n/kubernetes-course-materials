-- name: CreateProduct :one
INSERT INTO products (id, created_at, modified_at, creator_id, name, price, description) 
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetProducts :many
SELECT id, name, price, description FROM products LIMIT $1 OFFSET $2;

-- name: GetProductById :one
SELECT * FROM products WHERE id=$1 FOR UPDATE NOWAIT;

-- name: UpdateProductDetails :one
UPDATE products SET modified_at=$1, name=$2, price=$3, description=$4
WHERE id=$5
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id=$1;