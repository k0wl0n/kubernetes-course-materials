### installing brew
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

### installing minikube
brew install minikube
minikube version

### Install kubectl
brew install kubectl
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
