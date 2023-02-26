package utils

import (
	"fmt"
	"net/http"
)

type (
	Error struct {
		RawError error        `json:"error"`
		Message  ErrorMessage `json:"message"`
	}

	ErrorMessage string
)

const (
	NotFound  ErrorMessage = "NotFound"
)

func NewError(rawError error, message ErrorMessage) *Error {
	return &Error{
		RawError: rawError,
		Message: message,
	}
}

func NewNotFoundError(errorMessage string) *Error {
	return &Error{
		RawError: fmt.Errorf(errorMessage),
		Message: NotFound,
	}
}

func ConvertMessageToCode(msg string) int {
	messages := map[string]int{
		string(NotFound):  http.StatusNotFound,
	}

	statusCode, ok := messages[msg]
	if !ok {
		return http.StatusInternalServerError
	}

	return statusCode
}
