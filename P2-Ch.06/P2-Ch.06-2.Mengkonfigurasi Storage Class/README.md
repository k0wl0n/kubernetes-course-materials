### P2-Ch.06-2.Mengkonfigurasi Storage Class

This project demonstrates how to configure different StorageClasses and PersistentVolumeClaims (PVCs) in Kubernetes to manage storage using various provisioning policies and storage backends.

### Files Overview

1. **`retain-storage-class.yaml`**: Defines a StorageClass with a Retain reclaim policy using Minikube hostpath.
2. **`retain-storage-class-pvc.yaml`**: Defines a PVC that uses the `retain-storage-class`.
3. **`delete-storage-class.yaml`**: Defines a StorageClass with a Delete reclaim policy using Minikube hostpath.
4. **`delete-storage-class-pvc.yaml`**: Defines a PVC that uses the `delete-storage-class`.
5. **`cloud_aws.yaml`**: Defines a StorageClass for AWS using EBS with gp3 type.
6. **`cloud_azure.yaml`**: Defines a StorageClass for Azure using Azure Disk with Standard_LRS type.
7. **`cloud_gcp.yaml`**: Defines a StorageClass for GCP using Google Cloud Engine Persistent Disk (GCE PD) with pd-standard type.

### StorageClass and PVC Configuration

#### 1. Retain Storage Class

- **File**: `retain-storage-class.yaml`
- **Provisioner**: `k8s.io/minikube-hostpath`
- **Reclaim Policy**: `Retain` — This keeps the volume after the PVC is deleted.
- **Volume Binding Mode**: `Immediate`

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: retain-storage-class
provisioner: k8s.io/minikube-hostpath
reclaimPolicy: Retain
volumeBindingMode: Immediate
```

- **PVC**: Uses `retain-storage-class` for provisioning 1Gi storage with `ReadWriteOnce` access mode.

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc-retain
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
  storageClassName: retain-storage-class
```

#### 2. Delete Storage Class

- **File**: `delete-storage-class.yaml`
- **Provisioner**: `k8s.io/minikube-hostpath`
- **Reclaim Policy**: `Delete` — This deletes the volume when the PVC is deleted.
- **Volume Binding Mode**: `Immediate`

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: delete-storage-class
provisioner: k8s.io/minikube-hostpath
reclaimPolicy: Delete
volumeBindingMode: Immediate
```

- **PVC**: Uses `delete-storage-class` for provisioning 1Gi storage with `ReadWriteOnce` access mode.

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc-delete
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
  storageClassName: delete-storage-class
```

#### 3. Cloud Provider Storage Classes

- **AWS Storage Class**: `cloud_aws.yaml`
  - **Provisioner**: `kubernetes.io/aws-ebs`
  - **Type**: `gp3`
  - **Reclaim Policy**: `Retain`
  - **Allow Volume Expansion**: `true`

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: standard
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp3
reclaimPolicy: Retain
allowVolumeExpansion: true
```

- **Azure Storage Class**: `cloud_azure.yaml`
  - **Provisioner**: `kubernetes.io/azure-disk`
  - **SKU Name**: `Standard_LRS`
  - **Location**: `eastus`

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: azure-disk
provisioner: kubernetes.io/azure-disk
parameters:
  skuName: Standard_LRS
  location: eastus
  storageaccounttype: Standard_LRS
```

- **GCP Storage Class**: `cloud_gcp.yaml`
  - **Provisioner**: `kubernetes.io/gce-pd`
  - **Type**: `pd-standard`

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: gcp-pd
provisioner: kubernetes.io/gce-pd
parameters:
  type: pd-standard
  replication-type: none
```

### Usage

1. **Apply the StorageClass and PVC Files**:

   Use the following command to create the StorageClasses and PersistentVolumeClaims:

   ```bash
   kubectl apply -f <file-name>.yaml
   ```

   Replace `<file-name>` with the relevant file to apply the configuration.

2. **Check the Status of PVCs**:

   Verify the status of the PVCs to ensure they are bound to the appropriate PersistentVolumes.

   ```bash
   kubectl get pvc
   ```

3. **Inspect Persistent Volumes**:

   Check the created PersistentVolumes and their status.

   ```bash
   kubectl get pv
   ```

This setup helps manage and configure storage in Kubernetes, utilizing different reclaim policies and storage types suitable for various cloud providers and local environments using Minikube.