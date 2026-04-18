# Variables
DOCKER_COMPOSE = docker-compose
GO = go

# 1. Infrastructure Commands
up:
	@echo "🚀 Starting infrastructure (Postgres, RabbitMQ, Kafka, ClickHouse)..."
	$(DOCKER_COMPOSE) up -d

down:
	@echo "🛑 Stopping all containers..."
	$(DOCKER_COMPOSE) down

# 2. Local Run Commands (Jab tak hum Dockerfiles nahi banate)
run-payment:
	@echo "💳 Running Payment Service..."
	cd services/payment && $(GO) run main.go

run-billing:
	@echo "🧾 Running Billing Service..."
	cd services/billing && $(GO) run main.go

run-notification:
	@echo "🔔 Running Notification Service..."
	cd services/notification && $(GO) run main.go

# Sab kuch ek sath chalane ke liye (Alag terminals ki zaroorat nahi padegi)
# Note: Iske liye aapko 'air' ya 'concurrently' jaise tools chahiye honge,
# par abhi ke liye hum manual run karenge.

# 3. Database & Tools Shortcuts
postgres:
	@echo "🐘 Entering Postgres CLI..."
	docker exec -it $$(docker ps -qf "name=postgres") psql -U admin -d postgres

clickhouse:
	docker exec -it $$(docker ps -qf "name=clickhouse") clickhouse-client

# 4. Testing Shortcut
test-payment:
	curl -X POST http://localhost:3001/api/v1/billing/subscribe \
	-H "Content-Type: application/json" \
	-d '{"user_id": "123", "amount": 2001.4, "currency": "INR"}'

# 5. Help (Default command)
help:
	@echo "AetherPay Management Commands:"
	@echo "  make up              - Start Docker infrastructure"
	@echo "  make down            - Stop Docker containers"
	@echo "  make run-payment     - Run Payment Service locally"
	@echo "  make run-billing     - Run Billing Service locally"
	@echo "  make run-notification- Run Notification Service locally"
	@echo "  make clickhouse      - Enter ClickHouse CLI"
	@echo "  make test-payment    - Send a test payment request"
