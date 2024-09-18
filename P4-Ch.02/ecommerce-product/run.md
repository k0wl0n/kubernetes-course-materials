export DB_URL=postgres://db_super:RXxCqLaPrkHVlm8dx5YwE6v@localhost:5432/product_db?sslmode=disable
goose -dir ./sql/schema/ postgres $DB_URL up 


goose -dir ./sql/schema/ postgres "postgres://db_super:RXxCqLaPrkHVlm8dx5YwE6v@localhost:5432/product_db?sslmode=disable" up 