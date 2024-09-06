resource "helm_release" "stackgres_operator" {
  name              = "stackgres-operator"
  namespace         = "stackgres"
  create_namespace  = true

  repository        = "https://stackgres.io/downloads/stackgres-k8s/stackgres/helm/"
  chart             = "stackgres-operator"
  version           = "1.6.1"

  # Optional: Add custom values file if needed
  values = [file("path/to/your/values.yaml")]
}
