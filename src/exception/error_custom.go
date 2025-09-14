package exception

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Errors  map[string]string `json:"errors:omitempty"`
	Message string            `json:"message"`
	Code    string            `json:"code"`
	Status  int               `json:"-"`
}

func (err *ErrorResponse) Error() string {
	return err.Message
}

func (err *ErrorResponse) GetStatusHttp() int {
	if err.Status == 0 {
		return http.StatusInternalServerError
	}

	return err.Status
}

// Predefined error codes
const (
	// Client errors
	ErrCodeBadRequest   = "BAD_REQUEST"
	ErrCodeValidation   = "VALIDATION_ERROR"
	ErrCodeUnauthorized = "UNAUTHORIZED"
	ErrCodeForbidden    = "FORBIDDEN"
	ErrCodeNotFound     = "NOT_FOUND"
	ErrCodeConflict     = "CONFLICT"

	// Server errors
	ErrCodeInternal = "INTERNAL_SERVER_ERROR"
	ErrCodeDatabase = "DATABASE_ERROR"
	ErrCodeExternal = "EXTERNAL_SERVICE_ERROR"
)

func NewBadReqeust(message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    ErrCodeBadRequest,
		Message: message,
		Status:  http.StatusBadRequest,
	}
}
func NewValidationError(message string, fieldErrors map[string]string) *ErrorResponse {
	return &ErrorResponse{
		Code:    ErrCodeValidation,
		Message: message,
		Status:  http.StatusBadRequest,
		Errors:  fieldErrors,
	}
}
func NewNotFoundData(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    ErrCodeNotFound,
		Status:  http.StatusNotFound,
	}
}

func NewDbError(message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    ErrCodeDatabase,
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

func NewInternalServerError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    ErrCodeInternal,
		Status:  http.StatusInternalServerError,
	}
}

func IsCustomError(err error) (*ErrorResponse, bool) {
	customErr, ok := err.(*ErrorResponse)
	return customErr, ok
}

func WrapError(err error, message string, code string, status int) *ErrorResponse {
	return &ErrorResponse{
		Message: fmt.Sprintf("%s: %v", message, err),
		Code:    code,
		Status:  status,
	}
}
