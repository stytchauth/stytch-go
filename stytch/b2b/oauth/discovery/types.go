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
	// DiscoveryOAuthToken: The Discovery OAuth token to authenticate.
	DiscoveryOAuthToken    string         `json:"discovery_oauth_token,omitempty"`
	SessionToken           string         `json:"session_token,omitempty"`
	SessionDurationMinutes int32          `json:"session_duration_minutes,omitempty"`
	SessionJWT             string         `json:"session_jwt,omitempty"`
	SessionCustomClaims    map[string]any `json:"session_custom_claims,omitempty"`
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
	// IntermediateSessionToken: The Intermediate Session Token. This token does not necessarily belong to a
	// specific instance of a Member, but represents a bag of factors that may be converted to a member
	// session. The token can be used with the
	// [OTP SMS Authenticate endpoint](https://stytch.com/docs/b2b/api/authenticate-otp-sms),
	// [TOTP Authenticate endpoint](https://stytch.com/docs/b2b/api/authenticate-totp), or
	// [Recovery Codes Recover endpoint](https://stytch.com/docs/b2b/api/recovery-codes-recover) to complete an
	// MFA flow and log in to the Organization. It can also be used with the
	// [Exchange Intermediate Session endpoint](https://stytch.com/docs/b2b/api/exchange-intermediate-session)
	// to join a specific Organization that allows the factors represented by the intermediate session token;
	// or the
	// [Create Organization via Discovery endpoint](https://stytch.com/docs/b2b/api/create-organization-via-discovery) to create a new Organization and Member.
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
	// ProviderType: Denotes the OAuth identity provider that the user has authenticated with, e.g. Google,
	// Microsoft, GitHub etc.
	ProviderType string `json:"provider_type,omitempty"`
	// ProviderTenantID: The tenant ID returned by the OAuth provider. This is typically used to identify an
	// organization or group within the provider's domain. For example, in HubSpot this is a Hub ID, in Slack
	// this is the Workspace ID, and in GitHub this is an organization ID. This field will only be populated if
	// exactly one tenant ID is returned from a successful OAuth authentication and developers should prefer
	// `provider_tenant_ids` over this since it accounts for the possibility of an OAuth provider yielding
	// multiple tenant IDs.
	ProviderTenantID string `json:"provider_tenant_id,omitempty"`
	// ProviderTenantIds: All tenant IDs returned by the OAuth provider. These is typically used to identify
	// organizations or groups within the provider's domain. For example, in HubSpot this is a Hub ID, in Slack
	// this is the Workspace ID, and in GitHub this is an organization ID. Some OAuth providers do not return
	// tenant IDs, some providers are guaranteed to return one, and some may return multiple. This field will
	// always be populated if at least one tenant ID was returned from the OAuth provider and developers should
	// prefer this field over `provider_tenant_id`.
	ProviderTenantIds []string `json:"provider_tenant_ids,omitempty"`
	// FullName: The full name of the authenticated end user, if available.
	FullName string `json:"full_name,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}
