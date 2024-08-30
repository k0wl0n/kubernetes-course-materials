### P2-Ch.05-4.Layanan LoadBalancer

This project demonstrates how to use a LoadBalancer Service to expose a backend application to external traffic in a Kubernetes cluster, allowing it to be accessed from outside the cluster.

### Files Overview

1. **`lb.yaml`**: Defines a Deployment for a backend application and a LoadBalancer Service to expose it.

### Backend Deployment and LoadBalancer Service

- **Deployment**: Deploys two replicas of a backend container running a simple web server.
  - **Image**: Uses `kowlon/go-echo:v2` to serve HTTP responses.
  - **Ports**: Exposes container port 8080.
  - **Resources**: Sets CPU and memory requests and limits.

- **Service**: Exposes the backend Deployment using a LoadBalancer Service, making it accessible from outside the Kubernetes cluster.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: backend-service
spec:
  type: LoadBalancer
  selector:
    app: backend
  ports:
    - protocol: TCP
      port: 8080       
      targetPort: 8080
```

### Usage

1. **Deploy the Backend Application and LoadBalancer Service**:

   Apply the `lb.yaml` file to create the backend Deployment and expose it via a LoadBalancer Service.

   ```bash
   kubectl apply -f lb.yaml
   ```

2. **Enable Minikube Tunnel**:

   Use Minikube to create a network tunnel that allows access to the LoadBalancer Service. This command is required for exposing LoadBalancer services in Minikube.

   ```bash
   minikube tunnel -c
   ```

3. **Access the Backend Service**:

   After running the Minikube tunnel, you can use the external IP provided by Minikube to access the backend service from your browser or using a curl command to test connectivity.

   ```bash
   minikube service backend-service --url
   ```

4. **Open the Provided URL**:

   Copy and paste the URL into a web browser or use a tool like `curl` to interact with the backend service. This demonstrates external access to the service using the LoadBalancer.

This setup showcases how to use a LoadBalancer Service to provide external access to applications running inside a Kubernetes cluster, allowing them to be accessed from outside the cluster network.