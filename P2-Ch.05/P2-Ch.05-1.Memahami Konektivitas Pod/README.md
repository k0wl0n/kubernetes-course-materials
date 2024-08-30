### P2-Ch.05-1.Memahami Konektivitas Pod

This project demonstrates how to set up a simple web server using a Kubernetes Deployment and expose it through a Service. It also includes a client Pod to interact with the web server, illustrating basic Pod connectivity.

### Files Overview

1. **`deployment.yaml`**: Defines a Deployment for an Nginx web server and a Service to expose it.
2. **`client.yaml`**: Defines a client Pod that can interact with the web server.

### Deployment and Service Configuration

- **Deployment**: Deploys an Nginx container using the `nginx:alpine` image.
  - **Replicas**: 1
  - **Resources**: Sets CPU and memory requests and limits.
  - **Ports**: Exposes container port 80.

- **Service**: Exposes the Nginx web server on port 80 using a `ClusterIP` for internal cluster communication.

### Client Pod Configuration

- **Pod Name**: `web-client`
- **Image**: `alpine/curl` for running Curl commands.
- **Command**: Sleeps for 3600 seconds to keep the Pod running for interaction.
- **Resources**: Sets CPU and memory requests and limits.

### Usage

1. **Deploy the Web Server and Service**:

   Use the following command to create the Deployment and Service:

   ```bash
   kubectl apply -f deployment.yaml
   ```

2. **Deploy the Client Pod**:

   Use the following command to create the client Pod:

   ```bash
   kubectl apply -f client.yaml
   ```

3. **Interact with the Web Server**:

   You can use the `web-client` Pod to interact with the `web-server` Pod by executing commands like Curl to test connectivity and response.

This configuration helps understand basic Pod connectivity within a Kubernetes cluster, using a web server and a client Pod setup.