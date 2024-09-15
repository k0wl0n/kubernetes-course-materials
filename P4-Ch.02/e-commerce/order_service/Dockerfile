
FROM golang:1.22.4-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

FROM alpine
COPY --from=builder /app/main /app/main
EXPOSE 8080
ENTRYPOINT ["/app/main"]
