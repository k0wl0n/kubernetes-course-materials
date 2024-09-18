## create secret

export VAULT_TOKEN=hvs.JBzCHZmWHSq5dW3cDIjkocey
export VAULT_ADDRESS=localhost:8200  

```vault kv put internal/production/ecommerce-user \
  REDIS_URL="redis-user.redis.svc.cluster.local:6379" \
  DB_URL="postgres://db_super:RXxCqLaPrkHVlm8dx5YwE6v@postgresql-user-service.postgresql.svc.cluster.local:5432/users_db?sslmode=disable" \
  ACCESS_TOKEN_EXP="7" \
  CGO_ENABLED="1" \
  GIN_MODE="debug" \
  INTERNAL_API_SECRET="1234_secret_4321" \
  PORT="8080" \
  REFRESH_TOKEN_EXP="14"
```
