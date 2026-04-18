# 1. Build Stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o notification-app main.go

# 2. Final Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/notification-app .


# Application start karein
CMD ["./notification-app"]
