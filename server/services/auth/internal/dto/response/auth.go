package response

import "time"

type MerchantResponse struct {
	MerchantID   string    `json:"merchant_id"`
	TenantID     string    `json:"tenant_id"`
	MerchantName string    `json:"merchant_name"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
}

type RegisterResponse struct {
	MerchantResponse
	Message string `json:"message"`
}

type ProfileResponse struct {
	MerchantResponse
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TenantID    string `json:"tenant_id"`
	ExpiresIn   int64  `json:"expires_in"`
}
