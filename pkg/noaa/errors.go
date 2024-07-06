package noaa

import (
	"fmt"
	"net/http"
)

// NOAAError represents a standard error message.
type NOAAError struct {
	Message string
	Err     error
}

func (e *NOAAError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("error: %v - %v", e.Message, e.Err)
	}
	return fmt.Sprintf("error: %v", e.Message)
}

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
func (e *NOAAError) Unwrap() error {
	return e.Err
}

// NewError creates a new NOAAError.
func NewError(message string, err error) *NOAAError {
	return &NOAAError{
		Message: message,
		Err:     err,
	}
}

// NOAARequestError represents an application error with a HTTP status code.
type NOAARequestError struct {
	Code    int
	Message string
	Err     error
}

func (e *NOAARequestError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("http error(%d): %v - %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("http error(%d): %v", e.Code, e.Message)
}

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
func (e *NOAARequestError) Unwrap() error {
	return e.Err
}

// NewRequestError creates a new NOAARequestError.
func NewRequestError(code int, message string, err error) *NOAARequestError {
	return &NOAARequestError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

var (
	ReqErrNotFound       = NewRequestError(http.StatusNotFound, "resource not found", nil)
	ReqErrInternalServer = NewRequestError(http.StatusInternalServerError, "internal server error", nil)
)
