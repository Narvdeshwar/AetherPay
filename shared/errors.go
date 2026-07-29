package shared

import "fmt"

// APIError represents a standardized API error response.
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Tenant  string `json:"tenant_id,omitempty"`
}

// Error implements the built-in error interface.
func (e *APIError) Error() string {
	return fmt.Sprintf("%s has an error -> %d with message-> %s",
		e.Tenant,
		e.Code,
		e.Message,
	)
}

// NewAPIError creates and returns a new APIError instance.
func NewAPIError(code int, message string, tenant string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
		Tenant:  tenant,
	}
}
