resource "helm_release" "redis_order" {
  name              = "redis-order"
  namespace         = "redis"
  create_namespace  = true

  repository        = "https://ot-container-kit.github.io/helm-charts/"
  chart             = "redis"
  version           = "0.16.0"
}