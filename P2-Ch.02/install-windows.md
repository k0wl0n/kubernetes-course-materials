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

