package rules

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"time"

	"github.com/stytchauth/stytch-go/v15/stytch/consumer/fraud"
)

// SetParams: Request type for `Rules.Set`.
type SetParams struct {
	// Action: The action that should be returned by a fingerprint lookup for that fingerprint or ID with a
	// `RULE_MATCH` reason. The following values are valid: `ALLOW`, `BLOCK`, `CHALLENGE`, or `NONE`. If a
	// `NONE` action is specified, it will clear the stored rule.
	Action fraud.RuleAction `json:"action,omitempty"`
	// VisitorID: The visitor ID we want to set a rule for. Only one fingerprint or ID can be specified in the
	// request.
	VisitorID string `json:"visitor_id,omitempty"`
	// BrowserID: The browser ID we want to set a rule for. Only one fingerprint or ID can be specified in the
	// request.
	BrowserID string `json:"browser_id,omitempty"`
	// VisitorFingerprint: The visitor fingerprint we want to set a rule for. Only one fingerprint or ID can be
	// specified in the request.
	VisitorFingerprint string `json:"visitor_fingerprint,omitempty"`
	// BrowserFingerprint: The browser fingerprint we want to set a rule for. Only one fingerprint or ID can be
	// specified in the request.
	BrowserFingerprint string `json:"browser_fingerprint,omitempty"`
	// HardwareFingerprint: The hardware fingerprint we want to set a rule for. Only one fingerprint or ID can
	// be specified in the request.
	HardwareFingerprint string `json:"hardware_fingerprint,omitempty"`
	// NetworkFingerprint: The network fingerprint we want to set a rule for. Only one fingerprint or ID can be
	// specified in the request.
	NetworkFingerprint string `json:"network_fingerprint,omitempty"`
	// ExpiresInMinutes: The number of minutes until this rule expires. If no `expires_in_minutes` is
	// specified, then the rule is kept permanently.
	ExpiresInMinutes int32 `json:"expires_in_minutes,omitempty"`
	// Description: An optional description for the rule.
	Description string `json:"description,omitempty"`
}

// SetResponse: Response type for `Rules.Set`.
type SetResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// Action: The action that will be returned for the specified fingerprint or ID.
	Action fraud.RuleAction `json:"action,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// VisitorID: The cookie stored on the user's device that uniquely identifies them.
	VisitorID string `json:"visitor_id,omitempty"`
	// BrowserID: Combination of VisitorID and NetworkFingerprint to create a clear identifier of a browser.
	BrowserID string `json:"browser_id,omitempty"`
	// VisitorFingerprint: Cookie-less way of identifying a unique user.
	VisitorFingerprint string `json:"visitor_fingerprint,omitempty"`
	// BrowserFingerprint: Combination of signals to identify a browser and its specific version.
	BrowserFingerprint string `json:"browser_fingerprint,omitempty"`
	// HardwareFingerprint: Combinations of signals to identify an operating system and architecture.
	HardwareFingerprint string `json:"hardware_fingerprint,omitempty"`
	// NetworkFingerprint: Combination of signals associated with a specific network commonly known as TLS
	// fingerprinting.
	NetworkFingerprint string `json:"network_fingerprint,omitempty"`
	// ExpiresAt: The timestamp when the rule expires. Values conform to the RFC 3339 standard and are
	// expressed in UTC, e.g. `2021-12-29T12:33:09Z`.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}
