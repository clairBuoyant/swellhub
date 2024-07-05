package errors

import (
	"fmt"
	"net/http"
)

// AppError represents an application error with a HTTP status code.
type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("status %d: %v - %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("status %d: %v", e.Code, e.Message)
}

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError creates a new AppError.
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

var (
	ErrNotFound       = NewAppError(http.StatusNotFound, "resource not found", nil)
	ErrInternalServer = NewAppError(http.StatusInternalServerError, "internal server error", nil)
)
