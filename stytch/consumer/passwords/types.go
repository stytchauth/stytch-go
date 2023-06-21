package passwords

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v8/stytch/consumer/sessions"
	"github.com/stytchauth/stytch-go/v8/stytch/consumer/users"
)

// Argon2Config: Fields:
//   - Salt: The salt value.
//   - IterationAmount: The iteration amount.
//   - Memory: The memory in kibibytes.
//   - Threads: The thread value, also known as the parallelism factor.
//   - KeyLength: The key length, also known as the hash length.
type Argon2Config struct {
	Salt            string `json:"salt,omitempty"`
	IterationAmount int32  `json:"iteration_amount,omitempty"`
	Memory          int32  `json:"memory,omitempty"`
	Threads         int32  `json:"threads,omitempty"`
	KeyLength       int32  `json:"key_length,omitempty"`
}

// AuthenticateParams: Request type for `Authenticate`.
// Fields:
//
//   - Email: The email address of the end user.
//
//   - Password: The password of the user
//
//   - SessionToken: The `session_token` associated with a User's existing Session.
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
//     If the `session_duration_minutes` parameter is not specified, a Stytch session will not be created.
//
//   - SessionJWT: The `session_jwt` associated with a User's existing Session.
//
//   - SessionCustomClaims: Add a custom claims map to the Session being authenticated. Claims are only
//     created if a Session is initialized by providing a value in `session_duration_minutes`. Claims will be
//     included on the Session object and in the JWT. To update a key in an existing Session, supply a new
//     value. To delete a key, supply a null value.
//
//     Custom claims made with reserved claims ("iss", "sub", "aud", "exp", "nbf", "iat", "jti") will be
//     ignored. Total custom claims size cannot exceed four kilobytes.
type AuthenticateParams struct {
	Email                  string         `json:"email,omitempty"`
	Password               string         `json:"password,omitempty"`
	SessionToken           string         `json:"session_token,omitempty"`
	SessionDurationMinutes int32          `json:"session_duration_minutes,omitempty"`
	SessionJWT             string         `json:"session_jwt,omitempty"`
	SessionCustomClaims    map[string]any `json:"session_custom_claims,omitempty"`
}

// CreateParams: Request type for `Create`.
// Fields:
//
//   - Email: The email address of the end user.
//
//   - Password: The password of the user
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
//     If the `session_duration_minutes` parameter is not specified, a Stytch session will not be created.
//
//   - SessionCustomClaims: Add a custom claims map to the Session being authenticated. Claims are only
//     created if a Session is initialized by providing a value in `session_duration_minutes`. Claims will be
//     included on the Session object and in the JWT. To update a key in an existing Session, supply a new
//     value. To delete a key, supply a null value.
//
//     Custom claims made with reserved claims ("iss", "sub", "aud", "exp", "nbf", "iat", "jti") will be
//     ignored. Total custom claims size cannot exceed four kilobytes.
//
//   - TrustedMetadata: The `trusted_metadata` field contains an arbitrary JSON object of
//     application-specific data. See the [Metadata](https://stytch.com/docs/api/metadata) reference for
//     complete field behavior details.
//
//   - UntrustedMetadata: The `untrusted_metadata` field contains an arbitrary JSON object of
//     application-specific data. Untrusted metadata can be edited by end users directly via the SDK, and
//     **cannot be used to store critical information.** See the
//     [Metadata](https://stytch.com/docs/api/metadata) reference for complete field behavior details.
//
//   - Name: The name of the user. Each field in the name object is optional.
type CreateParams struct {
	Email                  string         `json:"email,omitempty"`
	Password               string         `json:"password,omitempty"`
	SessionDurationMinutes int32          `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]any `json:"session_custom_claims,omitempty"`
	TrustedMetadata        map[string]any `json:"trusted_metadata,omitempty"`
	UntrustedMetadata      map[string]any `json:"untrusted_metadata,omitempty"`
	Name                   users.Name     `json:"name,omitempty"`
}

// Feedback: Fields:
//   - Warning: For `zxcvbn` validation, contains an end user consumable warning if the password is valid
//     but not strong enough.
//   - Suggestions: For `zxcvbn` validation, contains end user consumable suggestions on how to improve the
//     strength of the password.
//   - LudsRequirements: Contains which LUDS properties are fulfilled by the password and which are missing
//     to convert an invalid password into a valid one. You'll use these fields to provide feedback to the user
//     on how to improve the password.
type Feedback struct {
	Warning          string           `json:"warning,omitempty"`
	Suggestions      []string         `json:"suggestions,omitempty"`
	LudsRequirements LUDSRequirements `json:"luds_requirements,omitempty"`
}

// LUDSRequirements: Fields:
//   - HasLowerCase: For LUDS validation, whether the password contains at least one lowercase letter.
//   - HasUpperCase: For LUDS validation, whether the password contains at least one uppercase letter.
//   - HasDigit: For LUDS validation, whether the password contains at least one digit.
//   - HasSymbol: For LUDS validation, whether the password contains at least one symbol. Any UTF8
//     character outside of a-z or A-Z may count as a valid symbol.
//   - MissingComplexity: For LUDS validation, the number of complexity requirements that are missing from
//     the password. Check the complexity fields to see which requirements are missing.
//   - MissingCharacters: For LUDS validation, this is the required length of the password that you've set
//     minus the length of the password being checked. The user will need to add this many characters to the
//     password to make it valid.
type LUDSRequirements struct {
	HasLowerCase      bool  `json:"has_lower_case,omitempty"`
	HasUpperCase      bool  `json:"has_upper_case,omitempty"`
	HasDigit          bool  `json:"has_digit,omitempty"`
	HasSymbol         bool  `json:"has_symbol,omitempty"`
	MissingComplexity int32 `json:"missing_complexity,omitempty"`
	MissingCharacters int32 `json:"missing_characters,omitempty"`
}

// MD5Config: Fields:
//   - PrependSalt: The salt that should be prepended to the migrated password.
//   - AppendSalt: The salt that should be appended to the migrated password.
type MD5Config struct {
	PrependSalt string `json:"prepend_salt,omitempty"`
	AppendSalt  string `json:"append_salt,omitempty"`
}

// MigrateParams: Request type for `Migrate`.
// Fields:
//   - Email: The email address of the end user.
//   - Hash: The password hash. For a Scrypt hash, the hash needs to be a base64 encoded string.
//   - HashType: The password hash used. Currently `bcrypt`, `scrypt`, `argon_2i`, `argon_2id`, `md_5`, and
//     `sha_1` are supported.
//   - Md5Config: Optional parameters for MD-5 hash types.
//   - Argon2Config: Required parameters if the argon2 hex form, as opposed to the encoded form, is
//     supplied.
//   - Sha1Config: Optional parameters for SHA-1 hash types.
//   - ScryptConfig: Required parameters if the scrypt is not provided in a [PHC encoded
//     form](https://github.com/P-H-C/phc-string-format/blob/master/phc-sf-spec.md#phc-string-format).
//   - TrustedMetadata: The `trusted_metadata` field contains an arbitrary JSON object of
//     application-specific data. See the [Metadata](https://stytch.com/docs/api/metadata) reference for
//     complete field behavior details.
//   - UntrustedMetadata: The `untrusted_metadata` field contains an arbitrary JSON object of
//     application-specific data. Untrusted metadata can be edited by end users directly via the SDK, and
//     **cannot be used to store critical information.** See the
//     [Metadata](https://stytch.com/docs/api/metadata) reference for complete field behavior details.
//   - Name: The name of the user. Each field in the name object is optional.
type MigrateParams struct {
	Email             string         `json:"email,omitempty"`
	Hash              string         `json:"hash,omitempty"`
	HashType          string         `json:"hash_type,omitempty"`
	Md5Config         MD5Config      `json:"md_5_config,omitempty"`
	Argon2Config      Argon2Config   `json:"argon_2_config,omitempty"`
	Sha1Config        SHA1Config     `json:"sha_1_config,omitempty"`
	ScryptConfig      ScryptConfig   `json:"scrypt_config,omitempty"`
	TrustedMetadata   map[string]any `json:"trusted_metadata,omitempty"`
	UntrustedMetadata map[string]any `json:"untrusted_metadata,omitempty"`
	Name              users.Name     `json:"name,omitempty"`
}

// SHA1Config: Fields:
//   - PrependSalt: The salt that should be prepended to the migrated password.
//   - AppendSalt: The salt that should be appended to the migrated password.
type SHA1Config struct {
	PrependSalt string `json:"prepend_salt,omitempty"`
	AppendSalt  string `json:"append_salt,omitempty"`
}

// ScryptConfig: Fields:
//   - Salt: The salt value, which should be in a base64 encoded string form.
//   - NParameter: The N value, also known as the iterations count. It must be a power of two greater than
//     1 and less than 262,145.
//     If your applicaiton's N parameter is larger than 262,144, please reach out to
//     [support@stytch.com](mailto:support@stytch.com)
//   - RParameter: The r parameter, also known as the block size.
//   - PParameter: The p parameter, also known as the parallelism factor.
//   - KeyLength: The key length, also known as the hash length.
type ScryptConfig struct {
	Salt       string `json:"salt,omitempty"`
	NParameter int32  `json:"n_parameter,omitempty"`
	RParameter int32  `json:"r_parameter,omitempty"`
	PParameter int32  `json:"p_parameter,omitempty"`
	KeyLength  int32  `json:"key_length,omitempty"`
}

// StrengthCheckParams: Request type for `StrengthCheck`.
// Fields:
//   - Password: The password of the user
//   - Email: The email address of the end user.
type StrengthCheckParams struct {
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

// AuthenticateResponse: Response type for `Authenticate`.
// Fields:
//
//   - RequestID: Globally unique UUID that is returned with every API call. This value is important to log
//     for debugging purposes; we may ask for this value to help identify a specific API call when helping you
//     debug an issue.
//
//   - UserID: The unique ID of the affected User.
//
//   - SessionToken: A secret token for a given Stytch Session.
//
//   - SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
//
//   - User: The `user` object affected by this API call. See the [Get user
//     endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
//
//   - StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
//     patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
//     are server errors.
//
//   - Session: If you initiate a Session, by including `session_duration_minutes` in your authenticate
//     call, you'll receive a full Session object in the response.
//
//     See [GET sessions](https://stytch.com/docs/api/session-get) for complete response fields.
type AuthenticateResponse struct {
	RequestID    string           `json:"request_id,omitempty"`
	UserID       string           `json:"user_id,omitempty"`
	SessionToken string           `json:"session_token,omitempty"`
	SessionJWT   string           `json:"session_jwt,omitempty"`
	User         users.User       `json:"user,omitempty"`
	StatusCode   int32            `json:"status_code,omitempty"`
	Session      sessions.Session `json:"session,omitempty"`
}

// CreateResponse: Response type for `Create`.
// Fields:
//
//   - RequestID: Globally unique UUID that is returned with every API call. This value is important to log
//     for debugging purposes; we may ask for this value to help identify a specific API call when helping you
//     debug an issue.
//
//   - UserID: The unique ID of the affected User.
//
//   - EmailID: The unique ID of a specific email address.
//
//   - SessionToken: A secret token for a given Stytch Session.
//
//   - SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
//
//   - User: The `user` object affected by this API call. See the [Get user
//     endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
//
//   - StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
//     patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
//     are server errors.
//
//   - Session: If you initiate a Session, by including `session_duration_minutes` in your authenticate
//     call, you'll receive a full Session object in the response.
//
//     See [GET sessions](https://stytch.com/docs/api/session-get) for complete response fields.
type CreateResponse struct {
	RequestID    string           `json:"request_id,omitempty"`
	UserID       string           `json:"user_id,omitempty"`
	EmailID      string           `json:"email_id,omitempty"`
	SessionToken string           `json:"session_token,omitempty"`
	SessionJWT   string           `json:"session_jwt,omitempty"`
	User         users.User       `json:"user,omitempty"`
	StatusCode   int32            `json:"status_code,omitempty"`
	Session      sessions.Session `json:"session,omitempty"`
}

// MigrateResponse: Response type for `Migrate`.
// Fields:
//   - RequestID: Globally unique UUID that is returned with every API call. This value is important to log
//     for debugging purposes; we may ask for this value to help identify a specific API call when helping you
//     debug an issue.
//   - UserID: The unique ID of the affected User.
//   - EmailID: The unique ID of a specific email address.
//   - UserCreated: In `login_or_create` endpoints, this field indicates whether or not a User was just
//     created.
//   - User: The `user` object affected by this API call. See the [Get user
//     endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
//   - StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
//     patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
//     are server errors.
type MigrateResponse struct {
	RequestID   string     `json:"request_id,omitempty"`
	UserID      string     `json:"user_id,omitempty"`
	EmailID     string     `json:"email_id,omitempty"`
	UserCreated bool       `json:"user_created,omitempty"`
	User        users.User `json:"user,omitempty"`
	StatusCode  int32      `json:"status_code,omitempty"`
}

// StrengthCheckResponse: Response type for `StrengthCheck`.
// Fields:
//   - RequestID: Globally unique UUID that is returned with every API call. This value is important to log
//     for debugging purposes; we may ask for this value to help identify a specific API call when helping you
//     debug an issue.
//   - ValidPassword: Returns `true` if the password passes our password validation. We offer two
//     validation options, [zxcvbn](https://stytch.com/docs/passwords#strength-requirements) is the default
//     option which offers a high level of sophistication. We also offer
//     [LUDS](https://stytch.com/docs/passwords#strength-requirements). If an email address is included in the
//     call we also require that the password hasn't been compromised using built-in breach detection powered
//     by [HaveIBeenPwned](https://haveibeenpwned.com/).
//   - Score: The score of the password determined by [zxcvbn](https://github.com/dropbox/zxcvbn). Values
//     will be between 1 and 4, a 3 or greater is required to pass validation.
//   - BreachedPassword: Returns `true` if the password has been breached. Powered by
//     [HaveIBeenPwned](https://haveibeenpwned.com/).
//   - StrengthPolicy: The strength policy type enforced, either `zxcvbn` or `luds`.
//   - BreachDetectionOnCreate: Will return `true` if breach detection will be evaluated. By default this
//     option is enabled. This option can be disabled by contacting
//     [support@stytch.com](mailto:support@stytch.com?subject=Password%20strength%20configuration). If this
//     value is `false` then `breached_password` will always be `false` as well.
//   - StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
//     patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
//     are server errors.
//   - Feedback: Feedback for how to improve the password's strength
//     [HaveIBeenPwned](https://haveibeenpwned.com/).
type StrengthCheckResponse struct {
	RequestID               string   `json:"request_id,omitempty"`
	ValidPassword           bool     `json:"valid_password,omitempty"`
	Score                   int32    `json:"score,omitempty"`
	BreachedPassword        bool     `json:"breached_password,omitempty"`
	StrengthPolicy          string   `json:"strength_policy,omitempty"`
	BreachDetectionOnCreate bool     `json:"breach_detection_on_create,omitempty"`
	StatusCode              int32    `json:"status_code,omitempty"`
	Feedback                Feedback `json:"feedback,omitempty"`
}
