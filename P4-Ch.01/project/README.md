# Overview of the stack

## prequisite

Credentials with admin role
    * AWS IAM access_key 
    * AWS IAM secret_key 

Recommended for windows to install WSL 2
    * https://learn.microsoft.com/en-us/windows/wsl/install
    * can follow this step https://www.youtube.com/watch?v=vxTW22y8zV8

Tools for provisioning and run the project
    package manager
    * windows
        * choco
    * macOs or linux
        * brew

    kubernetes
    * kubectl
    * kubectx
    * kubens
    * k9s (recommended for k8s operational)

    provisioning
    * terraform
    * helm
    * kustomize

    development
    * golang
    * vscode
    * goose
    * sqlc
    * psql or dbeaver ( db client )

production cluster kubernetes high availability
    * setup production eks kubernetes cluster on AWS using terraform
        * Create s3 bucket and dynamoDB state lock
            * P4-Ch.01/eks-cluster-bucket
        * Create cluster kubernetes
            * P4-Ch.01/eks-cluster-production
    * setup istio 
        * P4-Ch.01/istio/README.MD
    * setup argoCD
        * P4-Ch.01/argocd/README.md
    * setup vault

production CICD
    * Create Github Repository
        * Repository for user_service
        * Repository for product_service
        * Repository for order_service
    * Create monorepo-gitops
    * Setup CICD for Service user_service
        * Containerizing user_service
        * Creating Kubernetes Manifests user_service
        * Create ECR for user_service
    * Setup CICD for Service product_service
        * Containerizing product_service
        * Creating Kubernetes Manifests product_service
        * Create ECR for product_service
    * Setup CICD for Service order_service
        * Containerizing order_service
        * Creating Kubernetes Manifests order_service
        * Create ECR for order_service
    * Setup postgresql operator
        * setup postgresql database for user_service
        * setup postgresql database for product_service
        * setup postgresql database for order_service
    * Setup redis operator
        * setup redis for user_service
        * setup redis for product_service
        * setup redis for order_service
    

production monitoring cluster
    * install logging/loki
    * install logging/promtail
    * install metrics/exporter/node_exporter
    * install metrics/exporter/kube-state-metrics
    * install metrics/vmcluster
    * install metrics/vmagent
    * Install visualize/grafana
        * setup grafana datasource loki and vmselect
        * add example dashboard metrics grafana
        * check logging on grafana explore menu
    * setup jaeger tracing/jeager
        * check jaeger UI
    * install alert/vmalert