package broker

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
	URL string
}

func NewRabbitMQPublisher(url string) *RabbitMQPublisher {
	return &RabbitMQPublisher{URL: url}
}

func (r *RabbitMQPublisher) PublishEmailEvent(txnId, userId string) {
	conn, err := amqp.Dial(r.URL)
	if err != nil {
		log.Printf("❌ Failed to connect to RabbitMQ: %v", err)
		return
	}
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	msgBody := fmt.Sprintf(`{"user_id": "%s", "transaction_id": "%s", "event": "PAYMENT_SUCCESS"}`, userId, txnId)

	err = ch.Publish("", "email_notifications", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(msgBody),
	})

	if err == nil {
		log.Println("🚀 Message sent to RabbitMQ!")
	}
}
