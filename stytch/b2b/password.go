package b2b

import "github.com/stytchauth/stytch-go/v9/stytch/shared"

type PasswordsAuthenticateParams struct {
	OrganizationID         string                 `json:"organization_id,omitempty"`
	EmailAddress           string                 `json:"email_address"`
	Password               string                 `json:"password"`
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
}

type PasswordsAuthenticateResponse struct {
	RequestID      string        `json:"request_id,omitempty"`
	StatusCode     int           `json:"status_code,omitempty"`
	MemberID       string        `json:"member_id,omitempty"`
	OrganizationID string        `json:"organization_id,omitempty"`
	Member         Member        `json:"member,omitempty"`
	SessionToken   string        `json:"session_token,omitempty"`
	SessionJWT     string        `json:"session_jwt,omitempty"`
	MemberSession  MemberSession `json:"member_session,omitempty"`
	Organization   Organization  `json:"organization,omitempty"`
}

type PasswordsStrengthCheckParams struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

type PasswordsStrengthCheckResponse struct {
	RequestID               string         `json:"request_id,omitempty"`
	StatusCode              int            `json:"status_code,omitempty"`
	ValidPassword           bool           `json:"valid_password,omitempty"`
	Score                   int            `json:"score,omitempty"`
	BreachedPassword        bool           `json:"breached_password,omitempty"`
	ZXCVBNFeedback          ZXCVBNFeedback `json:"zxcvbn_feedback,omitempty"`
	LUDSFeedback            LUDSFeedback   `json:"luds_feedback,omitempty"`
	StrengthPolicy          string         `json:"strength_policy,omitempty"`
	BreachDetectionOnCreate bool           `json:"breach_detection_on_create,omitempty"`
}

type ZXCVBNFeedback struct {
	Warning     string   `json:"warning,omitempty"`
	Suggestions []string `json:"suggestions,omitempty"`
}

type LUDSFeedback struct {
	HasLowerCase      bool `json:"has_lower_case,omitempty"`
	HasUpperCase      bool `json:"has_upper_case,omitempty"`
	HasDigit          bool `json:"has_digit,omitempty"`
	HasSymbol         bool `json:"has_symbol,omitempty"`
	MissingComplexity int  `json:"missing_complexity,omitempty"`
	MissingCharacters int  `json:"missing_characters,omitempty"`
}

type PasswordsMigrateParams struct {
	OrganizationID    string                 `json:"organization_id"`
	EmailAddress      string                 `json:"email_address"`
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

type PasswordsDeleteResponse struct {
	RequestID   string `json:"request_id,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	EmailID     string `json:"email_id,omitempty"`
	UserCreated bool   `json:"user_created,omitempty"`
}

// PASSWORD - EMAIL
type PasswordEmailResetStartParams struct {
	OrganizationID                 string `json:"organization_id"`
	EmailAddress                   string `json:"email_address"`
	LoginRedirectURL               string `json:"login_redirect_url,omitempty"`
	ResetPasswordRedirectURL       string `json:"reset_password_redirect_url,omitempty"`
	ResetPasswordExpirationMinutes int32  `json:"reset_password_expiration_minutes,omitempty"`
	ResetPasswordTemplateID        string `json:"reset_password_template_id,omitempty"`
	CodeChallenge                  string `json:"code_challenge,omitempty"`
	Locale                         string `json:"locale,omitempty"`
}

type PasswordEmailResetStartResponse struct {
	RequestID     string `json:"request_id,omitempty"`
	StatusCode    int    `json:"status_code,omitempty"`
	MemberID      string `json:"user_id,omitempty"`
	MemberEmailID string `json:"email_id,omitempty"`
}

type PasswordEmailResetParams struct {
	PasswordResetToken     string                 `json:"password_reset_token"`
	Password               string                 `json:"password"`
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
	CodeVerifier           string                 `json:"code_verifier,omitempty"`
}

type PasswordEmailResetResponse struct {
	RequestID      string        `json:"request_id,omitempty"`
	StatusCode     int           `json:"status_code,omitempty"`
	MemberID       string        `json:"member_id,omitempty"`
	OrganizationID string        `json:"organization_id,omitempty"`
	Member         Member        `json:"member,omitempty"`
	SessionToken   string        `json:"session_token,omitempty"`
	SessionJWT     string        `json:"session_jwt,omitempty"`
	MemberSession  MemberSession `json:"member_session,omitempty"`
	Organization   Organization  `json:"organization,omitempty"`
}

// PASSWORD - EXISTING PASSWORD
type PasswordExistingPasswordResetParams struct {
	OrganizationID         string                 `json:"organization_id"`
	EmailAddress           string                 `json:"email_address"`
	ExistingPassword       string                 `json:"existing_password"`
	NewPassword            string                 `json:"new_password"`
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
}

type PasswordExistingPasswordResetResponse struct {
	RequestID      string        `json:"request_id,omitempty"`
	StatusCode     int           `json:"status_code,omitempty"`
	MemberID       string        `json:"member_id,omitempty"`
	OrganizationID string        `json:"organization_id,omitempty"`
	Member         Member        `json:"member,omitempty"`
	SessionToken   string        `json:"session_token,omitempty"`
	SessionJWT     string        `json:"session_jwt,omitempty"`
	MemberSession  MemberSession `json:"member_session,omitempty"`
	Organization   Organization  `json:"organization,omitempty"`
}

// PASSWORD - SESSION
type PasswordSessionResetParams struct {
	OrganizationID string `json:"organization_id"`
	Password       string `json:"password"`
	SessionToken   string `json:"session_token,omitempty"`
	SessionJWT     string `json:"session_jwt,omitempty"`
}

type PasswordSessionResetResponse struct {
	RequestID     string        `json:"request_id,omitempty"`
	StatusCode    int           `json:"status_code,omitempty"`
	MemberID      string        `json:"member_id,omitempty"`
	Member        Member        `json:"member,omitempty"`
	MemberSession MemberSession `json:"member_session,omitempty"`
	Organization  Organization  `json:"organization,omitempty"`
}
