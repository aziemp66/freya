package error

import (
	"fmt"
	"net/http"
)

type ClientError struct {
	Code    int
	Message string
}

func (e ClientError) Error() string {
	return fmt.Sprintf("%d\t%s", e.Code, e.Message)
}

func NewInvariantError(msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

func NewNotFoundError(msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func NewForbiddenError(msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusForbidden,
		Message: msg,
	}
}

func NewUnauthorizedError(msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}

func NewInternalServerError(msg string) *ClientError {
	return &ClientError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}
