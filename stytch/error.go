package stytch

import (
	"fmt"
	"strings"
)

type Error struct {
	StatusCode   int          `json:"status_code,omitempty"`
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
	var es strings.Builder
	es.WriteString("Stytch Error - ")

	if e.RequestID != "" {
		fmt.Fprintf(&es, "request ID: %s, ", e.RequestID)
	}

	fmt.Fprintf(&es, "status code: %d, type: %s, message: %s",
		e.StatusCode, e.ErrorType, e.ErrorMessage)

	if e.ErrorURL != "" {
		fmt.Fprintf(&es, " error_url: %s", e.ErrorURL)
	}
	return es.String()
}

func newInternalServerError(message string) error {
	if message == "" {
		message = "Oops, something seems to have gone wrong"
	}

	return Error{
		StatusCode:   500,
		ErrorType:    "internal_server_error",
		ErrorMessage: ErrorMessage(message),
		ErrorURL:     "https://stytch.com/docs/api/errors/500",
	}
}
