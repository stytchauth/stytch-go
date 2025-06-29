package totps

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/sessions"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/users"
)

// AuthenticateParams: Request type for `TOTPs.Authenticate`.
type AuthenticateParams struct {
	// UserID: The `user_id` of an active user the TOTP registration should be tied to. You may use an
	// `external_id` here if one is set for the user.
	UserID string `json:"user_id,omitempty"`
	// TOTPCode: The TOTP code to authenticate. The TOTP code should consist of 6 digits.
	TOTPCode string `json:"totp_code,omitempty"`
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

// CreateParams: Request type for `TOTPs.Create`.
type CreateParams struct {
	// UserID: The `user_id` of an active user the TOTP registration should be tied to. You may use an
	// `external_id` here if one is set for the user.
	UserID string `json:"user_id,omitempty"`
	// ExpirationMinutes: The expiration for the TOTP instance. If the newly created TOTP is not authenticated
	// within this time frame the TOTP will be unusable. Defaults to 1440 (1 day) with a minimum of 5 and a
	// maximum of 1440.
	ExpirationMinutes int32 `json:"expiration_minutes,omitempty"`
}

// RecoverParams: Request type for `TOTPs.Recover`.
type RecoverParams struct {
	// UserID: The `user_id` of an active user the TOTP registration should be tied to. You may use an
	// `external_id` here if one is set for the user.
	UserID string `json:"user_id,omitempty"`
	// RecoveryCode: The recovery code to authenticate.
	RecoveryCode string `json:"recovery_code,omitempty"`
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

// RecoveryCodesParams: Request type for `TOTPs.RecoveryCodes`.
type RecoveryCodesParams struct {
	// UserID: The `user_id` of an active user the TOTP registration should be tied to. You may use an
	// `external_id` here if one is set for the user.
	UserID string `json:"user_id,omitempty"`
}

// TOTP:
type TOTP struct {
	// TOTPID: The unique ID for a TOTP instance.
	TOTPID string `json:"totp_id,omitempty"`
	// Verified: The verified boolean denotes whether or not this send method, e.g. phone number, email
	// address, etc., has been successfully authenticated by the User.
	Verified bool `json:"verified,omitempty"`
	// RecoveryCodes: The recovery codes used to authenticate the user without an authenticator app.
	RecoveryCodes []string `json:"recovery_codes,omitempty"`
}

// AuthenticateResponse: Response type for `TOTPs.Authenticate`.
type AuthenticateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// SessionToken: A secret token for a given Stytch Session.
	SessionToken string `json:"session_token,omitempty"`
	// TOTPID: The unique ID for a TOTP instance.
	TOTPID string `json:"totp_id,omitempty"`
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
	//   See [Session object](https://stytch.com/docs/api/session-object) for complete response fields.
	//
	Session *sessions.Session `json:"session,omitempty"`
}

// CreateResponse: Response type for `TOTPs.Create`.
type CreateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// TOTPID: The unique ID for a TOTP instance.
	TOTPID string `json:"totp_id,omitempty"`
	// Secret: The TOTP secret key shared between the authenticator app and the server used to generate TOTP
	// codes.
	Secret string `json:"secret,omitempty"`
	// QrCode: The QR code image encoded in base64.
	QrCode string `json:"qr_code,omitempty"`
	// RecoveryCodes: The recovery codes used to authenticate the user without an authenticator app.
	RecoveryCodes []string `json:"recovery_codes,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User users.User `json:"user,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// RecoverResponse: Response type for `TOTPs.Recover`.
type RecoverResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// TOTPID: The unique ID for a TOTP instance.
	TOTPID string `json:"totp_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
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
	//   See [Session object](https://stytch.com/docs/api/session-object) for complete response fields.
	//
	Session *sessions.Session `json:"session,omitempty"`
}

// RecoveryCodesResponse: Response type for `TOTPs.RecoveryCodes`.
type RecoveryCodesResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// TOTPs: An array containing a list of all TOTP instances (along with their recovery codes) for a given
	// User in the Stytch API.
	TOTPs []TOTP `json:"totps,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}
