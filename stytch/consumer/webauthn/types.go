package webauthn

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v8/stytch/consumer/sessions"
	"github.com/stytchauth/stytch-go/v8/stytch/consumer/users"
)

// AuthenticateParams: Request type for `WebAuthn.Authenticate`.
type AuthenticateParams struct {
	// PublicKeyCredential: The response of the
	// [navigator.credentials.create()](https://www.w3.org/TR/webauthn-2/#sctn-createCredential).
	PublicKeyCredential string `json:"public_key_credential,omitempty"`
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
}

// AuthenticateStartParams: Request type for `WebAuthn.AuthenticateStart`.
type AuthenticateStartParams struct {
	// UserID: The `user_id` of an active user the WebAuthn registration should be tied to.
	UserID string `json:"user_id,omitempty"`
	// Domain: The domain for WebAuthn. Defaults to `window.location.hostname`.
	Domain string `json:"domain,omitempty"`
}

// RegisterParams: Request type for `WebAuthn.Register`.
type RegisterParams struct {
	// UserID: The `user_id` of an active user the WebAuthn registration should be tied to.
	UserID string `json:"user_id,omitempty"`
	// PublicKeyCredential: The response of the
	// [navigator.credentials.create()](https://www.w3.org/TR/webauthn-2/#sctn-createCredential).
	PublicKeyCredential string `json:"public_key_credential,omitempty"`
}

// RegisterStartParams: Request type for `WebAuthn.RegisterStart`.
type RegisterStartParams struct {
	// UserID: The `user_id` of an active user the WebAuthn registration should be tied to.
	UserID string `json:"user_id,omitempty"`
	// Domain: The domain for WebAuthn. Defaults to `window.location.hostname`.
	Domain string `json:"domain,omitempty"`
	// UserAgent: The user agent of the User.
	UserAgent string `json:"user_agent,omitempty"`
	// AuthenticatorType: The requested authenticator type of the WebAuthn device. The two valid value are
	// platform and cross-platform. If no value passed, we assume both values are allowed.
	AuthenticatorType string `json:"authenticator_type,omitempty"`
}

// AuthenticateResponse: Response type for `WebAuthn.Authenticate`.
type AuthenticateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// WebauthnRegistrationID: The unique ID for the WebAuthn registration.
	WebauthnRegistrationID string `json:"webauthn_registration_id,omitempty"`
	// SessionToken: A secret token for a given Stytch Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
	SessionJWT string `json:"session_jwt,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User users.User `json:"user,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// Session: If you initiate a Session, by including `session_duration_minutes` in your authenticate call,
	// you'll receive a full Session object in the response.
	//
	//   See [GET sessions](https://stytch.com/docs/api/session-get) for complete response fields.
	//
	Session sessions.Session `json:"session,omitempty"`
}

// AuthenticateStartResponse: Response type for `WebAuthn.AuthenticateStart`.
type AuthenticateStartResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// PublicKeyCredentialRequestOptions: Options used for WebAuthn authentication.
	PublicKeyCredentialRequestOptions string `json:"public_key_credential_request_options,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// RegisterResponse: Response type for `WebAuthn.Register`.
type RegisterResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// WebauthnRegistrationID: The unique ID for the WebAuthn registration.
	WebauthnRegistrationID string `json:"webauthn_registration_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// RegisterStartResponse: Response type for `WebAuthn.RegisterStart`.
type RegisterStartResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// PublicKeyCredentialCreationOptions: Options used for WebAuthn registration.
	PublicKeyCredentialCreationOptions string `json:"public_key_credential_creation_options,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}
