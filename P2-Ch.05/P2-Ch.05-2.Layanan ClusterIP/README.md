### P2-Ch.05-2.Layanan ClusterIP

This project demonstrates the use of a ClusterIP Service to enable communication between a frontend and backend application within a Kubernetes cluster.

### Files Overview

1. **`frontend.yaml`**: Defines a Deployment for a frontend application that uses a curl script to communicate with the backend service.
2. **`backend.yaml`**: Defines a Deployment for a backend application and a ClusterIP Service to expose the backend.

### Frontend Deployment and ConfigMap

- **Deployment**: Deploys a single replica of a frontend container that runs a curl command to interact with the backend.
  - **Image**: Uses `alpine` and installs `curl`.
  - **Volume**: Mounts a ConfigMap volume containing the `curl-script.sh` script.
  - **Resources**: Sets CPU and memory requests and limits.

- **ConfigMap**: Contains a script `curl-script.sh` that continuously sends HTTP requests to the backend service and prints the response.

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: curl-script-config
data:
  curl-script.sh: |
    #!/bin/sh
    while true; do
      response=$(curl -s http://backend-service:8080)
      echo "Received from backend: $response"
      sleep 5
    done
```

### Backend Deployment and Service

- **Deployment**: Deploys two replicas of a backend container running a simple web server.
  - **Image**: Uses `kowlon/go-echo:v2` to serve HTTP responses.
  - **Ports**: Exposes container port 8080.
  - **Resources**: Sets CPU and memory requests and limits.

- **Service**: Exposes the backend Deployment using a ClusterIP Service, making it accessible only within the cluster.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: backend-service
spec:
  selector:
    app: backend
  ports:
  - protocol: TCP
    port: 8080
    targetPort: 8080
  type: ClusterIP
```

### Usage

1. **Deploy the Backend and Service**:

   Apply the `backend.yaml` file to create the backend Deployment and expose it via a ClusterIP Service.

   ```bash
   kubectl apply -f backend.yaml
   ```

2. **Deploy the Frontend**:

   Apply the `frontend.yaml` file to create the frontend Deployment that will interact with the backend.

   ```bash
   kubectl apply -f frontend.yaml
   ```

3. **Monitor the Frontend Output**:

   View the logs of the frontend Pod to see the responses received from the backend service.

   ```bash
   kubectl logs -f <frontend-pod-name>
   ```

Replace `<frontend-pod-name>` with the actual name of the running frontend Pod. This setup illustrates the use of a ClusterIP Service to facilitate internal communication between frontend and backend applications within a Kubernetes cluster.