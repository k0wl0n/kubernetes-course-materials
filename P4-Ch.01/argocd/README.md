
# apply  argocd using terraform
terraform init
terraform plan
terraform apply

# port forward

 kubectl port-forward svc/argocd-server -n argocd 8080:443