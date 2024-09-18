
terraform init
terraform plan
terraform apply


kubectl exec -ti  hashicorp-vault-0  -n vault -- sh                   
vault operator init

save unseal key, example

```
Unseal Key 1: 8fS0ZzjQZyQ5fDnDzbR8Uf6BeyQjxY/U5v2MItKuavkT
Unseal Key 2: h3CLUn6gqDWhANZdFMTSPeWC15H7jCcfgZmvmCh44MVM
Unseal Key 3: 7SXHJpV7MAgTlv/7lPAoTM4jgCdCVFqcYAh4St4LwtRb
Unseal Key 4: uL8oo1djcXBhIUYF0tVTqH5vUr8mmKbRODpJX40usP6s
Unseal Key 5: S8RcyP0JwSrAe8VOjyjoWKM6DztS44C6lCQHy9cuUD76

Initial Root Token: hvs.JBzCHZmWHSq5dW3cDIjkocey

Vault initialized with 5 key shares and a key threshold of 3. Please securely
distribute the key shares printed above. When the Vault is re-sealed,
restarted, or stopped, you must supply at least 3 of these keys to unseal it
before it can start servicing requests.

Vault does not store the generated root key. Without at least 3 keys to
reconstruct the root key, Vault will remain permanently sealed!

It is possible to generate new unseal keys, provided you have a quorum of
existing unseal keys shares. See "vault operator rekey" for more information.
```

## unseal the vault
kubectl port-forward -n vault svc/hashicorp-vault  8200:8200

open http://localhost:8200
insert 3 unseal key

## setup root token

export VAULT_TOKEN=hvs.JBzCHZmWHSq5dW3cDIjkocey
export VAULT_ADDRESS=http://localhost:8200  


## create secret internal

vault secrets enable -path=internal kv-v2

## create secret

```vault kv put internal/production/ecommerce-user \
  ACCESS_TOKEN_EXP="7" \
  CGO_ENABLED="1" \
  DB_URL="postgres://db_super:asdas@<db-address>" \
  GIN_MODE="debug" \
  INTERNAL_API_SECRET="1234_secret_4321" \
  PORT="8080" \
  REDIS_URL="redis-user.redis.svc.cluster.local:6379" \
  REFRESH_TOKEN_EXP="14"
```

### example

```
<<K9s-Shell>> Pod: vault/hashicorp-vault-0 | Container: vault

/ $ export VAULT_TOKEN
/ $ vault secrets enable -path=internal kv-v2
Success! Enabled the kv-v2 secrets engine at: internal/
/ $ vault kv put internal/production/ecommerce-user \
>   ACCESS_TOKEN_EXP="7" \
>   CGO_ENABLED="1" \
>   DB_URL="postgres://db_super:asdas@<db-address>" \
>   GIN_MODE="debug" \
>   INTERNAL_API_SECRET="1234_secret_4321" \
>   PORT="8080" \
>   REDIS_URL="redis-user.redis.svc.cluster.local:6379" \
>   REFRESH_TOKEN_EXP="14"
======== Secret Path ========
internal/data/database/config

======= Metadata =======
Key                Value
---                -----
created_time       2024-09-14T19:26:23.491462939Z
custom_metadata    <nil>
deletion_time      n/a
destroyed          false
version            1
```

## Enable kubernetes authentication

vault auth enable kubernetes

vault write auth/kubernetes/config token_reviewer_jwt="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" kubernetes_host="https://$KUBERNETES_PORT_443_TCP_ADDR:443" kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt

Success! Data written to: auth/kubernetes/config


## create permission

vault policy write internal-app - <<EOF
path "internal/data/production*" {
   capabilities = ["read"]
}
EOF

internal/data/production/ecommerce-user
internal/data/production/ecommerce-product
internal/data/production/ecommerce-order

## assign permision

vault write auth/kubernetes/role/internal-app \
      bound_service_account_names=internal-app \
      bound_service_account_namespaces=production \
      policies=internal-app \
      ttl=24h

## create service account on production namespace
kubectl create ns production
kubectl create sa internal-app -n production

### test secret
kubectl apply -f P4-Ch.02/vault/env-check.yaml