package discovery

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/discovery"
)

// AuthenticateParams: Request type for `Discovery.Authenticate`.
type AuthenticateParams struct {
	// EmailAddress: The email address of the Member.
	EmailAddress string `json:"email_address,omitempty"`
	// Password: The password to authenticate, reset, or set for the first time. Any UTF8 character is allowed,
	// e.g. spaces, emojis, non-English characters, etc.
	Password string `json:"password,omitempty"`
}

// AuthenticateResponse: Response type for `Discovery.Authenticate`.
type AuthenticateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// EmailAddress: The email address.
	EmailAddress string `json:"email_address,omitempty"`
	// IntermediateSessionToken: The returned Intermediate Session Token contains a password factor associated
	// with the Member. If this value is non-empty, the member must complete an MFA step to finish logging in
	// to the Organization. The token can be used with the
	// [OTP SMS Authenticate endpoint](https://stytch.com/docs/b2b/api/authenticate-otp-sms),
	// [TOTP Authenticate endpoint](https://stytch.com/docs/b2b/api/authenticate-totp), or
	// [Recovery Codes Recover endpoint](https://stytch.com/docs/b2b/api/recovery-codes-recover) to complete an
	// MFA flow and log in to the Organization. The token has a default expiry of 10 minutes. Password factors
	// are not transferable between Organizations, so the intermediate session token is not valid for use with
	// discovery endpoints.
	IntermediateSessionToken string `json:"intermediate_session_token,omitempty"`
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
