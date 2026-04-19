# 1. Build Stage
FROM golang:1.25-alpine AS builder

# Pura project Docker mein daalein
WORKDIR /app
COPY . .

# Payment folder mein jayein
WORKDIR /app/services/payment

# DIRECT BUILD (go mod download hata diya hai)
# Binary ko direct /app folder mein save kar rahe hain taaki dhoondhne mein aasan ho
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/payment-app main.go

# 2. Final Stage
FROM alpine:latest
WORKDIR /root/

# Binary ko /app se copy karein
COPY --from=builder /app/payment-app .

EXPOSE 3002
CMD ["./payment-app"]
