package email

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v15/stytch/consumer/attribute"
	"github.com/stytchauth/stytch-go/v15/stytch/consumer/magiclinks"
	"github.com/stytchauth/stytch-go/v15/stytch/consumer/sessions"
	"github.com/stytchauth/stytch-go/v15/stytch/consumer/users"
)

// ResetParams: Request type for `Email.Reset`.
type ResetParams struct {
	// Token: The Passwords `token` from the `?token=` query parameter in the URL.
	//
	//       In the redirect URL, the `stytch_token_type` will be `login` or `reset_password`.
	//
	//       See examples and read more about redirect URLs
	// [here](https://stytch.com/docs/guides/dashboard/redirect-urls).
	Token string `json:"token,omitempty"`
	// Password: The password of the user
	Password string `json:"password,omitempty"`
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
	// CodeVerifier: A base64url encoded one time secret used to validate that the request starts and ends on
	// the same device.
	CodeVerifier string `json:"code_verifier,omitempty"`
	// SessionCustomClaims: Add a custom claims map to the Session being authenticated. Claims are only created
	// if a Session is initialized by providing a value in `session_duration_minutes`. Claims will be included
	// on the Session object and in the JWT. To update a key in an existing Session, supply a new value. To
	// delete a key, supply a null value.
	//
	//   Custom claims made with reserved claims ("iss", "sub", "aud", "exp", "nbf", "iat", "jti") will be
	// ignored. Total custom claims size cannot exceed four kilobytes.
	SessionCustomClaims map[string]any `json:"session_custom_claims,omitempty"`
	// Attributes: Provided attributes help with fraud detection.
	Attributes *attribute.Attributes `json:"attributes,omitempty"`
	// Options: Specify optional security settings.
	Options *magiclinks.Options `json:"options,omitempty"`
}

// ResetStartParams: Request type for `Email.ResetStart`.
type ResetStartParams struct {
	// Email: The email of the User that requested the password reset.
	Email string `json:"email,omitempty"`
	// ResetPasswordRedirectURL: The url that the user clicks from the password reset email to finish the reset
	// password flow.
	//   This should be a url that your app receives and parses before showing your app's reset password page.
	//   After the user submits a new password to your app, it should send an API request to complete the
	// password reset process.
	//   If this value is not passed, the default reset password redirect URL that you set in your Dashboard is
	// used.
	//   If you have not set a default reset password redirect URL, an error is returned.
	ResetPasswordRedirectURL string `json:"reset_password_redirect_url,omitempty"`
	// ResetPasswordExpirationMinutes: Set the expiration for the password reset, in minutes. By default, it
	// expires in 30 minutes.
	//   The minimum expiration is 5 minutes and the maximum is 7 days (10080 mins).
	ResetPasswordExpirationMinutes int32 `json:"reset_password_expiration_minutes,omitempty"`
	// CodeChallenge: A base64url encoded SHA256 hash of a one time secret used to validate that the request
	// starts and ends on the same device.
	CodeChallenge string `json:"code_challenge,omitempty"`
	// Attributes: Provided attributes help with fraud detection.
	Attributes *attribute.Attributes `json:"attributes,omitempty"`
	// LoginRedirectURL: The URL Stytch redirects to after the OAuth flow is completed for a user that already
	// exists. This URL should be a route in your application which will run `oauth.authenticate` (see below)
	// and finish the login.
	//
	//   The URL must be configured as a Login URL in the [Redirect URL page](/dashboard/redirect-urls). If the
	// field is not specified, the default Login URL will be used.
	LoginRedirectURL string `json:"login_redirect_url,omitempty"`
	// Locale: Used to determine which language to use when sending the user this delivery method. Parameter is
	// a [IETF BCP 47 language tag](https://www.w3.org/International/articles/language-tags/), e.g. `"en"`.
	//
	// Currently supported languages are English (`"en"`), Spanish (`"es"`), and Brazilian Portuguese
	// (`"pt-br"`); if no value is provided, the copy defaults to English.
	//
	// Request support for additional languages
	// [here](https://docs.google.com/forms/d/e/1FAIpQLScZSpAu_m2AmLXRT3F3kap-s_mcV6UTBitYn6CdyWP0-o7YjQ/viewform?usp=sf_link")!
	//
	Locale ResetStartRequestLocale `json:"locale,omitempty"`
	// ResetPasswordTemplateID: Use a custom template for password reset emails. By default, it will use your
	// default email template.
	//   The template must be a template using our built-in customizations or a custom HTML email for Passwords
	// - Password reset.
	ResetPasswordTemplateID string `json:"reset_password_template_id,omitempty"`
}

// ResetResponse: Response type for `Email.Reset`.
type ResetResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
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
	//   See [GET sessions](https://stytch.com/docs/api/session-get) for complete response fields.
	//
	Session *sessions.Session `json:"session,omitempty"`
}

// ResetStartResponse: Response type for `Email.ResetStart`.
type ResetStartResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// EmailID: The unique ID of a specific email address.
	EmailID string `json:"email_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

type ResetStartRequestLocale string

const (
	ResetStartRequestLocaleEn   ResetStartRequestLocale = "en"
	ResetStartRequestLocaleEs   ResetStartRequestLocale = "es"
	ResetStartRequestLocalePtbr ResetStartRequestLocale = "pt-br"
)
