package api

type ErrorResponse struct {
	Error     bool   `json:"error"`
	Retryable bool   `json:"retryable"`
	Message   string `json:"message"`
}

func NewErrorResponse(err error) *ErrorResponse {
	//
	// TODO: future improvements, if we were to run this in production:
	//  - Determine retryable (transient) errors - TCP timeouts, connection errors or SQL transaction isolation faults
	//  - Be careful to not expose sensitive information in the error message. Ideally we should only allow certain
	//    error messages to be visible to the public
	//
	return &ErrorResponse{
		Error:     true,
		Retryable: false,
		Message:   err.Error(),
	}
}
