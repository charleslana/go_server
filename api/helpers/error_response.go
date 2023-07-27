package helpers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type APIError struct {
	Error      bool   `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func ErrorResponse(err error, status ...int) APIError {
	var statusCode int
	if len(status) > 0 {
		statusCode = status[0]
	} else {
		statusCode = http.StatusBadRequest
	}

	var message string
	if ve, ok := err.(validator.ValidationErrors); ok {
		message = FormatValidationErrorMessage(ve).Error()
	} else {
		message = err.Error()
	}
	return APIError{Error: true, Message: message, StatusCode: statusCode}
}

type validationError struct {
	msg string
}

func (e *validationError) Error() string {
	return e.msg
}

func FormatValidationErrorMessage(err validator.ValidationErrors) error {
	var message string
	for _, e := range err {
		message += e.Tag() + " validation failed on the " + e.Field() + " field; "
	}
	return &validationError{msg: message}
}
