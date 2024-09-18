# Docker build

docker build -t ecommerce-user .

# run in local
docker run -p 8001:8080 -v $(pwd)/.env-dev:/app/.env ecommerce-user:busybox

docker tag ecommerce-user:latest kowlon/ecommerce-user:latest
docker push kowlon/ecommerce-user:latest

# Docker Build with multi plaform

## create buildx

docker buildx create --name builder --use
docker buildx inspect --bootstrap

## build multiplatform
docker buildx build --tag ecommerce-user:latest --platform=linux/arm64,linux/amd64 .

## build multiplatform and push
docker buildx build --push --tag kowlon/ecommerce-user --platform=linux/arm64,linux/amd64 .

## push to ECR

docker tag ecommerce-user:stg-v.0.0.5 490004610217.dkr.ecr.ap-southeast-1.amazonaws.com/ecommerce-user:stg-v.0.0.5
docker push 490004610217.dkr.ecr.ap-southeast-1.amazonaws.com/ecommerce-user:stg-v.0.0.5

## push to ECR using buildx
docker buildx build --push --tag 490004610217.dkr.ecr.ap-southeast-1.amazonaws.com/ecommerce-user:stg-v.0.0.6 --platform=linux/arm64,linux/amd64 .