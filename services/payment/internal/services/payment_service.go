package service

import (
	"aetherpay/payment/internal/broker"
	"aetherpay/payment/internal/repository"
	"aetherpay/shared"
)

type PaymentService struct {
	Repo   *repository.PaymentRepository
	Kafka  *broker.KafkaPublisher
	Rabbit *broker.RabbitMQPublisher
}

func NewPaymentService(repo *repository.PaymentRepository, k *broker.KafkaPublisher, r *broker.RabbitMQPublisher) *PaymentService {
	return &PaymentService{
		Repo:   repo,
		Kafka:  k,
		Rabbit: r,
	}
}

func (s *PaymentService) ExecutePayment(req shared.PaymentRequest) (string, error) {
	txnId := "1x23fdfg" // Isko aage dynamic karenge

	// 1. Database mein save karein
	newTxn := shared.Transaction{
		TransactionId: txnId,
		UserId:        req.UserId,
		Amount:        req.Amount,
		Currency:      req.Currency,
		Status:        "SUCCESS",
	}

	if err := s.Repo.SaveTransaction(&newTxn); err != nil {
		return "", err
	}

	// 2. Background Events Trigger Karein (Goroutines)
	go s.Rabbit.PublishEmailEvent(txnId, req.UserId)
	go s.Kafka.PublishPaymentEvent(txnId, req.UserId, req.Amount, req.Currency)

	return txnId, nil
}
