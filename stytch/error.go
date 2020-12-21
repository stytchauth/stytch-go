package stytch

import (
	"fmt"
)

type Error struct {
	Status       int          `json:"status,omitempty"`
	RequestID    string       `json:"request_id,omitempty"`
	ErrorType    ErrorType    `json:"error_type,omitempty"`
	ErrorMessage ErrorMessage `json:"error_message,omitempty"`
	ErrorURL     ErrorURL     `json:"error_url,omitempty"`
}

type (
	ErrorType    string
	ErrorMessage string
	ErrorURL     string
)

func (e Error) Error() string {
	error := fmt.Sprintf("Stytch Error - request ID: %s, http status: %d, "+
		"type: %s, message: %s",
		e.RequestID, e.Status, e.ErrorType, e.ErrorMessage)
	if e.ErrorURL != "" {
		error = error + fmt.Sprintf(" error_url: %s", e.ErrorURL)
	}
	return error
}

func newInternalServerError(message string) error {
	if message == "" {
		message = "Oops, something seems to have gone wrong"
	}

	return Error{
		Status:       500,
		ErrorType:    "internal_server_error",
		ErrorMessage: ErrorMessage(message),
		ErrorURL:     "https://docs.stytch.com/reference#500-internal-server-error",
	}
}
