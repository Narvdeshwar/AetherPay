package main

import (
	"aetherpay/notification/internal/broker"
	service "aetherpay/notification/internal/services"
)

func main() {
	// 1. Initialize Service
	svc := service.NewNotificationService()

	// 2. Initialize Consumer (RabbitMQ container ka URL use karenge)
	consumer := broker.NewRabbitMQConsumer("amqp://guest:guest@rabbitmq:5672/", svc)

	// 3. Start the worker (Yeh infinite loop mein chalega)
	consumer.StartConsuming()
}
