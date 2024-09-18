To set up Docker `buildx`, which allows for extended features such as multi-platform builds and cache management, follow the steps below. Docker `buildx` is an extension of the Docker build command, making it easier to build images for different platforms (e.g., Linux, Windows, ARM) and use advanced build features.

### Steps to Set Up Docker `buildx`:

#### 1. Install Docker (If Not Installed)
If you haven't installed Docker yet, you can follow the instructions to install Docker based on your operating system.

- **For Linux**:
  ```bash
  sudo apt update
  sudo apt install docker.io
  sudo systemctl start docker
  sudo systemctl enable docker
  sudo usermod -aG docker $USER  # Add your user to the docker group
  ```

- **For macOS and Windows**:
  Install Docker Desktop from [Docker's official site](https://www.docker.com/products/docker-desktop/).

#### 2. Enable Docker `buildx`
If Docker is installed, `buildx` comes included starting with Docker 19.03. To enable it, you need to follow these steps.

- Run the following command to enable Docker `buildx` as the default builder:
  ```bash
  docker buildx create --use
  ```

  You should see a message like `buildx-builder` created and set as the current builder.

#### 3. Verify Docker `buildx` Setup
After creating the buildx instance, verify that it is working properly by running:

```bash
docker buildx inspect --bootstrap
```

This will provide details about your current `buildx` setup.

You can also list all available `buildx` instances by running:

```bash
docker buildx ls
```

#### 4. Example Usage of Docker `buildx`
Now that `buildx` is set up, you can build multi-platform images using the following command:

```bash
docker buildx build --platform linux/amd64,linux/arm64 -t your-image:tag .
```

In this example:
- `--platform` specifies the platforms for which you want to build the image (e.g., `linux/amd64`, `linux/arm64`).
- `-t` specifies the tag for the image.

If you want to push the built image to a Docker registry (like Docker Hub), append the `--push` flag:

```bash
docker buildx build --platform linux/amd64,linux/arm64 -t your-image:tag --push .
```

#### 5. (Optional) Create a Custom Builder Instance
If you want to create a custom builder instance, run:

```bash
docker buildx create --name mybuilder --driver docker-container --use
docker buildx inspect mybuilder --bootstrap
```

This creates a custom `buildx` builder named `mybuilder` and sets it as the current builder. You can now use this builder for your builds.

#### 6. Using Build Caching with `buildx`
To enable caching and speed up builds, use `buildx` cache options like:

```bash
docker buildx build --cache-from type=registry,ref=your-image:cache --cache-to type=inline --platform linux/amd64 -t your-image:tag .
```

This command uses the cache stored in the Docker registry to accelerate builds.

### Common Use Cases
- **Multi-platform Builds**: Easily build images for different architectures like ARM, x86.
- **Build Caching**: Speed up builds with caching strategies.
- **Distributed Builds**: Use multiple nodes to parallelize builds.



docker buildx build --push --tag kowlon/vault-env-parser --platform=linux/arm64,linux/amd64 .