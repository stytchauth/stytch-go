package oauth

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"time"

	"github.com/stytchauth/stytch-go/v16/stytch/consumer/sessions"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/users"
)

// AttachParams: Request type for `OAuth.Attach`.
type AttachParams struct {
	// Provider: The OAuth provider's name.
	Provider string `json:"provider,omitempty"`
	// UserID: The unique ID of a specific User. You may use an `external_id` here if one is set for the user.
	UserID string `json:"user_id,omitempty"`
	// SessionToken: The `session_token` associated with a User's existing Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The `session_jwt` associated with a User's existing Session.
	SessionJWT string `json:"session_jwt,omitempty"`
}

// AuthenticateParams: Request type for `OAuth.Authenticate`.
type AuthenticateParams struct {
	// Token: The OAuth `token` from the `?token=` query parameter in the URL.
	//
	//       The redirect URL will look like
	// `https://example.com/authenticate?stytch_token_type=oauth&token=rM_kw42CWBhsHLF62V75jELMbvJ87njMe3tFVj7Qupu7`
	//
	//       In the redirect URL, the `stytch_token_type` will be `oauth`. See
	// [here](https://stytch.com/docs/workspace-management/redirect-urls) for more detail.
	Token string `json:"token,omitempty"`
	// SessionToken: Reuse an existing session instead of creating a new one. If you provide us with a
	// `session_token`, then we'll update the session represented by this session token with this OAuth factor.
	// If this `session_token` belongs to a different user than the OAuth token, the session_jwt will be
	// ignored. This endpoint will error if both `session_token` and `session_jwt` are provided.
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
	//   If the `session_duration_minutes` parameter is not specified, a Stytch session will not be created.
	SessionDurationMinutes int32 `json:"session_duration_minutes,omitempty"`
	// SessionJWT: Reuse an existing session instead of creating a new one. If you provide us with a
	// `session_jwt`, then we'll update the session represented by this JWT with this OAuth factor. If this
	// `session_jwt` belongs to a different user than the OAuth token, the session_jwt will be ignored. This
	// endpoint will error if both `session_token` and `session_jwt` are provided.
	SessionJWT string `json:"session_jwt,omitempty"`
	// SessionCustomClaims: Add a custom claims map to the Session being authenticated. Claims are only created
	// if a Session is initialized by providing a value in `session_duration_minutes`. Claims will be included
	// on the Session object and in the JWT. To update a key in an existing Session, supply a new value. To
	// delete a key, supply a null value.
	//
	//   Custom claims made with reserved claims ("iss", "sub", "aud", "exp", "nbf", "iat", "jti") will be
	// ignored. Total custom claims size cannot exceed four kilobytes.
	SessionCustomClaims map[string]any `json:"session_custom_claims,omitempty"`
	// CodeVerifier: A base64url encoded one time secret used to validate that the request starts and ends on
	// the same device.
	CodeVerifier string `json:"code_verifier,omitempty"`
}

// ProviderValues:
type ProviderValues struct {
	// AccessToken: The `access_token` that you may use to access the User's data in the provider's API.
	AccessToken string `json:"access_token,omitempty"`
	// RefreshToken: The `refresh_token` that you may use to obtain a new `access_token` for the User within
	// the provider's API.
	RefreshToken string `json:"refresh_token,omitempty"`
	// IDToken: The `id_token` returned by the OAuth provider. ID Tokens are JWTs that contain structured
	// information about a user. The exact content of each ID Token varies from provider to provider. ID Tokens
	// are returned from OAuth providers that conform to the [OpenID Connect](https://openid.net/foundation/)
	// specification, which is based on OAuth.
	IDToken string `json:"id_token,omitempty"`
	// Scopes: The OAuth scopes included for a given provider. See each provider's section above to see which
	// scopes are included by default and how to add custom scopes.
	Scopes []string `json:"scopes,omitempty"`
	// ExpiresAt: The timestamp when the Session expires. Values conform to the RFC 3339 standard and are
	// expressed in UTC, e.g. `2021-12-29T12:33:09Z`.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// AttachResponse: Response type for `OAuth.Attach`.
type AttachResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// OAuthAttachToken: A single-use token for connecting the Stytch User selection from an OAuth Attach
	// request to the corresponding OAuth Start request.
	OAuthAttachToken string `json:"oauth_attach_token,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// AuthenticateResponse: Response type for `OAuth.Authenticate`.
type AuthenticateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// ProviderSubject: The unique identifier for the User within a given OAuth provider. Also commonly called
	// the "sub" or "Subject field" in OAuth protocols.
	ProviderSubject string `json:"provider_subject,omitempty"`
	// ProviderType: Denotes the OAuth identity provider that the user has authenticated with, e.g. Google,
	// Facebook, GitHub etc.
	ProviderType string `json:"provider_type,omitempty"`
	// SessionToken: A secret token for a given Stytch Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
	SessionJWT string `json:"session_jwt,omitempty"`
	// ProviderValues: The `provider_values` object lists relevant identifiers, values, and scopes for a given
	// OAuth provider. For example this object will include a provider's `access_token` that you can use to
	// access the provider's API for a given user.
	//
	//   Note that these values will vary based on the OAuth provider in question, e.g. `id_token` is only
	// returned by OIDC compliant identity providers.
	ProviderValues ProviderValues `json:"provider_values,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User users.User `json:"user,omitempty"`
	// ResetSessions: Indicates if all other of the User's Sessions need to be reset. You should check this
	// field if you aren't using Stytch's Session product. If you are using Stytch's Session product, we revoke
	// the User's other sessions for you.
	ResetSessions bool `json:"reset_sessions,omitempty"`
	// OAuthUserRegistrationID: The unique ID for an OAuth registration.
	OAuthUserRegistrationID string `json:"oauth_user_registration_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// UserSession: A `Session` object. For backwards compatibility reasons, the session from an OAuth
	// authenticate call is labeled as `user_session`, but is otherwise just a standard stytch `Session` object.
	//
	//   See [Session object](https://stytch.com/docs/api/session-object) for complete response fields.
	//
	UserSession *sessions.Session `json:"user_session,omitempty"`
}
