-- +goose Up

CREATE TABLE orders (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL,
    product_id UUID NOT NULL
);

-- +goose Down
DROP TABLE orders;