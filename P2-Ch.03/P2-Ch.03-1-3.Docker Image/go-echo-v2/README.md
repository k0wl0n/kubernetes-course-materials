### Init go project

go mod init go-echo

go mod tidy

### Build Docker Image

docker build -t go-echo .

### Run Docker Image

docker run -p 8080:8080 go-echo

### push docker image
#### Register to dockerhub
https://hub.docker.com/

#### Docker Login
docker login
insert user and password

#### Tag and push your Docker image
docker tag local-image:tagname new-repo:tagname
docker push new-repo:tagname

kowlon/go-echo

docker tag go-echo kowlon/go-echo:tagname
docker push kowlon/go-echo:tagname

### Build multiple Architecture

#### Install buildx
https://github.com/docker/buildx?tab=readme-ov-file#installing

#### Create a builder
docker buildx create --name mybuilder

#### Use a builder
docker buildx use mybuilder

#### Register to dockerhub
https://hub.docker.com/

#### Docker Login
docker login
insert user and password

#### Build Docker and Push to docker hub 

docker buildx build --push --tag [docker-hubid/image-tag] --platform=linux/arm64,linux/amd64 .

example
docker buildx build --push --tag kowlon/go-echo:v1 --platform=linux/arm64,linux/amd64 .
