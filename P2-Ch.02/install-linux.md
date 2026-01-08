### Instalasi Minikube

curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
minikube version

### Install kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x kubectl
sudo mv kubectl /usr/local/bin/

### Verify the installation
kubectl version --client

### Install kubens & kubectx
sudo git clone https://github.com/ahmetb/kubectx /opt/kubectx
sudo ln -s /opt/kubectx/kubectx /usr/local/bin/kubectx
sudo ln -s /opt/kubectx/kubens /usr/local/bin/kubens

### Verify the installation:
kubectx --help
kubens --help

### Start Minikube with Specific Kubernetes Version
By default, Minikube starts the latest supported Kubernetes version. To use a specific version (e.g., v1.30.0), use the `--kubernetes-version` flag:

```bash
minikube start --kubernetes-version=v1.30.0
```
