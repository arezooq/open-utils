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
	// common
	ErrNotFound        = New("NOT_FOUND", "Resource not found", http.StatusNotFound)
	ErrUnauthorized    = New("UNAUTHORIZED", "Resource not found", http.StatusNotFound)
	ErrBadRequest      = New("BAD_REQUEST", "Invalid request", http.StatusBadRequest)
	ErrInternal        = New("INTERNAL_ERROR", "Internal server error", http.StatusInternalServerError)
	ErrDuplicate       = New("DUPLICATE", "Duplicate resource", http.StatusConflict)
	ErrValidation      = New("VALIDATION_ERROR", "Validation failed", http.StatusUnprocessableEntity)
	ErrInvalidateInput = New("INVALID_INPUT", "Input validation error", http.StatusBadRequest)
	ErrMissingToken    = New("MISSING_TOKEN", "Authorization token is missing", http.StatusUnauthorized)
	ErrInvalidToken    = New("INVALID_TOKEN", "Authorization token is invalid", http.StatusUnauthorized)

	// dependency
	ErrRedisConnect   = New("REDIS_CONNECT_ERROR", "Failed to connect to Redis", http.StatusInternalServerError)
	ErrRedisOperation = New("REDIS_OPERATION_ERROR", "Redis operation failed", http.StatusInternalServerError)

	ErrNATSConnect   = New("NATS_CONNECT_ERROR", "Failed to connect to NATS", http.StatusInternalServerError)
	ErrNATSJetStream = New("NATS_JETSTREAM_ERROR", "Failed to init JetStream", http.StatusInternalServerError)

	ErrDBConnect = New("DB_CONNECT_ERROR", "Failed to connect to database", http.StatusInternalServerError)
	ErrDBQuery   = New("DB_QUERY_ERROR", "Database query failed", http.StatusInternalServerError)
)
