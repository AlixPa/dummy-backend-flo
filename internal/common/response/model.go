package response

// MessageResponse represents a simple message response
type Message struct {
	Message string `json:"message"`
}

// SuccessResponse represents the standard success response format
type Data struct {
	Data any `json:"data"`
}
