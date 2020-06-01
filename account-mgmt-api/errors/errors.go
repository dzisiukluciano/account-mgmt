package errors

import "net/http"

// APIError holds an operation's error detail
type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// BadRequest creates an APIError for constraints not satisfied
func BadRequest(message string) *APIError {
	return &APIError{
		Status:  http.StatusBadRequest,
		Message: message,
		Code:    "bad_request",
	}
}

// NotFound creates an not-found APIError
func NotFound() *APIError {
	return &APIError{
		Status:  http.StatusNotFound,
		Message: "Transaction not found",
		Code:    "not_found",
	}
}

// InternalServer creates an APIError with custom message
func InternalServer(message string) *APIError {
	return &APIError{
		Status:  http.StatusInternalServerError,
		Message: message,
		Code:    "internal_server_error",
	}
}
