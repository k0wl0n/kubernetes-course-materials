# Stage 1: Build the application
FROM golang:1.22.4-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# Stage 2: Run the application
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/main /app/main

# Expose port 8080
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/app/main"]
