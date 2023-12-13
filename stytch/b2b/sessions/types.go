package sessions

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stytchauth/stytch-go/v11/stytch/b2b/mfa"
	"github.com/stytchauth/stytch-go/v11/stytch/b2b/organizations"
	"github.com/stytchauth/stytch-go/v11/stytch/consumer/sessions"
)

// AuthenticateParams: Request type for `Sessions.Authenticate`.
type AuthenticateParams struct {
	// SessionToken: A secret token for a given Stytch Session.
	SessionToken string `json:"session_token,omitempty"`
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
	// SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
	SessionJWT string `json:"session_jwt,omitempty"`
	// SessionCustomClaims: Add a custom claims map to the Session being authenticated. Claims are only created
	// if a Session is initialized by providing a value in
	//   `session_duration_minutes`. Claims will be included on the Session object and in the JWT. To update a
	// key in an existing Session, supply a new value. To
	//   delete a key, supply a null value. Custom claims made with reserved claims (`iss`, `sub`, `aud`,
	// `exp`, `nbf`, `iat`, `jti`) will be ignored.
	//   Total custom claims size cannot exceed four kilobytes.
	SessionCustomClaims map[string]any `json:"session_custom_claims,omitempty"`
	// AuthorizationCheck: (Coming Soon) If an `authorization_check` object is passed in, this endpoint will
	// also check if the Member is
	//   authorized to perform the given action on the given Resource in the specified Organization. A Member
	// is authorized if
	//   their Member Session contains a Role, assigned
	//   [explicitly or implicitly](https://github.com/docs/b2b/guides/rbac/role-assignment), with adequate
	// permissions.
	//   In addition, the `organization_id` passed in the authorization check must match the Member's
	// Organization.
	//
	//   The Roles on the Member Session may differ from the Roles you see on the Member object - Roles that
	// are implicitly
	//   assigned by SSO connection or SSO group will only be valid for a Member Session if there is at least
	// one authentication
	//   factor on the Member Session from the specified SSO connection.
	//
	//   If the Member is not authorized to perform the specified action on the specified Resource, or if the
	//   `organization_id` does not match the Member's Organization, a 403 error will be thrown.
	//   Otherwise, the response will contain a list of Roles that satisfied the authorization check.
	AuthorizationCheck *AuthorizationCheck `json:"authorization_check,omitempty"`
}

// AuthorizationCheck:
type AuthorizationCheck struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// ResourceID: A unique identifier of the RBAC Resource, provided by the developer and intended to be
	// human-readable.
	//
	//   A `resource_id` is not allowed to start with `stytch`, which is a special prefix used for Stytch
	// default Resources with reserved  `resource_id`s. These include:
	//
	//   * `stytch.organization`
	//   * `stytch.member`
	//   * `stytch.sso`
	//   * `stytch.self`
	//
	//   Check out the
	// [guide on Stytch default Resources](https://stytch.com/docs/b2b/guides/rbac/stytch-defaults) for a more
	// detailed explanation.
	//
	//
	ResourceID string `json:"resource_id,omitempty"`
	// Action: An action to take on a Resource.
	Action string `json:"action,omitempty"`
}

type AuthorizationVerdict struct {
	Authorized    bool     `json:"authorized,omitempty"`
	GrantingRoles []string `json:"granting_roles,omitempty"`
}

// ExchangeParams: Request type for `Sessions.Exchange`.
type ExchangeParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// SessionToken: The `session_token` belonging to the member that you wish to associate the email with.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The `session_jwt` belonging to the member that you wish to associate the email with.
	SessionJWT string `json:"session_jwt,omitempty"`
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
	// Locale: If the Member needs to complete an MFA step, and the Member has a phone number, this endpoint
	// will pre-emptively send a one-time passcode (OTP) to the Member's phone number. The locale argument will
	// be used to determine which language to use when sending the passcode.
	//
	// Parameter is a [IETF BCP 47 language tag](https://www.w3.org/International/articles/language-tags/),
	// e.g. `"en"`.
	//
	// Currently supported languages are English (`"en"`), Spanish (`"es"`), and Brazilian Portuguese
	// (`"pt-br"`); if no value is provided, the copy defaults to English.
	//
	// Request support for additional languages
	// [here](https://docs.google.com/forms/d/e/1FAIpQLScZSpAu_m2AmLXRT3F3kap-s_mcV6UTBitYn6CdyWP0-o7YjQ/viewform?usp=sf_link")!
	//
	Locale ExchangeRequestLocale `json:"locale,omitempty"`
}

// GetJWKSParams: Request type for `Sessions.GetJWKS`.
type GetJWKSParams struct {
	// ProjectID: The `project_id` to get the JWKS for.
	ProjectID string `json:"project_id,omitempty"`
}

// GetParams: Request type for `Sessions.Get`.
type GetParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member. The `member_id` is critical to perform
	// operations on a Member, so be sure to preserve this value.
	MemberID string `json:"member_id,omitempty"`
}

// MemberSession:
type MemberSession struct {
	// MemberSessionID: Globally unique UUID that identifies a specific Session.
	MemberSessionID string `json:"member_session_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// StartedAt: The timestamp when the Session was created. Values conform to the RFC 3339 standard and are
	// expressed in UTC, e.g. `2021-12-29T12:33:09Z`.
	StartedAt *time.Time `json:"started_at,omitempty"`
	// LastAccessedAt: The timestamp when the Session was last accessed. Values conform to the RFC 3339
	// standard and are expressed in UTC, e.g. `2021-12-29T12:33:09Z`.
	LastAccessedAt *time.Time `json:"last_accessed_at,omitempty"`
	// ExpiresAt: The timestamp when the Session expires. Values conform to the RFC 3339 standard and are
	// expressed in UTC, e.g. `2021-12-29T12:33:09Z`.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// AuthenticationFactors: An array of different authentication factors that comprise a Session.
	AuthenticationFactors []sessions.AuthenticationFactor `json:"authentication_factors,omitempty"`
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string   `json:"organization_id,omitempty"`
	Roles          []string `json:"roles,omitempty"`
	// CustomClaims: The custom claims map for a Session. Claims can be added to a session during a Sessions
	// authenticate call.
	CustomClaims map[string]any `json:"custom_claims,omitempty"`
}

// RevokeParams: Request type for `Sessions.Revoke`.
type RevokeParams struct {
	// MemberSessionID: Globally unique UUID that identifies a specific Session in the Stytch API. The
	// `member_session_id` is critical to perform operations on an Session, so be sure to preserve this value.
	MemberSessionID string `json:"member_session_id,omitempty"`
	// SessionToken: A secret token for a given Stytch Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
	SessionJWT string `json:"session_jwt,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member. The `member_id` is critical to perform
	// operations on a Member, so be sure to preserve this value.
	MemberID string `json:"member_id,omitempty"`
}

// AuthenticateResponse: Response type for `Sessions.Authenticate`.
type AuthenticateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberSession: The [Session object](https://stytch.com/docs/b2b/api/session-object).
	MemberSession MemberSession `json:"member_session,omitempty"`
	// SessionToken: A secret token for a given Stytch Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
	SessionJWT string `json:"session_jwt,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object)
	Member organizations.Member `json:"member,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// Verdict: (Coming Soon) If an `authorization_check` is provided in the request and the check succeeds,
	// this field will return
	//   the complete list of Roles that gave the Member permission to perform the specified action on the
	// specified Resource.
	Verdict *AuthorizationVerdict `json:"verdict,omitempty"`
}

// ExchangeResponse: Response type for `Sessions.Exchange`.
type ExchangeResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// MemberSession: The [Session object](https://stytch.com/docs/b2b/api/session-object).
	MemberSession MemberSession `json:"member_session,omitempty"`
	// SessionToken: A secret token for a given Stytch Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
	SessionJWT string `json:"session_jwt,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object)
	Member organizations.Member `json:"member,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
	// MemberAuthenticated: Indicates whether the Member is fully authenticated. If false, the Member needs to
	// complete an MFA step to log in to the Organization.
	MemberAuthenticated bool `json:"member_authenticated,omitempty"`
	// IntermediateSessionToken: The returned Intermediate Session Token contains any Email Magic Link or OAuth
	// factors from the original member session that are valid for the target Organization.
	//       The token can be used with the
	// [OTP SMS Authenticate endpoint](https://stytch.com/docs/b2b/api/authenticate-otp-sms) to complete the
	// MFA flow and log in to the target Organization.
	//       It can also be used with the
	// [Exchange Intermediate Session endpoint](https://stytch.com/docs/b2b/api/exchange-intermediate-session)
	// to join a different existing Organization,
	//       or the
	// [Create Organization via Discovery endpoint](https://stytch.com/docs/b2b/api/create-organization-via-discovery) to create a new Organization.
	IntermediateSessionToken string `json:"intermediate_session_token,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// MFARequired: Information about the MFA requirements of the Organization and the Member's options for
	// fulfilling MFA.
	MFARequired *mfa.MfaRequired `json:"mfa_required,omitempty"`
}

// GetJWKSResponse: Response type for `Sessions.GetJWKS`.
type GetJWKSResponse struct {
	// Keys: The JWK
	Keys []sessions.JWK `json:"keys,omitempty"`
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// GetResponse: Response type for `Sessions.Get`.
type GetResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberSessions: An array of [Session objects](https://stytch.com/docs/b2b/api/session-object).
	MemberSessions []MemberSession `json:"member_sessions,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// RevokeResponse: Response type for `Sessions.Revoke`.
type RevokeResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

type ExchangeRequestLocale string

const (
	ExchangeRequestLocaleEn   ExchangeRequestLocale = "en"
	ExchangeRequestLocaleEs   ExchangeRequestLocale = "es"
	ExchangeRequestLocalePtbr ExchangeRequestLocale = "pt-br"
)

// MANUAL(Types)(TYPES)
// ADDIMPORT: "errors"
// ADDIMPORT: "strings"
// ADDIMPORT: "github.com/golang-jwt/jwt/v5"
// ADDIMPORT: "github.com/stytchauth/stytch-go/v11/stytch/consumer/sessions"

var ErrJWTTooOld = errors.New("JWT too old")

// AuthenticateJWTParams: Request type for `Sessions.AuthenticateJWT`.
type AuthenticateJWTParams struct {
	MaxTokenAge time.Duration
	Body        *AuthenticateParams
}

type OrgClaim struct {
	ID   string `json:"organization_id"`
	Slug string `json:"slug"`
}

type Claims struct {
	Session      sessions.SessionClaim `json:"https://stytch.com/session"`
	Organization OrgClaim              `json:"https://stytch.com/organization"`
	jwt.RegisteredClaims
}

type ClaimsWrapper struct {
	Claims map[string]any `json:"custom_claims"`
}

type SessionWrapper struct {
	Session ClaimsWrapper `json:"session"`
}

// ENDMANUAL(Types)
