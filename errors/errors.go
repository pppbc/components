package errors

import (
	"errors"
	"fmt"
)

type Error struct {
	Code     int               `json:"code,omitempty"`
	Message  string            `json:"message,omitempty"`
	MetaData map[string]string `json:"metaData,omitempty"`
}

const (
	// UnknownCode is unknown code for error info.
	UnknownCode = 500
	// UnknownReason is unknown reason for error info.
	UnknownReason = ""
)

// New returns an error object for the code, message.
func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e *Error) Error() string {
	if e.MetaData == nil {
		return fmt.Sprintf("error: code = %d message = %s", e.Code, e.Message)
	}
	return fmt.Sprintf("error: code = %d message = %s metaData = %v", e.Code, e.Message, e.MetaData)
}

// Is matches each error in the chain with the target value.
func (e *Error) Is(err error) bool {
	if se := new(Error); errors.As(err, &se) {
		return se.Message == e.Message
	}
	return false
}

// WithMetadata with an MD formed by the mapping of key, value.
func (e *Error) WithMetadata(md map[string]string) {
	e.MetaData = md
}

// Newf New(code fmt.Sprintf(format, a...))
func Newf(code int, format string, a ...interface{}) *Error {
	return New(code, fmt.Sprintf(format, a...))
}

// Code returns the http code for a error.
// It supports wrapped errors.
func Code(err error) int {
	if err == nil {
		return 200
	}

	if se := FromError(err); se != nil {
		return se.Code
	}
	return UnknownCode
}

// FromError try to convert an error to *Error.
// It supports wrapped errors.
func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if se := new(Error); errors.As(err, &se) {
		return se
	}
	return New(UnknownCode, err.Error())
}
