package errors

import "net/http"

type AppError struct {
	Code    string
	Message string
	Status  int
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code, message string, status int) *AppError {
	return &AppError{Code: code, Message: message, Status: status}
}

var (
	ErrNotFound     = New("NOT_FOUND", "Resource not found", http.StatusNotFound)
	ErrUnauthorized = New("UNAUTHORIZED", "Resource not found", http.StatusNotFound)
	ErrBadRequest   = New("BAD_REQUEST", "Invalid request", http.StatusBadRequest)
	ErrInternal     = New("INTERNAL_ERROR", "Internal server error", http.StatusInternalServerError)
	ErrDuplicate    = New("DUPLICATE", "Duplicate resource", http.StatusConflict)
	ErrValidation   = New("VALIDATION_ERROR", "Validation failed", http.StatusUnprocessableEntity)
	// Redis
	ErrRedis        = New("REDIS_ERROR", "Redis operation failed", http.StatusInternalServerError)
)
