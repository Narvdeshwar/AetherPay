FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .

WORKDIR /app/services/notification
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/notification-app main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/notification-app .

# No EXPOSE needed here
CMD ["./notification-app"]
