package broker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaPublisher struct {
	Writer *kafka.Writer
}

func NewKafkaPublisher(brokerAddress string) *KafkaPublisher {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(brokerAddress),
		Topic:                  "payment_events",
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}
	return &KafkaPublisher{Writer: writer}
}

func (k *KafkaPublisher) PublishPaymentEvent(txnId, userId string, amount float64, currency string) {
	kafkaMessage := fmt.Sprintf(`{"transaction_id": "%s", "user_id": "%s", "amount": %.2f, "currency": "%s", "status": "SUCCESS", "timestamp": "%s"}`,
		txnId, userId, amount, currency, time.Now().Format(time.RFC3339))

	err := k.Writer.WriteMessages(context.Background(),
		kafka.Message{Key: []byte(txnId), Value: []byte(kafkaMessage)},
	)
	if err != nil {
		log.Printf("❌ Failed to publish to Kafka: %v", err)
	} else {
		log.Println("📊 Event saved in Kafka Ledger!")
	}
}
