package sso

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"time"

	"github.com/stytchauth/stytch-go/v9/stytch/b2b/organizations"
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/sessions"
)

// AuthenticateParams: Request type for `SSO.Authenticate`.
type AuthenticateParams struct {
	// SSOToken: The token to authenticate.
	SSOToken string `json:"sso_token,omitempty"`
	// PkceCodeVerifier: A base64url encoded one time secret used to validate that the request starts and ends
	// on the same device.
	PkceCodeVerifier string `json:"pkce_code_verifier,omitempty"`
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
}

// DeleteConnectionParams: Request type for `SSO.DeleteConnection`.
type DeleteConnectionParams struct {
	// OrganizationID: The organization ID that the SSO connection belongs to.
	OrganizationID string `json:"organization_id,omitempty"`
	// ConnectionID: The ID of the SSO connection. Both SAML and OIDC connection IDs can be provided.
	ConnectionID string `json:"connection_id,omitempty"`
}

// GetConnectionsParams: Request type for `SSO.GetConnections`.
type GetConnectionsParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
}
type OIDCConnection struct {
	OrganizationID   string `json:"organization_id,omitempty"`
	ConnectionID     string `json:"connection_id,omitempty"`
	Status           string `json:"status,omitempty"`
	DisplayName      string `json:"display_name,omitempty"`
	RedirectURL      string `json:"redirect_url,omitempty"`
	ClientID         string `json:"client_id,omitempty"`
	ClientSecret     string `json:"client_secret,omitempty"`
	Issuer           string `json:"issuer,omitempty"`
	AuthorizationURL string `json:"authorization_url,omitempty"`
	TokenURL         string `json:"token_url,omitempty"`
	UserinfoURL      string `json:"userinfo_url,omitempty"`
	JWKSURL          string `json:"jwks_url,omitempty"`
}
type SAMLConnection struct {
	OrganizationID           string            `json:"organization_id,omitempty"`
	ConnectionID             string            `json:"connection_id,omitempty"`
	Status                   string            `json:"status,omitempty"`
	IdpEntityID              string            `json:"idp_entity_id,omitempty"`
	DisplayName              string            `json:"display_name,omitempty"`
	IdpSSOURL                string            `json:"idp_sso_url,omitempty"`
	AcsURL                   string            `json:"acs_url,omitempty"`
	AudienceURI              string            `json:"audience_uri,omitempty"`
	SigningCertificates      []X509Certificate `json:"signing_certificates,omitempty"`
	VerificationCertificates []X509Certificate `json:"verification_certificates,omitempty"`
	AttributeMapping         map[string]any    `json:"attribute_mapping,omitempty"`
}
type X509Certificate struct {
	CertificateID string     `json:"certificate_id,omitempty"`
	Certificate   string     `json:"certificate,omitempty"`
	Issuer        string     `json:"issuer,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	ExpiresAt     *time.Time `json:"expires_at,omitempty"`
}

// AuthenticateResponse: Response type for `SSO.Authenticate`.
type AuthenticateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object).
	Member organizations.Member `json:"member,omitempty"`
	// SessionToken: A secret token for a given Stytch Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
	SessionJWT string `json:"session_jwt,omitempty"`
	// ResetSession: Indicates if all Sessions linked to the Member need to be reset. You should check this
	// field if you aren't using
	//     Stytch's Session product. If you are using Stytch's Session product, we revoke the Member’s other
	// Sessions for you.
	ResetSession bool `json:"reset_session,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// MemberSession: The [Session object](https://stytch.com/docs/b2b/api/session-object).
	MemberSession sessions.MemberSession `json:"member_session,omitempty"`
}

// DeleteConnectionResponse: Response type for `SSO.DeleteConnection`.
type DeleteConnectionResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// ConnectionID: The `connection_id` that was deleted as part of the delete request.
	ConnectionID string `json:"connection_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// GetConnectionsResponse: Response type for `SSO.GetConnections`.
type GetConnectionsResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// SAMLConnections: The list of [SAML Connections](https://stytch.com/docs/b2b/api/saml-connection-object)
	// owned by this organization.
	SAMLConnections []SAMLConnection `json:"saml_connections,omitempty"`
	// OIDCConnections: The list of [OIDC Connections](https://stytch.com/docs/b2b/api/oidc-connection-object)
	// owned by this organization.
	OIDCConnections []OIDCConnection `json:"oidc_connections,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}
