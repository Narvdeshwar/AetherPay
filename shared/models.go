package shared

type PaymentRequest struct {
	UserId   string  `json:"user_id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type PaymentResponse struct{
	TransactionId string `json:"transaction_id"`
	Status string `json:"status"`
}
