# variables
DOCKER_COMPOSE= docker-compose
GO= go

# 1. infrastructure command
up:
	@echo "=============== Starting infrastructure (Postgres, RabbitMQ, Kafka, ClickHouse) ================="
	$(DOCKER_COMPOSE) up -d

down:
	@echo "=================== Stopping all containers ================================"
	$(DOCKER_COMPOSE) down


