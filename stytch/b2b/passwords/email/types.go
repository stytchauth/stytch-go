package email

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v8/stytch/b2b/organizations"
	"github.com/stytchauth/stytch-go/v8/stytch/b2b/sessions"
)

// ResetParams: Request type for `Reset`.
// Fields:
//
//   - PasswordResetToken: The password reset token to authenticate.
//
//   - Password: The password to reset.
//
//   - SessionToken: Reuse an existing session instead of creating a new one. If you provide a
//     `session_token`, Stytch will update the session.
//     If the `session_token` and `magic_links_token` belong to different Members, the `session_token`
//     will be ignored. This endpoint will error if
//     both `session_token` and `session_jwt` are provided.
//
//   - SessionDurationMinutes: Set the session lifetime to be this many minutes from now. This will start a
//     new session if one doesn't already exist,
//     returning both an opaque `session_token` and `session_jwt` for this session. Remember that the
//     `session_jwt` will have a fixed lifetime of
//     five minutes regardless of the underlying session duration, and will need to be refreshed over time.
//
//     This value must be a minimum of 5 and a maximum of 527040 minutes (366 days).
//
//     If a `session_token` or `session_jwt` is provided then a successful authentication will continue to
//     extend the session this many minutes.
//
//     If the `session_duration_minutes` parameter is not specified, a Stytch session will be created with a
//     60 minute duration. If you don't want
//     to use the Stytch session product, you can ignore the session fields in the response.
//
//   - SessionJWT: Reuse an existing session instead of creating a new one. If you provide a `session_jwt`,
//     Stytch will update the session. If the `session_jwt`
//     and `magic_links_token` belong to different Members, the `session_jwt` will be ignored. This
//     endpoint will error if both `session_token` and `session_jwt`
//     are provided.
//
//   - CodeVerifier: A base64url encoded one time secret used to validate that the request starts and ends
//     on the same device.
//
//   - SessionCustomClaims: Add a custom claims map to the Session being authenticated. Claims are only
//     created if a Session is initialized by providing a value in
//     `session_duration_minutes`. Claims will be included on the Session object and in the JWT. To update a
//     key in an existing Session, supply a new value. To
//     delete a key, supply a null value. Custom claims made with reserved claims (`iss`, `sub`, `aud`,
//     `exp`, `nbf`, `iat`, `jti`) will be ignored.
//     Total custom claims size cannot exceed four kilobytes.
type ResetParams struct {
	PasswordResetToken     string         `json:"password_reset_token,omitempty"`
	Password               string         `json:"password,omitempty"`
	SessionToken           string         `json:"session_token,omitempty"`
	SessionDurationMinutes int32          `json:"session_duration_minutes,omitempty"`
	SessionJWT             string         `json:"session_jwt,omitempty"`
	CodeVerifier           string         `json:"code_verifier,omitempty"`
	SessionCustomClaims    map[string]any `json:"session_custom_claims,omitempty"`
}

// ResetStartParams: Request type for `ResetStart`.
// Fields:
//
//   - OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id`
//     is critical to perform operations on an Organization, so be sure to preserve this value.
//
//   - EmailAddress: The email address of the Member to start the email reset process for.
//
//   - ResetPasswordRedirectURL: The URL that the Member clicks from the reset password link. This URL
//     should be an endpoint in the backend server that verifies the request by querying
//     Stytch's authenticate endpoint and finishes the reset password flow. If this value is not passed, the
//     default `reset_password_redirect_url` that you set in your Dashboard is used.
//     If you have not set a default `reset_password_redirect_url`, an error is returned.
//
//   - ResetPasswordExpirationMinutes: Sets a time limit after which the email link to reset the member's
//     password will no longer be valid.
//
//   - CodeChallenge: A base64url encoded SHA256 hash of a one time secret used to validate that the
//     request starts and ends on the same device.
//
//   - LoginRedirectURL: The URL that the member clicks from the reset without password link. This URL
//     should be an endpoint in the backend server
//     that verifies the request by querying Stytch's authenticate endpoint and finishes the magic link
//     flow. If this value is not passed, the
//     default `login_redirect_url` that you set in your Dashboard is used. This value is only used if
//     magic links are enabled for the member. If
//     you have not set a default `login_redirect_url` and magic links are not enabled for the member, an
//     error is returned.
//
//   - Locale: Used to determine which language to use when sending the user this delivery method.
//     Parameter is a [IETF BCP 47 language tag](https://www.w3.org/International/articles/language-tags/),
//     e.g. `"en"`.
//
//     Currently supported languages are English (`"en"`), Spanish (`"es"`), and Brazilian Portuguese
//     (`"pt-br"`); if no value is provided, the copy defaults to English.
//
//     Request support for additional languages
//     [here](https://docs.google.com/forms/d/e/1FAIpQLScZSpAu_m2AmLXRT3F3kap-s_mcV6UTBitYn6CdyWP0-o7YjQ/viewform?usp=sf_link")!
//
//   - ResetPasswordTemplateID: Use a custom template for reset password emails. By default, it will use
//     your default email template. The template must be a template using our built-in customizations or a
//     custom HTML email for Magic Links - Reset Password.
type ResetStartParams struct {
	OrganizationID                 string `json:"organization_id,omitempty"`
	EmailAddress                   string `json:"email_address,omitempty"`
	ResetPasswordRedirectURL       string `json:"reset_password_redirect_url,omitempty"`
	ResetPasswordExpirationMinutes int32  `json:"reset_password_expiration_minutes,omitempty"`
	CodeChallenge                  string `json:"code_challenge,omitempty"`
	LoginRedirectURL               string `json:"login_redirect_url,omitempty"`
	Locale                         string `json:"locale,omitempty"`
	ResetPasswordTemplateID        string `json:"reset_password_template_id,omitempty"`
}

// ResetResponse: Response type for `Reset`.
// Fields:
//   - RequestID: Globally unique UUID that is returned with every API call. This value is important to log
//     for debugging purposes; we may ask for this value to help identify a specific API call when helping you
//     debug an issue.
//   - MemberID: Globally unique UUID that identifies a specific Member.
//   - MemberEmailID: Globally unique UUID that identifies a member's email
//   - OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id`
//     is critical to perform operations on an Organization, so be sure to preserve this value.
//   - Member: The [Member object](https://stytch.com/docs/b2b/api/member-object).
//   - SessionToken: A secret token for a given Stytch Session.
//   - SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
//   - Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
//   - StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
//     patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
//     are server errors.
//   - MemberSession: The [Session object](https://stytch.com/docs/b2b/api/session-object).
type ResetResponse struct {
	RequestID      string                     `json:"request_id,omitempty"`
	MemberID       string                     `json:"member_id,omitempty"`
	MemberEmailID  string                     `json:"member_email_id,omitempty"`
	OrganizationID string                     `json:"organization_id,omitempty"`
	Member         organizations.Member       `json:"member,omitempty"`
	SessionToken   string                     `json:"session_token,omitempty"`
	SessionJWT     string                     `json:"session_jwt,omitempty"`
	Organization   organizations.Organization `json:"organization,omitempty"`
	StatusCode     int32                      `json:"status_code,omitempty"`
	MemberSession  sessions.MemberSession     `json:"member_session,omitempty"`
}

// ResetStartResponse: Response type for `ResetStart`.
// Fields:
//   - RequestID: Globally unique UUID that is returned with every API call. This value is important to log
//     for debugging purposes; we may ask for this value to help identify a specific API call when helping you
//     debug an issue.
//   - MemberID: Globally unique UUID that identifies a specific Member.
//   - MemberEmailID: Globally unique UUID that identifies a member's email
//   - StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
//     patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
//     are server errors.
type ResetStartResponse struct {
	RequestID     string `json:"request_id,omitempty"`
	MemberID      string `json:"member_id,omitempty"`
	MemberEmailID string `json:"member_email_id,omitempty"`
	StatusCode    int32  `json:"status_code,omitempty"`
}
