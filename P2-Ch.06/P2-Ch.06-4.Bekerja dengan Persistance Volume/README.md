### P2-Ch.06-4.Bekerja dengan Persistent Volume

This project demonstrates how to work with Persistent Volumes (PV), Persistent Volume Claims (PVC), and StatefulSets in Kubernetes to manage persistent storage for applications.

### Files Overview

1. **`pv.yaml`**: Defines a PersistentVolume for storing data.
2. **`pvc.yaml`**: Defines a PersistentVolumeClaim to request storage from a PersistentVolume.
3. **`deploy.yaml`**: Defines a Pod configuration that uses a PVC to mount persistent storage.
4. **`go-sts.yaml`**: Defines a StatefulSet and a PersistentVolume to manage stateful applications with persistent storage.

### Persistent Volume and Claim Configuration

#### 1. Persistent Volume (PV)

- **File**: `pv.yaml`
- **Name**: `webserver-pv`
- **Capacity**: 1Gi
- **Access Modes**: `ReadWriteOnce` — Allows read/write access by a single node.
- **Storage Path**: `/mnt/data` on the host system.

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: webserver-pv
spec:
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: /mnt/data
```

#### 2. Persistent Volume Claim (PVC)

- **File**: `pvc.yaml`
- **Name**: `webserver-pvc`
- **Requested Storage**: 1Gi
- **Access Modes**: `ReadWriteOnce`

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: webserver-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
```

### Pod Using PVC

- **File**: `deploy.yaml`
- **Pod Name**: `webserver-pod`
- **Container Image**: `nginx`
- **Volume Mount**: Mounts the PVC `webserver-pvc` to `/usr/share/nginx/html` to serve static files from persistent storage.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: webserver-pod
spec:
  containers:
    - name: nginx
      image: nginx
      resources:
        requests:
          cpu: "10m"
          memory: "32Mi"
        limits:
          cpu: "100m"
          memory: "128Mi"
      volumeMounts:
        - mountPath: "/usr/share/nginx/html"
          name: webserver-volume
  volumes:
    - name: webserver-volume
      persistentVolumeClaim:
        claimName: webserver-pvc
```

### StatefulSet and Persistent Volume

#### 1. Persistent Volume for StatefulSet

- **File**: `go-sts.yaml`
- **Name**: `go-echo-stateful-volume`
- **Capacity**: 10Mi
- **Access Modes**: `ReadWriteOnce`
- **Storage Path**: `/data/location` on the host system.

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: go-echo-stateful-volume
spec:
  accessModes:
    - ReadWriteOnce
  capacity:
    storage: 10Mi
  hostPath:
    path: /data/location
```

#### 2. StatefulSet Configuration

- **StatefulSet Name**: `go-echo-stateful`
- **Replicas**: 3
- **Container Image**: `kowlon/go-echo:v2`
- **Volume Mount**: Each replica mounts its own persistent storage to `/app/data`.

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: go-echo-stateful
  labels:
    name: go-echo-stateful
spec:
  serviceName: go-echo-stateful-service
  replicas: 3
  selector:
    matchLabels:
      name: go-echo-stateful
  template:
    metadata:
      labels:
        name: go-echo-stateful
    spec:
      containers:
      - name: go-echo-stateful
        image: kowlon/go-echo:v2
        env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        volumeMounts:
        - name: go-echo-stateful-volume-claim
          mountPath: /app/data
  volumeClaimTemplates:
  - metadata:
      name: go-echo-stateful-volume-claim
    spec:
      accessModes:
        - ReadWriteOnce
      resources:
        requests:
          storage: 10Mi
```

### Usage

1. **Create Persistent Volume and Claim**:

   Apply the PV and PVC configurations to create the necessary storage resources.

   ```bash
   kubectl apply -f pv.yaml
   kubectl apply -f pvc.yaml
   ```

2. **Deploy the Webserver Pod**:

   Apply the Pod configuration to deploy a webserver that uses the persistent storage.

   ```bash
   kubectl apply -f deploy.yaml
   ```

3. **Deploy the StatefulSet**:

   Apply the StatefulSet configuration to deploy a stateful application with persistent storage.

   ```bash
   kubectl apply -f go-sts.yaml
   ```

4. **Verify the Setup**:

   Check the status of the PV, PVC, Pods, and StatefulSets to ensure they are correctly bound and running.

   ```bash
   kubectl get pv
   kubectl get pvc
   kubectl get pods
   kubectl get statefulsets
   ```

This setup demonstrates the use of Persistent Volumes and StatefulSets to manage persistent data in Kubernetes, ensuring data remains available even when Pods are deleted or rescheduled.