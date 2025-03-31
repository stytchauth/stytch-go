package email

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/mfa"
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/organizations"
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/sessions"
	"github.com/stytchauth/stytch-go/v16/stytch/methodoptions"
)

// RequireResetParams: Request type for `Email.RequireReset`.
type RequireResetParams struct {
	// EmailAddress: The email address of the Member to start the email reset process for.
	EmailAddress string `json:"email_address,omitempty"`
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value. You may also use
	// the organization_slug here as a convenience.
	OrganizationID string `json:"organization_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member. The `member_id` is critical to perform
	// operations on a Member, so be sure to preserve this value. You may use an external_id here if one is set
	// for the member.
	MemberID string `json:"member_id,omitempty"`
}

// ResetParams: Request type for `Email.Reset`.
type ResetParams struct {
	// PasswordResetToken: The password reset token to authenticate.
	PasswordResetToken string `json:"password_reset_token,omitempty"`
	// Password: The password to authenticate, reset, or set for the first time. Any UTF8 character is allowed,
	// e.g. spaces, emojis, non-English characers, etc.
	Password string `json:"password,omitempty"`
	// SessionToken: Reuse an existing session instead of creating a new one. If you provide a `session_token`,
	// Stytch will update the session.
	//       If the `session_token` and `magic_links_token` belong to different Members, the `session_token`
	// will be ignored. This endpoint will error if
	//       both `session_token` and `session_jwt` are provided.
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
	// SessionJWT: Reuse an existing session instead of creating a new one. If you provide a `session_jwt`,
	// Stytch will update the session. If the `session_jwt`
	//       and `magic_links_token` belong to different Members, the `session_jwt` will be ignored. This
	// endpoint will error if both `session_token` and `session_jwt`
	//       are provided.
	SessionJWT string `json:"session_jwt,omitempty"`
	// CodeVerifier: A base64url encoded one time secret used to validate that the request starts and ends on
	// the same device.
	CodeVerifier string `json:"code_verifier,omitempty"`
	// SessionCustomClaims: Add a custom claims map to the Session being authenticated. Claims are only created
	// if a Session is initialized by providing a value in
	//   `session_duration_minutes`. Claims will be included on the Session object and in the JWT. To update a
	// key in an existing Session, supply a new value. To
	//   delete a key, supply a null value. Custom claims made with reserved claims (`iss`, `sub`, `aud`,
	// `exp`, `nbf`, `iat`, `jti`) will be ignored.
	//   Total custom claims size cannot exceed four kilobytes.
	SessionCustomClaims map[string]any `json:"session_custom_claims,omitempty"`
	// Locale: If the needs to complete an MFA step, and the Member has a phone number, this endpoint will
	// pre-emptively send a one-time passcode (OTP) to the Member's phone number. The locale argument will be
	// used to determine which language to use when sending the passcode.
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
	Locale *ResetRequestLocale `json:"locale,omitempty"`
	// IntermediateSessionToken: Adds this primary authentication factor to the intermediate session token. If
	// the resulting set of factors satisfies the organization's primary authentication requirements and MFA
	// requirements, the intermediate session token will be consumed and converted to a member session. If not,
	// the same intermediate session token will be returned.
	IntermediateSessionToken string `json:"intermediate_session_token,omitempty"`
}

// ResetStartParams: Request type for `Email.ResetStart`.
type ResetStartParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value. You may also use
	// the organization_slug here as a convenience.
	OrganizationID string `json:"organization_id,omitempty"`
	// EmailAddress: The email address of the Member to start the email reset process for.
	EmailAddress string `json:"email_address,omitempty"`
	// ResetPasswordRedirectURL: The URL that the Member clicks from the reset password link. This URL should
	// be an endpoint in the backend server that verifies the request by querying
	//   Stytch's authenticate endpoint and finishes the reset password flow. If this value is not passed, the
	// default `reset_password_redirect_url` that you set in your Dashboard is used.
	//   If you have not set a default `reset_password_redirect_url`, an error is returned.
	ResetPasswordRedirectURL string `json:"reset_password_redirect_url,omitempty"`
	// ResetPasswordExpirationMinutes: Sets a time limit after which the email link to reset the member's
	// password will no longer be valid.
	ResetPasswordExpirationMinutes int32 `json:"reset_password_expiration_minutes,omitempty"`
	// CodeChallenge: A base64url encoded SHA256 hash of a one time secret used to validate that the request
	// starts and ends on the same device.
	CodeChallenge string `json:"code_challenge,omitempty"`
	// LoginRedirectURL: The URL that the member clicks from the reset without password link. This URL should
	// be an endpoint in the backend server
	//       that verifies the request by querying Stytch's authenticate endpoint and finishes the magic link
	// flow. If this value is not passed, the
	//       default `login_redirect_url` that you set in your Dashboard is used. This value is only used if
	// magic links are enabled for the member. If
	//       you have not set a default `login_redirect_url` and magic links are not enabled for the member, an
	// error is returned.
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
	Locale *ResetStartRequestLocale `json:"locale,omitempty"`
	// ResetPasswordTemplateID: Use a custom template for reset password emails. By default, it will use your
	// default email template. The template must be a template using our built-in customizations or a custom
	// HTML email for Magic Links - Reset Password.
	ResetPasswordTemplateID string `json:"reset_password_template_id,omitempty"`
	VerifyEmailTemplateID   string `json:"verify_email_template_id,omitempty"`
}

// RequireResetRequestOptions:
type RequireResetRequestOptions struct {
	// Authorization: Optional authorization object.
	// Pass in an active Stytch Member session token or session JWT and the request
	// will be run using that member's permissions.
	Authorization methodoptions.Authorization `json:"authorization,omitempty"`
}

func (o *RequireResetRequestOptions) AddHeaders(headers map[string][]string) map[string][]string {
	headers = o.Authorization.AddHeaders(headers)
	return headers
}

// RequireResetResponse: Response type for `Email.RequireReset`.
type RequireResetResponse struct {
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object)
	Member *organizations.Member `json:"member,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization *organizations.Organization `json:"organization,omitempty"`
}

// ResetResponse: Response type for `Email.Reset`.
type ResetResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// MemberEmailID: Globally unique UUID that identifies a member's email
	MemberEmailID string `json:"member_email_id,omitempty"`
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object)
	Member organizations.Member `json:"member,omitempty"`
	// SessionToken: A secret token for a given Stytch Session.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
	SessionJWT string `json:"session_jwt,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
	// IntermediateSessionToken: The returned Intermediate Session Token contains a password factor associated
	// with the Member. If this value is non-empty, the member must complete an MFA step to finish logging in
	// to the Organization. The token can be used with the
	// [OTP SMS Authenticate endpoint](https://stytch.com/docs/b2b/api/authenticate-otp-sms),
	// [TOTP Authenticate endpoint](https://stytch.com/docs/b2b/api/authenticate-totp), or
	// [Recovery Codes Recover endpoint](https://stytch.com/docs/b2b/api/recovery-codes-recover) to complete an
	// MFA flow and log in to the Organization. Password factors are not transferable between Organizations, so
	// the intermediate session token is not valid for use with discovery endpoints.
	IntermediateSessionToken string `json:"intermediate_session_token,omitempty"`
	// MemberAuthenticated: Indicates whether the Member is fully authenticated. If false, the Member needs to
	// complete an MFA step to log in to the Organization.
	MemberAuthenticated bool `json:"member_authenticated,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// MemberSession: The [Session object](https://stytch.com/docs/b2b/api/session-object).
	MemberSession *sessions.MemberSession `json:"member_session,omitempty"`
	// MFARequired: Information about the MFA requirements of the Organization and the Member's options for
	// fulfilling MFA.
	MFARequired *mfa.MfaRequired `json:"mfa_required,omitempty"`
}

// ResetStartResponse: Response type for `Email.ResetStart`.
type ResetStartResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// MemberEmailID: Globally unique UUID that identifies a member's email
	MemberEmailID string `json:"member_email_id,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object)
	Member organizations.Member `json:"member,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

type ResetRequestLocale string

const (
	ResetRequestLocaleEn   ResetRequestLocale = "en"
	ResetRequestLocaleEs   ResetRequestLocale = "es"
	ResetRequestLocalePtbr ResetRequestLocale = "pt-br"
)

type ResetStartRequestLocale string

const (
	ResetStartRequestLocaleEn   ResetStartRequestLocale = "en"
	ResetStartRequestLocaleEs   ResetStartRequestLocale = "es"
	ResetStartRequestLocalePtbr ResetStartRequestLocale = "pt-br"
)
