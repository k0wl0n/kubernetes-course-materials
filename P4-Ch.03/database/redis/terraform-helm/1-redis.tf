resource "helm_release" "redis_operator" {
  name              = "redis-operator"
  namespace         = "redis-operator"
  create_namespace  = true

  repository        = "https://ot-container-kit.github.io/helm-charts/"
  chart             = "redis-operator"  # Corrected chart name
  version           = "0.18.0"

}
