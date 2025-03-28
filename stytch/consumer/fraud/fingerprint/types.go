package fingerprint

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"time"

	"github.com/stytchauth/stytch-go/v16/stytch/consumer/fraud"
)

// LookupParams: Request type for `Fingerprint.Lookup`.
type LookupParams struct {
	// TelemetryID: The telemetry ID associated with the fingerprint getting looked up.
	TelemetryID string `json:"telemetry_id,omitempty"`
	// ExternalMetadata: External identifiers that you wish to associate with the given telemetry ID. You will
	// be able to search for fingerprint results by these identifiers in the DFP analytics dashboard. External
	// metadata fields may not exceed 65 characters. They may only contain alphanumerics and the characters `_`
	// `-` `+` `.` or `@`.
	ExternalMetadata *fraud.Metadata `json:"external_metadata,omitempty"`
}

// LookupResponse: Response type for `Fingerprint.Lookup`.
type LookupResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// TelemetryID: The telemetry ID associated with the fingerprint getting looked up.
	TelemetryID string `json:"telemetry_id,omitempty"`
	// Fingerprints: A Stytch fingerprint consists of the following identifiers:
	Fingerprints fraud.Fingerprints `json:"fingerprints,omitempty"`
	// Verdict: The metadata associated with each fingerprint
	Verdict fraud.Verdict `json:"verdict,omitempty"`
	// ExternalMetadata: External identifiers that you wish to associate with the given telemetry ID. You will
	// be able to search for fingerprint results by these identifiers in the DFP analytics dashboard. External
	// metadata fields may not exceed 65 characters. They may only contain alphanumerics and the characters `_`
	// `-` `+` `.` or `@`.
	ExternalMetadata fraud.Metadata `json:"external_metadata,omitempty"`
	// CreatedAt: The time when the fingerprint was taken. Values conform to the RFC 3339 standard and are
	// expressed in UTC, e.g. `2021-12-29T12:33:09Z`.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ExpiresAt: The timestamp when the fingerprint expires. Values conform to the RFC 3339 standard and are
	// expressed in UTC, e.g. `2021-12-29T12:33:09Z`.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// Properties: Additional information about the user's browser and network.
	Properties *fraud.Properties `json:"properties,omitempty"`
}
