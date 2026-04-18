# 1. Build Stage
FROM golang:1.21-alpine AS builder

# Working directory set karein
WORKDIR /app

# Go modules copy karein aur download karein
COPY go.mod go.sum ./
RUN go mod download

# Poora code copy karein
COPY . .

# Binary build karein (Isse 'billing-app' naam ki file banegi)
RUN CGO_ENABLED=0 GOOS=linux go build -o billing-app main.go

# 2. Final Stage (Choti image banane ke liye)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/billing-app .

# Port expose karein
EXPOSE 3001

# Application start karein
CMD ["./billing-app"]
