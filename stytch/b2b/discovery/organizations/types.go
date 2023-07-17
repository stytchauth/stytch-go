package organizations

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/discovery"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/organizations"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/sessions"
)

// CreateParams: Request type for `Organizations.Create`.
type CreateParams struct {
	// IntermediateSessionToken: The Intermediate Session Token. This token does not belong to a specific
	// instance of a member, but may be exchanged for an existing Member Session or used to create a new
	// organization.
	IntermediateSessionToken string `json:"intermediate_session_token,omitempty"`
	// OrganizationName: The name of the Organization. If the name is not specified, a default name will be
	// created based on the email used to initiate the discovery flow. If the email domain is a common email
	// provider such as gmail.com, or if the email is a .edu email, the organization name will be generated
	// based on the name portion of the email. Otherwise, the organization name will be generated based on the
	// email domain.
	OrganizationName string `json:"organization_name,omitempty"`
	// OrganizationSlug: The unique URL slug of the Organization. A minimum of two characters is required. The
	// slug only accepts alphanumeric characters and the following reserved characters: `-` `.` `_` `~`. If the
	// slug is not specified, a default slug will be created based on the email used to initiate the discovery
	// flow. If the email domain is a common email provider such as gmail.com, or if the email is a .edu email,
	// the organization slug will be generated based on the name portion of the email. Otherwise, the
	// organization slug will be generated based on the email domain.
	OrganizationSlug string `json:"organization_slug,omitempty"`
	// SessionDurationMinutes: Set the session lifetime to be this many minutes from now. This will start a new
	// session if one doesn't already exist,
	//   returning both an opaque `session_token` and `session_jwt` for this session. Remember that the
	// `session_jwt` will have a fixed lifetime of
	//   five minutes regardless of the underlying session duration, and will need to be refreshed over time.
	//
	//   This value must be a minimum of 5 and a maximum of 527040 minutes (366 days).
	//
	//   If a `session_token` or `session_jwt` is provided then a successful authentication will continue to
	// extend the session this many minutes.
	//
	//   If the `session_duration_minutes` parameter is not specified, a Stytch session will be created with a
	// 60 minute duration. If you don't want
	//   to use the Stytch session product, you can ignore the session fields in the response.
	SessionDurationMinutes int32 `json:"session_duration_minutes,omitempty"`
	// SessionCustomClaims: Add a custom claims map to the Session being authenticated. Claims are only created
	// if a Session is initialized by providing a value in
	//   `session_duration_minutes`. Claims will be included on the Session object and in the JWT. To update a
	// key in an existing Session, supply a new value. To
	//   delete a key, supply a null value. Custom claims made with reserved claims (`iss`, `sub`, `aud`,
	// `exp`, `nbf`, `iat`, `jti`) will be ignored.
	//   Total custom claims size cannot exceed four kilobytes.
	SessionCustomClaims map[string]any `json:"session_custom_claims,omitempty"`
	// OrganizationLogoURL: The image URL of the Organization logo.
	OrganizationLogoURL string `json:"organization_logo_url,omitempty"`
	// TrustedMetadata: An arbitrary JSON object for storing application-specific data or
	// identity-provider-specific data.
	TrustedMetadata map[string]any `json:"trusted_metadata,omitempty"`
	// SsoJitProvisioning: The authentication setting that controls the JIT provisioning of Members when
	// authenticating via SSO. The accepted values are:
	//
	//   `ALL_ALLOWED` – new Members will be automatically provisioned upon successful authentication via any
	// of the Organization's `sso_active_connections`.
	//
	//   `RESTRICTED` – only new Members with SSO logins that comply with
	// `sso_jit_provisioning_allowed_connections` can be provisioned upon authentication.
	//
	//   `NOT_ALLOWED` – disable JIT provisioning via SSO.
	//
	SsoJitProvisioning string `json:"sso_jit_provisioning,omitempty"`
	// EmailAllowedDomains: An array of email domains that allow invites or JIT provisioning for new Members.
	// This list is enforced when either `email_invites` or `email_jit_provisioning` is set to `RESTRICTED`.
	//
	//
	//     Common domains such as `gmail.com` are not allowed. See the
	// [common email domains resource](https://stytch.com/docs/b2b/api/common-email-domains) for the full list.
	EmailAllowedDomains []string `json:"email_allowed_domains,omitempty"`
	// EmailJitProvisioning: The authentication setting that controls how a new Member can be provisioned by
	// authenticating via Email Magic Link. The accepted values are:
	//
	//   `RESTRICTED` – only new Members with verified emails that comply with `email_allowed_domains` can be
	// provisioned upon authentication via Email Magic Link.
	//
	//   `NOT_ALLOWED` – disable JIT provisioning via Email Magic Link.
	//
	EmailJitProvisioning string `json:"email_jit_provisioning,omitempty"`
	// EmailInvites: The authentication setting that controls how a new Member can be invited to an
	// organization by email. The accepted values are:
	//
	//   `ALL_ALLOWED` – any new Member can be invited to join via email.
	//
	//   `RESTRICTED` – only new Members with verified emails that comply with `email_allowed_domains` can be
	// invited via email.
	//
	//   `NOT_ALLOWED` – disable email invites.
	//
	EmailInvites string `json:"email_invites,omitempty"`
	// AuthMethods: The setting that controls which authentication methods can be used by Members of an
	// Organization. The accepted values are:
	//
	//   `ALL_ALLOWED` – the default setting which allows all authentication methods to be used.
	//
	//   `RESTRICTED` – only methods that comply with `allowed_auth_methods` can be used for authentication.
	// This setting does not apply to Members with `is_breakglass` set to `true`.
	//
	AuthMethods string `json:"auth_methods,omitempty"`
	// AllowedAuthMethods:
	//   An array of allowed authentication methods. This list is enforced when `auth_methods` is set to
	// `RESTRICTED`.
	//   The list's accepted values are: `sso`, `magic_link`, `password`, `google_oauth`, and `microsoft_oauth`.
	//
	AllowedAuthMethods []string `json:"allowed_auth_methods,omitempty"`
}

// ListParams: Request type for `Organizations.List`.
type ListParams struct {
	// IntermediateSessionToken: The Intermediate Session Token. This token does not belong to a specific
	// instance of a member, but may be exchanged for an existing Member Session or used to create a new
	// organization.
	IntermediateSessionToken string `json:"intermediate_session_token,omitempty"`
	// SessionToken: A secret token for a given Stytch Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
	SessionJWT string `json:"session_jwt,omitempty"`
}

// CreateResponse: Response type for `Organizations.Create`.
type CreateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// SessionToken: A secret token for a given Stytch Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
	SessionJWT string `json:"session_jwt,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object).
	Member organizations.Member `json:"member,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// MemberSession: The [Session object](https://stytch.com/docs/b2b/api/session-object).
	MemberSession sessions.MemberSession `json:"member_session,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
}

// ListResponse: Response type for `Organizations.List`.
type ListResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
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
