package existingpassword

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v15/stytch/b2b/mfa"
	"github.com/stytchauth/stytch-go/v15/stytch/b2b/organizations"
	"github.com/stytchauth/stytch-go/v15/stytch/b2b/sessions"
)

// ResetParams: Request type for `ExistingPassword.Reset`.
type ResetParams struct {
	// EmailAddress: The email address of the Member.
	EmailAddress string `json:"email_address,omitempty"`
	// ExistingPassword: The Member's current password that they supplied.
	ExistingPassword string `json:"existing_password,omitempty"`
	// NewPassword: The Member's elected new password.
	NewPassword string `json:"new_password,omitempty"`
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
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
	Locale ResetRequestLocale `json:"locale,omitempty"`
}

// ResetResponse: Response type for `ExistingPassword.Reset`.
type ResetResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
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

type ResetRequestLocale string

const (
	ResetRequestLocaleEn   ResetRequestLocale = "en"
	ResetRequestLocaleEs   ResetRequestLocale = "es"
	ResetRequestLocalePtbr ResetRequestLocale = "pt-br"
)
