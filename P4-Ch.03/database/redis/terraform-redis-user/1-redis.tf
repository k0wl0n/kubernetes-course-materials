resource "helm_release" "redis_user" {
  name              = "redis-user"
  namespace         = "redis"
  create_namespace  = true

  repository        = "https://ot-container-kit.github.io/helm-charts/"
  chart             = "redis"
  version           = "0.16.0" 
}