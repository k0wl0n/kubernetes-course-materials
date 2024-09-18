resource "helm_release" "redis_operator" {
  name              = "redis-operator"
  namespace         = "redis-operator"
  create_namespace  = true

  repository        = "https://ot-container-kit.github.io/helm-charts/"
  chart             = "redis-operator" 
  version           = "0.18.0"

}
