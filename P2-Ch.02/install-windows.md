### Open Power Shell with Run As Administrator
### Run this for installing Chocolatey

Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))

### Install kubectl
choco install kubernetes-cli

### Verify the installation
kubectl version --client

### Install minikube
choco install minikube

### Install kubectx
choco install kubectx

### Install kubens
choco install kubens

### Start Minikube with Specific Kubernetes Version
By default, Minikube starts the latest supported Kubernetes version. To use a specific version (e.g., v1.30.0), use the `--kubernetes-version` flag:

```bash
minikube start --kubernetes-version=v1.30.0
```
