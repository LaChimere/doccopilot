package error

import (
	"fmt"
	"log"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`
	message string   `json:"message"`
	details []string `json:"details"`
}

var errorCodes = map[int]string{}

func NewError(code int, message string) *Error {
	if _, ok := errorCodes[code]; ok {
		log.Panicf("error code %d already exists", code)
	}

	errorCodes[code] = message
	return &Error{
		code:    code,
		message: message,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code %d: %s", e.code, e.message)
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = append(newError.details, details...)
	return &newError
}

func (e *Error) StatusCode() int {
	switch e.code {
	case Success.code:
		return http.StatusOK
	case ServerError.code:
		return http.StatusInternalServerError
	case InvalidParams.code:
		return http.StatusBadRequest
	case Unauthorized.code:
		return http.StatusUnauthorized
	case TooManyRequests.code:
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
