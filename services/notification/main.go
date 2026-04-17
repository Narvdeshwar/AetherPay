package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// 1. Rabbitmq se connection setup krna
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Unable to connect with rabbitmq")
		return
	}
	defer conn.Close()

	// 2. channel open krna
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Unable to open the channel", err)
		return
	}
	defer ch.Close()

	// 3. Queue declare kro agr queue exits nhi krti hai to
	q, err := ch.QueueDeclare(
		"email_notifications", // queue ka naam
		true,                 // Durable (RabbitMQ restart hone par bhi bachegi
		false,                // Delete when unused
		false,                // Exclusive
		false,                // No wait
		nil,                  // Arguments
	)

	if err != nil {
		log.Fatal("Unable to declare the queue", err)
	}
	// 4. Queue se messages read karna
	msgs, err := ch.Consume(
		q.Name, // Queue
		"",     // Consumer ID
		true,   // Auto-Ack (Message padhte hi queue se delete ho jayega)
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Args
	)

	if err != nil {
		log.Fatal("Unable to read the messgae", err)
	}
	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("Received a message %s\n", d.Body)
			log.Printf("Sending Email to user... Done!\n")
			log.Println("------------------------------------")
		}
	}()
	log.Println("Notification services is running.. Waiting for messages....")
	<-forever
}
