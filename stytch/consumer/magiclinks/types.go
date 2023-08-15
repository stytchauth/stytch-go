package magiclinks

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v11/stytch/consumer/attribute"
	"github.com/stytchauth/stytch-go/v11/stytch/consumer/sessions"
	"github.com/stytchauth/stytch-go/v11/stytch/consumer/users"
)

// AuthenticateParams: Request type for `MagicLinks.Authenticate`.
type AuthenticateParams struct {
	// Token: The token to authenticate.
	Token string `json:"token,omitempty"`
	// Attributes: Provided attributes help with fraud detection.
	Attributes *attribute.Attributes `json:"attributes,omitempty"`
	// Options: Specify optional security settings.
	Options *Options `json:"options,omitempty"`
	// SessionToken: The `session_token` associated with a User's existing Session.
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
	// SessionJWT: The `session_jwt` associated with a User's existing Session.
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

// CreateParams: Request type for `MagicLinks.Create`.
type CreateParams struct {
	// UserID: The unique ID of a specific User.
	UserID string `json:"user_id,omitempty"`
	// ExpirationMinutes: Set the expiration for the Magic Link `token` in minutes. By default, it expires in 1
	// hour. The minimum expiration is 5 minutes and the maximum is 7 days (10080 mins).
	ExpirationMinutes int32 `json:"expiration_minutes,omitempty"`
	// Attributes: Provided attributes help with fraud detection.
	Attributes *attribute.Attributes `json:"attributes,omitempty"`
}

// Options:
type Options struct {
	// IPMatchRequired: Require that the IP address the Magic Link was requested from matches the IP address
	// it's clicked from.
	IPMatchRequired bool `json:"ip_match_required,omitempty"`
	// UserAgentMatchRequired: Require that the user agent the Magic Link was requested from matches the user
	// agent it's clicked from.
	UserAgentMatchRequired bool `json:"user_agent_match_required,omitempty"`
}

// AuthenticateResponse: Response type for `MagicLinks.Authenticate`.
type AuthenticateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// MethodID: The `email_id` or `phone_id` involved in the given authentication.
	MethodID string `json:"method_id,omitempty"`
	// SessionToken: A secret token for a given Stytch Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
	SessionJWT string `json:"session_jwt,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User users.User `json:"user,omitempty"`
	// ResetSessions: Indicates if all other of the User's Sessions need to be reset. You should check this
	// field if you aren't using Stytch's Session product. If you are using Stytch's Session product, we revoke
	// the User's other sessions for you.
	ResetSessions bool `json:"reset_sessions,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// Session: If you initiate a Session, by including `session_duration_minutes` in your authenticate call,
	// you'll receive a full Session object in the response.
	//
	//   See [GET sessions](https://stytch.com/docs/api/session-get) for complete response fields.
	//
	Session *sessions.Session `json:"session,omitempty"`
}

// CreateResponse: Response type for `MagicLinks.Create`.
type CreateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// Token: The Magic Link `token` that you'll include in your contact method of choice, e.g. email or SMS.
	Token string `json:"token,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}
