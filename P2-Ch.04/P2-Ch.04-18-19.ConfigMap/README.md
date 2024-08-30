### README for `P2-Ch.04-18-19.ConfigMap`

This guide provides steps to create and use ConfigMaps in Kubernetes, demonstrating how to manage application configuration using different methods within the `P2-Ch.04-18-19.ConfigMap` folder.

### Overview

This project demonstrates various ways to create and use ConfigMaps in Kubernetes. The examples include creating ConfigMaps from literal values, using ConfigMaps in Pods, and mounting multiple ConfigMaps into a single Pod.

### Files

- **`alpine-cm-one.yaml`**: ConfigMap containing application properties.
- **`alpine-cm-two.yaml`**: ConfigMap containing a simple shell script.
- **`alpine-pod.yaml`**: Pod configuration that mounts the ConfigMaps `alpine-cm-one` and `alpine-cm-two`.
- **`cm-env.yaml`**: ConfigMap with environment variables.
- **`cm-pod.yaml`**: Pod configuration that uses the `config-env` ConfigMap for environment variables.
- **`multi-cm.yaml`**: ConfigMap with multiple configuration files for application and database properties.
- **`multi-pod.yaml`**: Pod configuration that mounts multiple ConfigMap files.

### Commands

#### 1. Create a ConfigMap from Literal Values

Create a ConfigMap directly from the command line using literal values:

```bash
kubectl create configmap my-config --from-literal=APP_COLOR=blue --from-literal=APP_MODE=production
```

#### 2. Apply ConfigMaps and Pods

Deploy the ConfigMaps and Pods by applying the corresponding YAML files:

```bash
kubectl apply -f alpine-cm-one.yaml
kubectl apply -f alpine-cm-two.yaml
kubectl apply -f alpine-pod.yaml
kubectl apply -f cm-env.yaml
kubectl apply -f cm-pod.yaml
kubectl apply -f multi-cm.yaml
kubectl apply -f multi-pod.yaml
```

#### 3. Check ConfigMaps and Pods

List the ConfigMaps and Pods to ensure they have been created successfully:

```bash
kubectl get configmap
kubectl get pod
```

#### 4. Describe ConfigMaps and Pods

Get detailed information about specific ConfigMaps and Pods:

```bash
kubectl describe configmap <configmap-name>
kubectl describe pod <pod-name>
```

Replace `<configmap-name>` and `<pod-name>` with the actual names of the ConfigMaps and Pods you wish to inspect.

### Summary

- **`alpine-cm-one`**: Contains key-value pairs for application configuration.
- **`alpine-cm-two`**: Contains a shell script as a data value.
- **`alpine-pod`**: Uses both `alpine-cm-one` and `alpine-cm-two` by mounting them as volumes.
- **`config-env`**: Provides environment variables for the container.
- **`cm-pod`**: Uses the `config-env` ConfigMap to set environment variables in a container.
- **`multi-cm`**: Combines application and database configuration into a single ConfigMap.
- **`multi-pod`**: Mounts both application and database configuration files from `multi-cm`.

This setup demonstrates how to manage application configuration using Kubernetes ConfigMaps, providing flexibility and scalability for handling configuration data.