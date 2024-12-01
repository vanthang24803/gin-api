package util

import (
	"fmt"
	"net/http"
	"time"
)

type TException struct {
	HttpCode  int    `json:"httpCode"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func (e *TException) Error() string {
	return fmt.Sprintf("API Error - %s:", e.Message)
}

func BadRequestException(message interface{}) *TException {
	var msg string

	switch v := message.(type) {
	case error:
		msg = v.Error()
	case string:
		msg = v
	default:
		msg = "Invalid error message type"
	}

	return &TException{
		HttpCode:  http.StatusBadRequest,
		Message:   msg,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func NotFoundException(message string) *TException {
	return &TException{
		HttpCode:  http.StatusNotFound,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func InternalServerErrorException() *TException {
	return &TException{
		HttpCode:  http.StatusInternalServerError,
		Message:   "Internal Server Error!",
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func UnauthorizedException() *TException {
	return &TException{
		HttpCode:  http.StatusUnauthorized,
		Message:   "Unauthorized!",
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func ForbiddenException() *TException {
	return &TException{
		HttpCode:  http.StatusForbidden,
		Message:   "Forbidden!",
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
