-- +goose Up

CREATE TABLE products (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    creator_id UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    price INTEGER NOT NULL,
    description TEXT NOT NULL,
    CONSTRAINT UniqueProduct UNIQUE (creator_id, name)
);

-- +goose Down
DROP TABLE products;