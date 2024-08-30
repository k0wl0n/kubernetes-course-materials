### P2-Ch.05-5.Menggunakan Ingress di Minikube

This guide demonstrates how to enable and use the NGINX Ingress controller in Minikube to manage external access to services running in a Kubernetes cluster.

### Steps

#### 1. Enable NGINX Ingress Controller

Enable the NGINX Ingress controller addon in Minikube to manage external traffic.

```bash
minikube addons enable ingress
```

#### 2. Verify the NGINX Ingress Controller

Check if the NGINX Ingress controller is running by listing the pods in the `ingress-nginx` namespace.

```bash
kubectl get pods -n ingress-nginx
```

#### 3. Create a Deployment and Expose It

Create a Deployment using an example web server image and expose it via a NodePort service.

```bash
kubectl create deployment web --image=kowlon/go-echo:v2
kubectl expose deployment web --type=NodePort --port=8080
```

#### 4. Enable Minikube Tunnel

Run the Minikube tunnel to provide external access to the Ingress service.

```bash
minikube tunnel -c
```

#### 5. Access the Ingress

Use the following command to access the service through the Ingress using a hostname (`hello-world.example`) mapped to localhost.

```bash
curl --resolve "hello-world.example:80:127.0.0.1" -i http://hello-world.example
```

#### 6. Optionally Visit the Ingress from Your Browser

You can access the service using `hello-world.example` in your web browser by adding the appropriate entry to your `/etc/hosts` file (or the equivalent on Windows).

- **On Linux**:
  
  1. Get the Minikube IP:

     ```bash
     minikube ip
     ```

  2. Add this line to `/etc/hosts`:

     ```
     <minikube-ip> hello-world.example
     ```

- **On Mac**:

  Add this line to `/etc/hosts`:

  ```
  127.0.0.1 hello-world.example
  ```

- **On Windows**:

  Edit the file `C:\Windows\System32\drivers\etc\hosts` and add:

  ```
  127.0.0.1 hello-world.example
  ```

  or

  ```
  <minikube-ip> hello-world.example
  ```

This setup illustrates how to use Ingress in Minikube to manage routing of external traffic to services within a Kubernetes cluster, providing a scalable solution for handling external access to multiple services.