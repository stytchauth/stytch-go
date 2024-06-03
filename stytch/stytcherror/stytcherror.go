package stytcherror

import (
	"errors"
	"fmt"
	"strings"

	"github.com/stytchauth/stytch-go/v15/stytch/config"
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

func NewSessionAuthorizationTenancyError(subjectOrgID string, requestOrgID string) error {
	msg := fmt.Sprintf("Subject organization_id %s does not match authZ request organization_id %s", subjectOrgID, requestOrgID)
	return Error{
		StatusCode:   403,
		ErrorType:    "session_authorization_tenancy_error",
		ErrorMessage: Message(msg + ", v" + config.APIVersion),
		ErrorURL:     "https://stytch.com/docs/api/errors/403",
	}
}

func NewPermissionError() error {
	msg := "The Member is not authorized to perform the requested action on that resource."
	return Error{
		StatusCode:   403,
		ErrorType:    "session_authorization_error",
		ErrorMessage: Message(msg + ", v" + config.APIVersion),
		ErrorURL:     "https://stytch.com/docs/api/errors/403",
	}
}

var ErrJWKSNotInitialized = errors.New("JWKS not initialized")
