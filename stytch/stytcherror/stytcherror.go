package stytcherror

import (
	"fmt"
	"strings"

	"github.com/stytchauth/stytch-go/v11/stytch/config"
)

type Error struct {
	StatusCode   int     `json:"status_code,omitempty"`
	RequestID    string  `json:"request_id,omitempty"`
	ErrorType    Type    `json:"error_type,omitempty"`
	ErrorMessage Message `json:"error_message,omitempty"`
	ErrorURL     URL     `json:"error_url,omitempty"`
}

type (
	Type    string
	Message string
	URL     string
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

type OAuth2Error struct {
	StatusCode   int     `json:"status_code,omitempty"`
	RequestID    string  `json:"request_id,omitempty"`
	ErrorType    Type    `json:"error,omitempty"`
	ErrorMessage Message `json:"error_description,omitempty"`
	ErrorURL     URL     `json:"error_uri,omitempty"`
}

func (e OAuth2Error) Error() string {
	var es strings.Builder
	es.WriteString("Stytch OAuth2 Error - ")

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

func NewClientLibraryError(message string) error {
	if message == "" {
		message = "Oops, something seems to have gone wrong with the Stytch Go client library"
	}

	return Error{
		StatusCode:   500,
		ErrorType:    "client_library_error",
		ErrorMessage: Message(message + ", v" + config.APIVersion),
		ErrorURL:     "https://stytch.com/docs/api/errors/500",
	}
}
