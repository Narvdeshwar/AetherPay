package request

type EmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type RegisterRequest struct {
	EmailRequest
	MerchantName string `json:"merchant_name" binding:"required"`
	Password     string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	EmailRequest
	Password string `json:"password" binding:"required"`
}
