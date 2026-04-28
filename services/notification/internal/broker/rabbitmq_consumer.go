package broker

import (
	service "aetherpay/notification/internal/services"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
	URL     string
	Service *service.NotificationService
}

func NewRabbitMQConsumer(url string, svc *service.NotificationService) *RabbitMQConsumer {
	return &RabbitMQConsumer{
		URL:     url,
		Service: svc,
	}
}

func (c *RabbitMQConsumer) StartConsuming() {
	conn, err := amqp.Dial(c.URL)
	if err != nil {
		log.Fatal("❌ Unable to connect with rabbitmq: ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("❌ Unable to open the channel: ", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"email_notifications", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatal("❌ Unable to declare the queue: ", err)
	}

	msgs, err := ch.Consume(
		q.Name, "", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatal("❌ Unable to read the message: ", err)
	}

	log.Println("🔔 Notification worker is running.. Waiting for messages....")

	// Messages ko sunna aur Service ko pass karna
	var forever chan struct{}
	go func() {
		for d := range msgs {
			c.Service.SendEmailNotification(d.Body)
		}
	}()

	<-forever
}
