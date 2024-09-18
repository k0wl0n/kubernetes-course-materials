# Overview of the Stack

## Prerequisites

- **Credentials with admin role**:
  - AWS IAM `access_key`
  - AWS IAM `secret_key`
  - aws cli

- **Recommended for Windows**:
  - Install [WSL 2](https://learn.microsoft.com/en-us/windows/wsl/install) (You can follow [this guide](https://www.youtube.com/watch?v=vxTW22y8zV8))

## Tools for Provisioning and Project Setup

### Package Managers
- **Windows**: [Chocolatey](https://chocolatey.org/)
- **macOS/Linux**: [Homebrew](https://brew.sh/)

### Kubernetes
- `kubectl`
- `kubectx`
- `kubens`
- [`k9s`](https://k9scli.io/) (recommended for k8s operations)

### Provisioning
- [Terraform](https://www.terraform.io/)
- [Helm](https://helm.sh/)
- [Kustomize](https://kustomize.io/)

### Development
- [Golang](https://go.dev/)
- [Visual Studio Code](https://code.visualstudio.com/)
- [Goose](https://github.com/pressly/goose) (DB migration tool)
- [SQLC](https://sqlc.dev/) (Go code generation from SQL queries)
- `psql` or [DBeaver](https://dbeaver.io/) (DB client)

## Production Kubernetes Cluster (High Availability)

1. **Setup Production EKS Kubernetes Cluster on AWS using Terraform**:
    - **Create S3 bucket and DynamoDB state lock**:
        - Path: `P4-Ch.01/eks-cluster-bucket`
    - **Create EKS cluster**:
        - Path: `P4-Ch.01/eks-cluster-production`

2. **Setup Istio**:
    - Path: `P4-Ch.01/istio/README.MD`

3. **Setup ArgoCD**:
    - Path: `P4-Ch.01/argocd/README.md`

4. **Setup Vault**:
    - Configure as required

## Production CI/CD

1. **Create GitHub Repositories**:
    - Repository for `user_service`
    - Repository for `product_service`
    - Repository for `order_service`
    - Monorepo for GitOps

2. **Setup CI/CD for Services**:
    - **User Service**:
        - Containerize `user_service`
        - Create Kubernetes manifests
        - Create ECR for `user_service`
    - **Product Service**:
        - Containerize `product_service`
        - Create Kubernetes manifests
        - Create ECR for `product_service`
    - **Order Service**:
        - Containerize `order_service`
        - Create Kubernetes manifests
        - Create ECR for `order_service`

3. **Setup PostgreSQL Operator**:
    - Setup PostgreSQL databases for:
        - `user_service`
        - `product_service`
        - `order_service`

4. **Setup Redis Operator**:
    - Setup Redis for:
        - `user_service`
        - `product_service`
        - `order_service`

## Production Monitoring Cluster

1. **Logging**:
    - Install Loki
    - Install Promtail

2. **Metrics**:
    - Install `node_exporter`
    - Install `kube-state-metrics`
    - Install VMCluster (VictoriaMetrics)
    - Install VMAgent

3. **Visualization**:
    - Install Grafana
        - Setup Grafana datasources for Loki and VMSelect
        - Add example Grafana dashboards
        - Check logs via Grafana Explore

4. **Tracing**:
    - Setup Jaeger
        - Check Jaeger UI

5. **Alerting**:
    - Install VMAlert