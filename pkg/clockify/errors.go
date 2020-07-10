package clockify

import (
	"fmt"
	"net/http"
)

const (
	StatusBadRequest          = http.StatusBadRequest
	StatusUnauthorized        = http.StatusUnauthorized
	StatusForbidden           = http.StatusForbidden
	StatusNotFound            = http.StatusNotFound
	StatusDecodeError         = 1400
	StatusEncodeError         = 1400
	StatusInternalClientError = 1500
)

// NewBadRequestError with message
func NewBadRequestError(msg string, args ...interface{}) *APIError {
	return NewAPIError(StatusBadRequest, msg, args...)
}

// NewUnauthorizedError with message
func NewUnauthorizedError(msg string, args ...interface{}) *APIError {
	return NewAPIError(StatusUnauthorized, msg, args...)
}

// NewForbiddenError with message
func NewForbiddenError(msg string, args ...interface{}) *APIError {
	return NewAPIError(StatusForbidden, msg, args...)
}

// NewNotFoundError with message
func NewNotFoundError(msg string, args ...interface{}) *APIError {
	return NewAPIError(StatusNotFound, msg, args...)
}

// NewDecodeError with message
func NewDecodeError(msg string, args ...interface{}) *APIError {
	return NewAPIError(StatusDecodeError, msg, args...)
}

// NewEncodeError with message
func NewEncodeError(msg string, args ...interface{}) *APIError {
	return NewAPIError(StatusEncodeError, msg, args...)
}

// NewInternalClientError with message
func NewInternalClientError(msg string, args ...interface{}) *APIError {
	return NewAPIError(StatusInternalClientError, msg, args...)
}

// APIError is an error with a message and clockify errors
type APIError struct {
	Code          int    `json:"code"`
	Message       string `json:"msg"`
	InternalError error  `json:"-"`
	ErrorID       string `json:"error_id,omitempty"`
}

// NewAPIError with a message
func NewAPIError(status int, msg string, args ...interface{}) *APIError {
	return &APIError{
		Code:    status,
		Message: fmt.Sprintf(msg, args...),
	}
}

func (a *APIError) Error() string {
	return fmt.Sprintf("%d: %s", a.Code, a.Message)
}

// WithInternalError add a error to the api error
func (a *APIError) WithInternalError(err error) *APIError {
	a.InternalError = err
	return a
}

// IsUnauthorized code
func (a *APIError) IsUnauthorized() bool {
	return a.Code == StatusUnauthorized
}

// IsForbidden code
func (a *APIError) IsForbidden() bool {
	return a.Code == StatusForbidden
}

// IsNotFound code
func (a *APIError) IsNotFound() bool {
	return a.Code == StatusNotFound
}

// IsBadRequest code
func (a *APIError) IsBadRequest() bool {
	return a.Code == StatusBadRequest
}
