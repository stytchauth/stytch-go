package passwords

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/mfa"
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/organizations"
	"github.com/stytchauth/stytch-go/v16/stytch/b2b/sessions"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/passwords"
)

// AuthenticateParams: Request type for `Passwords.Authenticate`.
type AuthenticateParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value. You may also use
	// the organization_slug here as a convenience.
	OrganizationID string `json:"organization_id,omitempty"`
	// EmailAddress: The email address of the Member.
	EmailAddress string `json:"email_address,omitempty"`
	// Password: The password to authenticate, reset, or set for the first time. Any UTF8 character is allowed,
	// e.g. spaces, emojis, non-English characers, etc.
	Password string `json:"password,omitempty"`
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
	Locale *AuthenticateRequestLocale `json:"locale,omitempty"`
	// IntermediateSessionToken: Adds this primary authentication factor to the intermediate session token. If
	// the resulting set of factors satisfies the organization's primary authentication requirements and MFA
	// requirements, the intermediate session token will be consumed and converted to a member session. If not,
	// the same intermediate session token will be returned.
	IntermediateSessionToken string `json:"intermediate_session_token,omitempty"`
}

// LudsFeedback:
type LudsFeedback struct {
	// HasLowerCase: For LUDS validation, whether the password contains at least one lowercase letter.
	HasLowerCase bool `json:"has_lower_case,omitempty"`
	// HasUpperCase: For LUDS validation, whether the password contains at least one uppercase letter.
	HasUpperCase bool `json:"has_upper_case,omitempty"`
	// HasDigit: For LUDS validation, whether the password contains at least one digit.
	HasDigit bool `json:"has_digit,omitempty"`
	// HasSymbol: For LUDS validation, whether the password contains at least one symbol. Any UTF8 character
	// outside of a-z or A-Z may count as a valid symbol.
	HasSymbol bool `json:"has_symbol,omitempty"`
	// MissingComplexity: For LUDS validation, the number of complexity requirements that are missing from the
	// password.
	//       Check the complexity fields to see which requirements are missing.
	MissingComplexity int32 `json:"missing_complexity,omitempty"`
	// MissingCharacters: For LUDS validation, this is the required length of the password that you've set
	// minus the length of the password being checked.
	//       The user will need to add this many characters to the password to make it valid.
	MissingCharacters int32 `json:"missing_characters,omitempty"`
}

// MigrateParams: Request type for `Passwords.Migrate`.
type MigrateParams struct {
	// EmailAddress: The email address of the Member.
	EmailAddress string `json:"email_address,omitempty"`
	// Hash: The password hash. For a Scrypt or PBKDF2 hash, the hash needs to be a base64 encoded string.
	Hash string `json:"hash,omitempty"`
	// HashType: The password hash used. Currently `bcrypt`, `scrypt`, `argon_2i`, `argon2_id`, `md_5`,
	// `sha_1`, and `pbkdf_2` are supported.
	HashType MigrateRequestHashType `json:"hash_type,omitempty"`
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value. You may also use
	// the organization_slug here as a convenience.
	OrganizationID string `json:"organization_id,omitempty"`
	// Md5Config: Optional parameters for MD-5 hash types.
	Md5Config *passwords.MD5Config `json:"md_5_config,omitempty"`
	// Argon2Config: Required parameters if the argon2 hex form, as opposed to the encoded form, is supplied.
	Argon2Config *passwords.Argon2Config `json:"argon_2_config,omitempty"`
	// Sha1Config: Optional parameters for SHA-1 hash types.
	Sha1Config *passwords.SHA1Config `json:"sha_1_config,omitempty"`
	// ScryptConfig: Required parameters if the scrypt is not provided in a **PHC encoded form**.
	ScryptConfig *passwords.ScryptConfig `json:"scrypt_config,omitempty"`
	// Pbkdf2Config: Required additional parameters for PBKDF2 hash keys. Note that we use the SHA-256 by
	// default, please contact [support@stytch.com](mailto:support@stytch.com) if you use another hashing
	// function.
	Pbkdf2Config *passwords.PBKDF2Config `json:"pbkdf_2_config,omitempty"`
	// Name: The name of the Member. Each field in the name object is optional.
	Name string `json:"name,omitempty"`
	// TrustedMetadata: An arbitrary JSON object for storing application-specific data or
	// identity-provider-specific data.
	TrustedMetadata map[string]any `json:"trusted_metadata,omitempty"`
	// UntrustedMetadata: An arbitrary JSON object of application-specific data. These fields can be edited
	// directly by the
	//   frontend SDK, and should not be used to store critical information. See the
	// [Metadata resource](https://stytch.com/docs/b2b/api/metadata)
	//   for complete field behavior details.
	UntrustedMetadata map[string]any `json:"untrusted_metadata,omitempty"`
	// Roles to explicitly assign to this Member.
	//  Will completely replace any existing explicitly assigned roles. See the
	//  [RBAC guide](https://stytch.com/docs/b2b/guides/rbac/role-assignment) for more information about role
	// assignment.
	//
	//    If a Role is removed from a Member, and the Member is also implicitly assigned this Role from an SSO
	// connection
	//    or an SSO group, we will by default revoke any existing sessions for the Member that contain any SSO
	//    authentication factors with the affected connection ID. You can preserve these sessions by passing in
	// the
	//    `preserve_existing_sessions` parameter with a value of `true`.
	Roles []string `json:"roles,omitempty"`
	// PreserveExistingSessions: Whether to preserve existing sessions when explicit Roles that are revoked are
	// also implicitly assigned
	//   by SSO connection or SSO group. Defaults to `false` - that is, existing Member Sessions that contain
	// SSO
	//   authentication factors with the affected SSO connection IDs will be revoked.
	PreserveExistingSessions bool   `json:"preserve_existing_sessions,omitempty"`
	MFAPhoneNumber           string `json:"mfa_phone_number,omitempty"`
	SetPhoneNumberVerified   bool   `json:"set_phone_number_verified,omitempty"`
	// ExternalID: If a new member is created, this will set an identifier that can be used in API calls
	// wherever a member_id is expected. This is a string consisting of alphanumeric, `.`, `_`, or `-`
	// characters with a maximum length of 128 characters. External IDs must be unique within an organization,
	// but may be reused across different organizations in the same project. Note that if a member already
	// exists, this field will be ignored.
	ExternalID string `json:"external_id,omitempty"`
}

// StrengthCheckParams: Request type for `Passwords.StrengthCheck`.
type StrengthCheckParams struct {
	// Password: The password to authenticate, reset, or set for the first time. Any UTF8 character is allowed,
	// e.g. spaces, emojis, non-English characers, etc.
	Password string `json:"password,omitempty"`
	// EmailAddress: The email address of the Member.
	EmailAddress string `json:"email_address,omitempty"`
}

// ZxcvbnFeedback:
type ZxcvbnFeedback struct {
	// Warning: For zxcvbn validation, contains an end user consumable warning if the password is valid but not
	// strong enough.
	Warning string `json:"warning,omitempty"`
	// Suggestions: For zxcvbn validation, contains end user consumable suggestions on how to improve the
	// strength of the password.
	Suggestions []string `json:"suggestions,omitempty"`
}

// AuthenticateResponse: Response type for `Passwords.Authenticate`.
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

// MigrateResponse: Response type for `Passwords.Migrate`.
type MigrateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member.
	MemberID string `json:"member_id,omitempty"`
	// MemberCreated: A flag indicating `true` if a new Member object was created and `false` if the Member
	// object already existed.
	MemberCreated bool `json:"member_created,omitempty"`
	// Member: The [Member object](https://stytch.com/docs/b2b/api/member-object)
	Member organizations.Member `json:"member,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization organizations.Organization `json:"organization,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// StrengthCheckResponse: Response type for `Passwords.StrengthCheck`.
type StrengthCheckResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// ValidPassword: Returns `true` if the password passes our password validation. We offer two validation
	// options,
	//   [zxcvbn](https://stytch.com/docs/passwords#strength-requirements) is the default option which offers a
	// high level of sophistication.
	//   We also offer [LUDS](https://stytch.com/docs/passwords#strength-requirements). If an email address is
	// included in the call we also
	//   require that the password hasn't been compromised using built-in breach detection powered by
	// [HaveIBeenPwned](https://haveibeenpwned.com/)
	ValidPassword bool `json:"valid_password,omitempty"`
	// Score: The score of the password determined by [zxcvbn](https://github.com/dropbox/zxcvbn). Values will
	// be between 1 and 4, a 3 or greater is required to pass validation.
	Score int32 `json:"score,omitempty"`
	// BreachedPassword: Returns `true` if the password has been breached. Powered by
	// [HaveIBeenPwned](https://haveibeenpwned.com/).
	BreachedPassword bool `json:"breached_password,omitempty"`
	// StrengthPolicy: The strength policy type enforced, either `zxcvbn` or `luds`.
	StrengthPolicy string `json:"strength_policy,omitempty"`
	// BreachDetectionOnCreate: Will return `true` if breach detection will be evaluated. By default this
	// option is enabled.
	//   This option can be disabled by contacting
	// [support@stytch.com](mailto:support@stytch.com?subject=Password%20strength%20configuration).
	//   If this value is false then `breached_password` will always be `false` as well.
	BreachDetectionOnCreate bool `json:"breach_detection_on_create,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// LudsFeedback: Feedback for how to improve the password's strength using
	// [luds](https://stytch.com/docs/passwords#strength-requirements).
	LudsFeedback *LudsFeedback `json:"luds_feedback,omitempty"`
	// ZxcvbnFeedback: Feedback for how to improve the password's strength using
	// [zxcvbn](https://stytch.com/docs/passwords#strength-requirements).
	ZxcvbnFeedback *ZxcvbnFeedback `json:"zxcvbn_feedback,omitempty"`
}

type AuthenticateRequestLocale string

const (
	AuthenticateRequestLocaleEn   AuthenticateRequestLocale = "en"
	AuthenticateRequestLocaleEs   AuthenticateRequestLocale = "es"
	AuthenticateRequestLocalePtbr AuthenticateRequestLocale = "pt-br"
)

type MigrateRequestHashType string

const (
	MigrateRequestHashTypeBcrypt   MigrateRequestHashType = "bcrypt"
	MigrateRequestHashTypeMd5      MigrateRequestHashType = "md_5"
	MigrateRequestHashTypeArgon2i  MigrateRequestHashType = "argon_2i"
	MigrateRequestHashTypeArgon2id MigrateRequestHashType = "argon_2id"
	MigrateRequestHashTypeSha1     MigrateRequestHashType = "sha_1"
	MigrateRequestHashTypeScrypt   MigrateRequestHashType = "scrypt"
	MigrateRequestHashTypePhpass   MigrateRequestHashType = "phpass"
	MigrateRequestHashTypePbkdf2   MigrateRequestHashType = "pbkdf_2"
)
