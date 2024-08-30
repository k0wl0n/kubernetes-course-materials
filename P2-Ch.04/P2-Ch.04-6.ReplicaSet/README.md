### README for `P2-Ch.04-6.ReplicaSet`

This guide provides steps to work with ReplicaSets and a Pod in Kubernetes within the `P2-Ch.04-6.ReplicaSet` folder, using Minikube for local development.

### Steps

#### 1. Apply First ReplicaSet

Create and apply the first ReplicaSet configuration:

```bash
kubectl apply -f rs-1.yaml
```

#### 2. Create and Access Alpine Pod

Deploy an Alpine Pod and access its shell:

```bash
kubectl apply -f alpine.yaml
kubectl exec -ti alpine sh
```

#### 3. Apply Second ReplicaSet

Create and apply the second ReplicaSet configuration:

```bash
kubectl apply -f rs-2.yaml
```

#### 4. Enable Minikube Tunnel

Run the Minikube tunnel to expose services:

```bash
minikube tunnel
```

#### 5. Loop to Access Service

Run a loop to continuously access a service running on localhost:

```bash
while true; do                
  curl -sSL http://localhost:8080/ && echo
  sleep 0.5
done
```

#### 6. Apply Third ReplicaSet

Create and apply the third ReplicaSet configuration:

```bash
kubectl apply -f rs-3.yaml
```

This setup demonstrates how to manage ReplicaSets, interact with a Pod, and use Minikube to expose services for testing and development.