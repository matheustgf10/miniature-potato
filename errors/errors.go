package errors

import (
	"net/http"
)

// !estrutura responsável por os erros
type RestErr struct {
	Message string   `json:"message,omitempty"`
	Err     string   `json:"error,omitempty"`
	Code    int64    `json:"code,omitempty"`
	Causes  []Causes `json:"causes,omitempty"`
}

// !estrutura responsável pelo campo Causes
type Causes struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

// !método que implementa a interface "error" retornando uma mensagem de erro, fazendo com que a
// !estrutura RestErr seja tratada como erro
func (r *RestErr) Error() string {
	return r.Message
}

// !função que cria uma instancia da estrutura RestErr
func NewRestErr(message, err string, code int64, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// ! função que cria um erro expecífico
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad_Request",
		Code:    http.StatusBadRequest,
	}
}

// ! função que cria um erro expecífico
func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad_Request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

// ! função que cria um erro expecífico
func InternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Internal Server error",
		Code:    http.StatusInternalServerError,
	}
}

// ! função que cria um erro expecífico
func NotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Not Found",
		Code:    http.StatusNotFound,
	}
}

// ! função que cria um erro expecífico
func UnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

//! função que cria um erro expecífico

func ForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
	}
}

// ! função que cria um erro expecífico
func ServiceUnavailableError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Service Unavailable",
		Code:    http.StatusServiceUnavailable,
	}
}

// ! função que cria um erro expecífico
func GatewayTimeoutError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Gateway Timeout",
		Code:    http.StatusGatewayTimeout,
	}
}

// ! função que cria um erro expecífico
func TooManyRequests(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Too Many Requests",
		Code:    http.StatusTooManyRequests,
	}
}
