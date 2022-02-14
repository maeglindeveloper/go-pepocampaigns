package pepo

// BaseResponse defines the basic REST API response
type BaseResponse struct {
	Error        string                 `json:"error,omitempty"`
	ErrorMessage map[string]interface{} `json:"error_message,omitempty"`
	Message      string                 `json:"message,omitempty"`
	// Data         interface{}            `json:"data,omitempty"`
}
