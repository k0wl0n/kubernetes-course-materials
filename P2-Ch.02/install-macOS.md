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



