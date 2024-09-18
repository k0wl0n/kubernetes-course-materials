export DB_URL=postgres://db_super:RXxCqLaPrkHVlm8dx5YwE6v@localhost:5432/order_db?sslmode=disable

goose -dir ./sql/schema/ postgres "postgres://db_super:RXxCqLaPrkHVlm8dx5YwE6v@localhost:5432/order_db?sslmode=disable" up 

postgres://db_super:RXxCqLaPrkHVlm8dx5YwE6v@postgresql-order-service.postgresql.svc.cluster.local:5432/order_db?sslmode=disable