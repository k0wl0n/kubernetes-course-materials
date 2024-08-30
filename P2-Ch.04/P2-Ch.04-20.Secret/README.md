### README for `P2-Ch.04-20.Secret`

This guide provides steps to create and use Kubernetes Secrets to store and manage sensitive data, such as passwords and configuration files. It includes examples of how to create Secrets and use them in Pods within the `P2-Ch.04-20.Secret` folder.

### Overview

This project demonstrates the use of Kubernetes Secrets to store sensitive information securely and how to use these Secrets in Pods. Examples include creating Secrets from base64-encoded data and mounting Secrets as volumes in Pods.

### Files Overview

1. **`secret.yaml`**: Defines a Secret named `secret` with base64-encoded username and password.
2. **`secret-pod.yaml`**: Configures a Pod named `alpine-secret-pod` that uses the `secret` Secret to access username and password data.
3. **`multi-secret.yaml`**: Defines a Secret named `multi-secret` containing multiple configuration files with base64-encoded data.
4. **`multi-secret-pod.yaml`**: Configures a Pod named `alpine-multi-secret-pod` that uses the `multi-secret` Secret to access multiple configuration files.

### Details

#### 1. Secret: `secret.yaml`

Defines a Secret to store sensitive data like username and password in base64-encoded format.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: secret
type: Opaque
data:
  username: YWRtaW4=  # base64 encoded "admin"
  password: MWYyZDFlMmU2N2Rm  # base64 encoded "1f2d1e2e67df"
```

#### 2. Pod: `secret-pod.yaml`

Configures a Pod to use the `secret` Secret by mounting it as a volume, allowing access to sensitive information.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: alpine-secret-pod
spec:
  containers:
  - name: alpine
    image: alpine
    command: ["/bin/sh", "-c", "cat /etc/secrets/username; cat /etc/secrets/password"]
    volumeMounts:
    - name: secret-volume
      mountPath: /etc/secrets
      readOnly: true
  volumes:
  - name: secret-volume
    secret:
      secretName: secret
```

#### 3. Secret: `multi-secret.yaml`

Stores multiple pieces of sensitive information, such as application and database properties, in a single Secret.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: multi-secret
type: Opaque
data:
  app.properties: QUFQX0NPTElSPWJsdWUKQVBQX01PREU9cHJvZHVjdGlvbg==  # base64 encoded values
  db.properties: REJfSE9TVD1sb2NhbGhvc3QKREJfUE9SVD01NDMy
```

#### 4. Pod: `multi-secret-pod.yaml`

Configures a Pod to use the `multi-secret` Secret, mounting it as multiple volumes for accessing different pieces of sensitive data.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: alpine-multi-secret-pod
spec:
  containers:
  - name: alpine
    image: alpine
    command: ["/bin/sh", "-c", "cat /etc/secrets/app.properties; cat /etc/secrets/db.properties"]
    volumeMounts:
    - name: app-secret-volume
      mountPath: /etc/secrets/app.properties
      subPath: app.properties
      readOnly: true
    - name: db-secret-volume
      mountPath: /etc/secrets/db.properties
      subPath: db.properties
      readOnly: true
  volumes:
  - name: app-secret-volume
    secret:
      secretName: multi-secret
  - name: db-secret-volume
    secret:
      secretName: multi-secret
```

### Usage

- To create these Secrets and Pods in your Kubernetes cluster, run the following commands:

  ```bash
  kubectl apply -f secret.yaml
  kubectl apply -f secret-pod.yaml
  kubectl apply -f multi-secret.yaml
  kubectl apply -f multi-secret-pod.yaml
  ```

- To view the created Secrets, use:

  ```bash
  kubectl get secrets
  ```

- To view the contents of a Secret, use:

  ```bash
  kubectl describe secret <secret-name>
  ```

Replace `<secret-name>` with the name of the Secret you want to inspect.

This setup demonstrates how to securely manage sensitive data using Kubernetes Secrets and shows how to use these Secrets in Pod configurations to access sensitive information safely.