package service

import (
	"aetherpay/billing/internal/clients"
	"aetherpay/shared"
)

type BillingService struct {
	PaymentClient *clients.PaymentClient
}

func NewBillingService(pc *clients.PaymentClient) *BillingService {
	return &BillingService{PaymentClient: pc}
}

func (s *BillingService) SubscribeUser(req shared.PaymentRequest) (*shared.PaymentResponse, error) {
	// Yahan aap future mein check kar sakte hain ki User pehle se subscribed toh nahi hai
	
	// Payment process karne ke liye client ko call karein
	resp, err := s.PaymentClient.ProcessPayment(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}