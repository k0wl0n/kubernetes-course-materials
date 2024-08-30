### P2-Ch.05-3.Layanan NodePort

This project demonstrates the use of a NodePort Service to expose a web server application to external traffic in a Kubernetes cluster.

### Files Overview

1. **`webserver.yaml`**: Defines a Deployment for an Nginx web server and a NodePort Service to expose it.

### Nginx Deployment and NodePort Service

- **Deployment**: Deploys two replicas of an Nginx container to serve web traffic.
  - **Image**: Uses `nginx:alpine` for a lightweight Nginx server.
  - **Ports**: Exposes container port 80.
  - **Resources**: Sets CPU and memory requests and limits.

- **Service**: Exposes the Nginx Deployment using a NodePort Service, making it accessible outside the cluster on a specific port.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx-nodeport
spec:
  type: NodePort
  selector:
    app: nginx
  ports:
  - protocol: TCP
    port: 80
    targetPort: 80
    nodePort: 32080
```

### Usage

1. **Deploy the Nginx Web Server and NodePort Service**:

   Apply the `webserver.yaml` file to create the Nginx Deployment and expose it via a NodePort Service.

   ```bash
   kubectl apply -f webserver.yaml
   ```

2. **Access the Nginx Service**:

   Use Minikube to get the URL of the exposed Nginx service. This command will provide a URL that can be accessed from your browser or a curl command to test connectivity.

   ```bash
   minikube service nginx-nodeport --url
   ```

3. **Open the Provided URL**:

   Copy and paste the provided URL into a web browser or use a tool like `curl` to access the Nginx web server. This will demonstrate external access to the service using the NodePort.

This setup showcases how to use a NodePort Service to expose internal applications to external network traffic, allowing access to applications running inside a Kubernetes cluster from outside the cluster.