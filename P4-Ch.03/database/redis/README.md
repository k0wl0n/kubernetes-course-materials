# Add the helm chart
helm repo add ot-helm https://ot-container-kit.github.io/helm-charts/

# Deploy the redis-operator
helm upgrade redis-operator ot-helm/redis-operator --install --create-namespace --namespace redis-operator

helm test redis-operator --namespace redis-operator

helm upgrade redis-user ot-helm/redis --install --namespace redis
helm upgrade redis-product ot-helm/redis --install --namespace redis
helm upgrade redis-order ot-helm/redis --install --namespace redis