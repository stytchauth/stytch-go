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

	if e.StatusCode != 0 {
		fmt.Fprintf(&es, "status code: %d, ", e.StatusCode)
	}

	fmt.Fprintf(&es, "type: %s, message: %s", e.ErrorType, e.ErrorMessage)

	if e.ErrorURL != "" {
		fmt.Fprintf(&es, " error_url: %s", e.ErrorURL)
	}
	return es.String()
}

func newClientLibraryError(message string) error {
	if message == "" {
		message = "Oops, something seems to have gone wrong with the Stytch Go client library"
	}

	return Error{
		ErrorType:    "client_library_error",
		ErrorMessage: ErrorMessage(message + ", v" + APIVersion),
		ErrorURL:     "https://stytch.com/docs/api/errors",
	}
}
