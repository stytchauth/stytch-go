package secrets

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v12/stytch/consumer/m2m"
)

// RotateCancelParams: Request type for `Secrets.RotateCancel`.
type RotateCancelParams struct {
	// ClientID: The ID of the client.
	ClientID string `json:"client_id,omitempty"`
}

// RotateParams: Request type for `Secrets.Rotate`.
type RotateParams struct {
	// ClientID: The ID of the client.
	ClientID string `json:"client_id,omitempty"`
}

// RotateStartParams: Request type for `Secrets.RotateStart`.
type RotateStartParams struct {
	// ClientID: The ID of the client.
	ClientID string `json:"client_id,omitempty"`
}

// RotateCancelResponse: Response type for `Secrets.RotateCancel`.
type RotateCancelResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// M2MClient: The M2M Client affected by this operation.
	M2MClient m2m.M2MClient `json:"m2m_client,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// RotateResponse: Response type for `Secrets.Rotate`.
type RotateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// M2MClient: The M2M Client affected by this operation.
	M2MClient m2m.M2MClient `json:"m2m_client,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// RotateStartResponse: Response type for `Secrets.RotateStart`.
type RotateStartResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// M2MClient: The M2M Client affected by this operation.
	M2MClient m2m.M2MClientWithNextClientSecret `json:"m2m_client,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}
