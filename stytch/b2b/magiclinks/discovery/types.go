package discovery

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/discovery"
)

// AuthenticateParams: Request type for `Discovery.Authenticate`.
type AuthenticateParams struct {
	// DiscoveryMagicLinksToken: The Discovery Email Magic Link token to authenticate.
	DiscoveryMagicLinksToken string `json:"discovery_magic_links_token,omitempty"`
	// PkceCodeVerifier: A base64url encoded one time secret used to validate that the request starts and ends
	// on the same device.
	PkceCodeVerifier string `json:"pkce_code_verifier,omitempty"`
}

// AuthenticateResponse: Response type for `Discovery.Authenticate`.
type AuthenticateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// IntermediateSessionToken: The Intermediate Session Token. This token does not belong to a specific
	// instance of a member, but may be exchanged for an existing Member Session or used to create a new
	// organization.
	IntermediateSessionToken string `json:"intermediate_session_token,omitempty"`
	// EmailAddress: The email address.
	EmailAddress string `json:"email_address,omitempty"`
	// DiscoveredOrganizations: An array of `discovered_organization` objects tied to the
	// `intermediate_session_token`, `session_token`, or `session_jwt`. See the
	// [Discovered Organization Object](https://stytch.com/docs/b2b/api/discovered-organization-object) for
	// complete details.
	//
	//   Note that Organizations will only appear here under any of the following conditions:
	//   1. The end user is already a Member of the Organization.
	//   2. The end user is invited to the Organization.
	//   3. The end user can join the Organization because:
	//
	//       a) The Organization allows JIT provisioning.
	//
	//       b) The Organizations' allowed domains list contains the Member's email domain.
	//
	//       c) The Organization has at least one other Member with a verified email address with the same
	// domain as the end user (to prevent phishing attacks).
	DiscoveredOrganizations []discovery.DiscoveredOrganization `json:"discovered_organizations,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}
