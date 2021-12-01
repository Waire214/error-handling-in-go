package main

import (
	"time"

	"github.com/google/uuid"
)

type CustomClientError struct {
	TimeStamp      string         `json:"timestamp"`
	ErrorReference string         `json:"error_reference"`
	Errors         []CustomErrors `json:"errors"`
}
type CustomErrors struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Source  string `json:"source"`
	Err     error  `json:"-"`
}

func (err CustomClientError) Error() string {
	var cE CustomErrors
	return cE.Err.Error()
}

func ErrMessageClient(cES []CustomErrors) error {
	return CustomClientError{
		TimeStamp:      time.Now().Format(time.RFC3339),
		ErrorReference: uuid.New().String(),
		Errors:         cES,
	}
}

func ErrArray(cES []CustomErrors, code int, message, source string, err error) []CustomErrors{

	cES = append(cES, CustomErrors{Code: code, Message: message, Source: source, Err: err})

	return cES
}