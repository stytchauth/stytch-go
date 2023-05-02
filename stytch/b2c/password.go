package b2c

import "github.com/stytchauth/stytch-go/v8/stytch/shared"

type PasswordsCreateParams struct {
	Email                  string                 `json:"email"`
	Password               string                 `json:"password"`
	Name                   Name                   `json:"name,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
	TrustedMetadata        map[string]interface{} `json:"trusted_metadata,omitempty"`
	UntrustedMetadata      map[string]interface{} `json:"untrusted_metadata,omitempty"`
}

type PasswordsCreateResponse struct {
	RequestID    string  `json:"request_id,omitempty"`
	StatusCode   int     `json:"status_code,omitempty"`
	UserID       string  `json:"user_id,omitempty"`
	EmailID      string  `json:"email_id,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
	Session      Session `json:"session,omitempty"`
	SessionJWT   string  `json:"session_jwt,omitempty"`
	User         User    `json:"user,omitempty"`
}

type PasswordsAuthenticateParams struct {
	Email                  string                 `json:"email"`
	Password               string                 `json:"password"`
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
}

type PasswordsAuthenticateResponse struct {
	RequestID    string  `json:"request_id,omitempty"`
	StatusCode   int     `json:"status_code,omitempty"`
	UserID       string  `json:"user_id,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
	Session      Session `json:"session,omitempty"`
	SessionJWT   string  `json:"session_jwt,omitempty"`
	User         User    `json:"user,omitempty"`
}

type PasswordsStrengthCheckParams struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type PasswordsStrengthCheckResponse struct {
	RequestID               string   `json:"request_id,omitempty"`
	StatusCode              int      `json:"status_code,omitempty"`
	ValidPassword           bool     `json:"valid_password,omitempty"`
	Score                   int      `json:"score,omitempty"`
	BreachedPassword        bool     `json:"breached_password,omitempty"`
	Feedback                Feedback `json:"feedback,omitempty"`
	StrengthPolicy          string   `json:"strength_policy,omitempty"`
	BreachDetectionOnCreate bool     `json:"breach_detection_on_create,omitempty"`
}

type Feedback struct {
	Warning          string           `json:"warning,omitempty"`
	Suggestions      []string         `json:"suggestions,omitempty"`
	LUDSRequirements LUDSRequirements `json:"luds_requirements,omitempty"`
}

type LUDSRequirements struct {
	HasLowerCase      bool `json:"has_lower_case,omitempty"`
	HasUpperCase      bool `json:"has_upper_case,omitempty"`
	HasDigit          bool `json:"has_digit,omitempty"`
	HasSymbol         bool `json:"has_symbol,omitempty"`
	MissingComplexity int  `json:"missing_complexity,omitempty"`
	MissingCharacters int  `json:"missing_characters,omitempty"`
}

type PasswordsMigrateParams struct {
	Email             string                 `json:"email"`
	Name              Name                   `json:"name,omitempty"`
	TrustedMetadata   map[string]interface{} `json:"trusted_metadata,omitempty"`
	UntrustedMetadata map[string]interface{} `json:"untrusted_metadata,omitempty"`
	Hash              string                 `json:"hash"`
	HashType          shared.HashType        `json:"hash_type"`
	MD5Config         shared.MD5Config       `json:"md_5_config,omitempty"`
	Argon2Config      shared.Argon2Config    `json:"argon_2_config,omitempty"`
	SHA1Config        shared.SHA1Config      `json:"sha_1_config,omitempty"`
	ScryptConfig      shared.ScryptConfig    `json:"scrypt_config,omitempty"`
}

type PasswordsMigrateResponse struct {
	RequestID   string `json:"request_id,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	EmailID     string `json:"email_id,omitempty"`
	UserCreated bool   `json:"user_created,omitempty"`
}

// PASSWORD - EMAIL
type PasswordEmailResetStartParams struct {
	Email                          string     `json:"email"`
	LoginRedirectURL               string     `json:"login_redirect_url,omitempty"`
	ResetPasswordRedirectURL       string     `json:"reset_password_redirect_url,omitempty"`
	ResetPasswordExpirationMinutes int32      `json:"reset_password_expiration_minutes,omitempty"`
	ResetPasswordTemplateID        string     `json:"reset_password_template_id,omitempty"`
	Attributes                     Attributes `json:"attributes,omitempty"`
	CodeChallenge                  string     `json:"code_challenge,omitempty"`
	Locale                         string     `json:"locale,omitempty"`
}

type PasswordEmailResetStartResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	EmailID    string `json:"email_id,omitempty"`
}

type PasswordEmailResetParams struct {
	Token                  string                 `json:"token,omitempty"`
	Password               string                 `json:"password,omitempty"`
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
	Options                Options                `json:"options,omitempty"`
	Attributes             Attributes             `json:"attributes,omitempty"`
	CodeVerifier           string                 `json:"code_verifier,omitempty"`
}

type PasswordEmailResetResponse struct {
	RequestID    string  `json:"request_id,omitempty"`
	StatusCode   int     `json:"status_code,omitempty"`
	UserID       string  `json:"user_id,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
	SessionJWT   string  `json:"session_jwt,omitempty"`
	Session      Session `json:"session,omitempty"`
	User         User    `json:"user,omitempty"`
}

// PASSWORD - EXISTING PASSWORD
type PasswordExistingPasswordResetParams struct {
	Email                  string                 `json:"email,omitempty"`
	ExistingPassword       string                 `json:"existing_password,omitempty"`
	NewPassword            string                 `json:"new_password,omitempty"`
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
}

type PasswordExistingPasswordResetResponse struct {
	RequestID    string  `json:"request_id,omitempty"`
	StatusCode   int     `json:"status_code,omitempty"`
	UserID       string  `json:"user_id,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
	SessionJWT   string  `json:"session_jwt,omitempty"`
	Session      Session `json:"session,omitempty"`
	User         User    `json:"user,omitempty"`
}

// PASSWORD - SESSION
type PasswordSessionResetParams struct {
	Password     string `json:"password,omitempty"`
	SessionToken string `json:"session_token,omitempty"`
	SessionJWT   string `json:"session_jwt,omitempty"`
}

type PasswordSessionResetResponse struct {
	RequestID  string  `json:"request_id,omitempty"`
	StatusCode int     `json:"status_code,omitempty"`
	UserID     string  `json:"user_id,omitempty"`
	User       User    `json:"user,omitempty"`
	Session    Session `json:"session,omitempty"`
}
