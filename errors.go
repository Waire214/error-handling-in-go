package main

import (
	"time"

	"github.com/google/uuid"
)

type CustomClientErrorBody struct {
	TimeStamp      string         `json:"timestamp"`
	ErrorReference string         `json:"error_reference"`
	Errors         []CustomError `json:"errors"`
}
type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Source  string `json:"source"`
	Err     error  `json:"-"`
}

//create a custom error
func (err CustomClientErrorBody) Error() string {
	var customError CustomError
	return customError.Err.Error()
}

//returns error struct as error
func ErrMessageClient(customErrorArray []CustomError) error {
	return CustomClientErrorBody{
		TimeStamp:      time.Now().Format(time.RFC3339),
		ErrorReference: uuid.New().String(),
		Errors:         customErrorArray,
	}
}

//returns an array of error
func ReturnErrorArray(customErrorArray []CustomError, code int, message, source string, err error) []CustomError{

	customErrorArray = append(customErrorArray, CustomError{Code: code, Message: message, Source: source, Err: err})

	return customErrorArray
}