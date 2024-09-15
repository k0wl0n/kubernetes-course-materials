export DB_URL=postgres://db_user:1234@localhost:5432/orders_db?sslmode=disable

goose -dir ./sql/schema/ postgres $DB_URL up 