package service

import (
	"encoding/json"
	"log"
)

// Event ki structure jo RabbitMQ se aayegi
type PaymentEvent struct {
	UserID        string `json:"user_id"`
	TransactionID string `json:"transaction_id"`
	Event         string `json:"event"`
}

type NotificationService struct{}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (s *NotificationService) SendEmailNotification(data []byte) {
	var event PaymentEvent

	// JSON ko Go struct mein convert karein
	if err := json.Unmarshal(data, &event); err != nil {
		log.Printf("❌ Failed to parse event data: %v\n", err)
		return
	}

	// Asli Email bhejne ka logic (Abhi ke liye sirf print)
	log.Printf("📧 Sending Email to User [%s] for Transaction [%s]\n", event.UserID, event.TransactionID)
	log.Println("✅ Email sent successfully!")
	log.Println("------------------------------------")
}
