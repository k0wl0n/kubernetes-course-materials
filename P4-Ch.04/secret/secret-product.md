## create secret

```vault kv put internal/production/ecommerce-product \
  ACCESS_TOKEN_EXP="7" \
  CGO_ENABLED="1" \
  DB_URL="postgres://db_super:RXxCqLaPrkHVlm8dx5YwE6v@postgresql-product-service.postgresql.svc.cluster.local:5432/product_db?sslmode=disable" \
  GIN_MODE="debug" \
  INTERNAL_API_SECRET="1234_secret_4321" \
  PORT="8080" \
  REDIS_URL="redis-product.redis.svc.cluster.local:6379" \
  USER_SERVICE_HOST=ecommerce-user.production.svc.cluster.local:8001
  REFRESH_TOKEN_EXP="14"
```
