package errors

import (
	"net/http"
	"reflect"
)

type Errors struct {
	Type    string
	Code    int64
	Message string
}

const (
	notFoundType            = "HTTPStatusNotFound"
	internalServerErrorType = "HTTPStatusInternalServerError"
	badRequestType          = "HTTPStatusBadRequest"
)

func (e *Errors) Error() string {
	return e.Message
}

func NewWithCode(code int64, message, errType string) error {
	errors := &Errors{
		Type:    errType,
		Code:    code,
		Message: message,
	}

	return errors
}

func NotFound(entity string) error {
	return NewWithCode(http.StatusNotFound, entity+" not found", notFoundType)
}

func InternalServerError(message string) error {
	return NewWithCode(http.StatusInternalServerError, message, internalServerErrorType)
}

func GetType(err error) string {
	if err == nil {
		return "HTTPStatusOK"
	}

	if reflect.TypeOf(err).String() == "*errors.Errors" {
		return err.(*Errors).Type
	}

	return internalServerErrorType
}

func GetCode(err error) int64 {
	if err == nil {
		return 200
	}

	if reflect.TypeOf(err).String() == "*errors.Errors" {
		return err.(*Errors).Code
	}

	return 500
}

func GetMessage(err error) string {
	if err == nil {
		return "OK"
	}

	if reflect.TypeOf(err).String() == "*errors.Errors" {
		return err.(*Errors).Message
	}

	return err.Error()
}
