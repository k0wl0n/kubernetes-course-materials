resource "helm_release" "hashicorp_vault" {
  name              = "hashicorp-vault"
  namespace         = "vault"
  create_namespace  = true

  repository        = "https://helm.releases.hashicorp.com"
  chart             = "vault"
  version           = "0.27.0" 

  # Optional: Add custom values file if needed
  values = [file("values/values.yaml")]
}
