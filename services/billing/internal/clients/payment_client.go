package clients

import (
	"aetherpay/shared"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PaymentClient struct {
	BaseURL string
}

func NewPaymentClient(baseURL string) *PaymentClient {
	return &PaymentClient{BaseURL: baseURL}
}

func (c *PaymentClient) ProcessPayment(paymentReq shared.PaymentRequest) (*shared.PaymentResponse, error) {
	reqData, err := json.Marshal(paymentReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payment request: %v", err)
	}

	url := fmt.Sprintf("%s/api/v1/payment/process", c.BaseURL)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqData))
	if err != nil {
		return nil, fmt.Errorf("failed to call payment service: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // io.ReadAll is the modern standard
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var paymentResp shared.PaymentResponse
	if err := json.Unmarshal(body, &paymentResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payment response: %v", err)
	}

	return &paymentResp, nil
}