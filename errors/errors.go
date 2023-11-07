package errors

import (
	"net/http"
)

type RestErr struct {
	Message string   `json:"message,omitempty"`
	Err     string   `json:"error,omitempty"`
	Code    int64    `json:"code,omitempty"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int64, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}

}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad_Request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad_Request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}

}

func InternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Internal Server error",
		Code:    http.StatusInternalServerError,
	}
}

func NotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Not Found",
		Code:    http.StatusNotFound,
	}
}

func UnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func ForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
	}
}

func ServiceUnavailableError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Service Unavailable",
		Code:    http.StatusServiceUnavailable,
	}
}

func GatewayTimeoutError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Gateway Timeout",
		Code:    http.StatusGatewayTimeout,
	}
}

func TooManyRequests(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Too Many Requests",
		Code:    http.StatusTooManyRequests,
	}
}
