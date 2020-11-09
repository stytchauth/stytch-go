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
	return fmt.Sprintf("Stytch Error - request ID: %s, http status: %d, "+
		"type: %s, code: %s, message: %s",
		e.RequestID, e.Status, e.ErrorType, e.ErrorMessage, e.ErrorURL)
}
