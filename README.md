# Kubernetes Manifest Repository

A collection of Kubernetes manifests arranged by chapter and topic. The materials were
originally created for the Fast Campus course **Docker & Kubernetes from Containerization to Real-World Applications (ft. AWS EKS)**,
but the examples are free to use for anyone exploring Kubernetes.

Fast Campus is Korea's leading edutech company with over 930,000 members and
high-quality online lectures for working professionals in Indonesia and beyond.

## Contents

Each chapter directory contains YAML manifests that can be applied directly with `kubectl` or
adapted for your own experiments.

### P2 – Kubernetes Fundamentals

- **Ch.02** – Introduction to Kubernetes basics.
- **Ch.03** – Building Docker images with the `go-echo` example application (`go-echo`, `go-echo-v1`, `go-echo-v2`, `go-echo-v3`).
- **Ch.04** – Core workloads: Pods, ReplicaSets, Deployments, DaemonSets, Jobs, CronJobs,
  StatefulSets, rollouts/rollbacks (`go-echo-v4`), ConfigMaps, Secrets, HPAs, resource
  limits and scheduling.
- **Ch.05** – Networking: pod connectivity, ClusterIP, NodePort, LoadBalancer and Ingress controllers.
- **Ch.06** – Storage: storage classes and persistent volumes.

### P3 – CI/CD & Observability

- **Ch.01** – Terraform configurations for infrastructure setup.
- **Ch.02** – Argo CD, ECR, CI/CD pipelines, monitoring stack and SSH key management.

### P4 – Production Cluster Examples

- **Ch.01–05** – End‑to‑end EKS cluster setup, Istio service mesh, Vault integration and other production tools.

## Getting Started

1. **Clone the repository**
   ```bash
   git clone https://github.com/k0wl0n/kubernetes-course-materials.git
   cd kubernetes-course-materials
   ```

2. **Navigate to an example**
   ```bash
   cd P2-Ch.04/P2-Ch.04-2-5.Pods
   ```

3. **Apply a manifest**
   ```bash
   kubectl apply -f <manifest>.yaml
   ```

Feel free to modify the manifests and use them as a starting point for your own Kubernetes experiments.

## License

This project is licensed under the MIT License. See the [`LICENSE`](LICENSE) file for details.

## Contact

Questions or suggestions? Open an issue in this repository.
