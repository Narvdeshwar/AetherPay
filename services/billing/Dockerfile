FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .

WORKDIR /app/services/billing
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/billing-app main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/billing-app .

EXPOSE 3001
CMD ["./billing-app"]
