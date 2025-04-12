package response

// ErrorResponse represents the standard error response format
type ErrorResponse struct {
	Error APIError `json:"error"`
}

// APIError represents a standardized error response
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// NewAPIError creates a new APIError
func NewAPIError(code int, message string, details string) APIError {
	return APIError{
		Code:    code,
		Message: message,
		Details: details,
	}
}
