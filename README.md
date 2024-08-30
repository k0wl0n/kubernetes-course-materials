# Kubernetes Manifest Repository

This repository contains Kubernetes manifests organized by chapter and topics to support learning and understanding of various Kubernetes features and concepts. Each directory contains specific examples and configurations related to the chapter topics.

## [All-in-One Package] Docker & Kubernetes: From Containerization to Real-World Applications (ft. AWS EKS)

🔗 Find the full course here: [Course Link](https://abit.ly/1poqhq)

This course is designed to take you from the basics of containerization to real-world application implementation, including using AWS EKS. It’s the perfect opportunity for anyone looking to develop their DevOps skills and master Kubernetes in a practical and applicable context.

### What will you learn?

- The basics of Docker and Kubernetes
- Building and managing Kubernetes clusters using Terraform and AWS
- Building CICD pipeline for Software Development Life Cycle using GitHub Actions
- Kubernetes observability with tools like Grafana, Loki, and Jaeger
- Real-world Go application implementation with AWS EKS

### About Fast Campus

Fast Campus is Korea's No. 1 edutech company with over 930,000 cumulative members. Recently, Fast Campus launched in Indonesia with many high-quality online lectures for working professionals.

## Directory Structure

The repository is organized into several chapters, each focusing on different aspects of Kubernetes.

## Chapters and Topics

### P2-Ch.02
- **P2-Ch.02-2**: Introduction to Kubernetes basics.

### P2-Ch.03
- **P2-Ch.03-1-3.Docker Image**: Building and managing Docker images.
  - `go-echo`: Basic Go application.
  - `go-echo-v1`: First version of the Go echo application.
  - `go-echo-v2`: Second version with enhancements.
  - `go-echo-v3`: Third version with additional features.

### P2-Ch.04
- **P2-Ch.04-2-5.Pods**: Basic pod management and configuration.
- **P2-Ch.04-6.ReplicaSet**: Ensuring the availability of pod replicas.
- **P2-Ch.04-7-8.Deployment**: Managing Deployments.
  - `go-echo-v3`: Example Deployment using the `go-echo-v3` app.
- **P2-Ch.04-9-10.DaemonSet**: Configuration and use cases for DaemonSets.
- **P2-Ch.04-11-12.Job**: Running batch Jobs in Kubernetes.
- **P2-Ch.04-13-14.CronJob**: Scheduling recurring tasks using CronJobs.
- **P2-Ch.04-15-16.Statefulset**: Managing stateful applications with StatefulSets.
- **P2-Ch.04-17.RolloutRollback**: Rollout and rollback strategies.
  - `go-echo-v4`: Example for demonstrating rollout and rollback.
- **P2-Ch.04-18-19.ConfigMap**: Using ConfigMaps to manage configuration data.
- **P2-Ch.04-20.Secret**: Storing and managing sensitive information with Secrets.
- **P2-Ch.04-21.HPA**: Horizontal Pod Autoscaling for workload management.
- **P2-Ch.04-22.Mengelola Resource Limits**: Managing resource limits for pods.
- **P2-Ch.04-23.Pengelolaan Scheduling Pod**: Advanced pod scheduling techniques.

### P2-Ch.05
- **P2-Ch.05-1.Memahami Konektivitas Pod**: Understanding pod connectivity.
- **P2-Ch.05-2.Layanan ClusterIP**: ClusterIP services for internal communication.
- **P2-Ch.05-3.Layanan NodePort**: Exposing services using NodePort.
- **P2-Ch.05-4.Layanan LoadBalancer**: Load balancing external traffic.
- **P2-Ch.05-5.Pengenalan Pengendali Ingress**: Introduction to Ingress controllers.

### P2-Ch.06
- **P2-Ch.06-2.Mengkonfigurasi Storage Class**: Configuring storage classes in Kubernetes.
- **P2-Ch.06-4.Bekerja dengan Persistance Volume**: Working with persistent volumes.

## How to Use This Repository

1. **Clone the repository to your local machine:**
    ```bash
    git clone https://github.com/k0wl0n/kubernetes-course-manifest.git
    ```

2. **Navigate to the desired chapter and sub-chapter to explore the examples and manifests:**
    ```bash
    cd kubernetes-manifest/P2-Ch.04/P2-Ch.04-2-5.Pods
    ```

3. **Apply the Kubernetes manifests using `kubectl`:**
    ```bash
    kubectl apply -f <manifest-file>.yaml
    ```

4. **Explore and modify the examples as needed for learning and experimentation.**

## License

This repository is licensed under the MIT License. See the `LICENSE` file for more information.

## Contact

For questions or discussions, please open an issue or reach out via this repo.
