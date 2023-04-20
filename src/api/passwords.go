package stytch

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

type PasswordsCreateParams struct {
	Email                  string         `json:"email,omitempty"`
	Password               string         `json:"password,omitempty"`
	SessionDurationMinutes int32          `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]any `json:"session_custom_claims,omitempty"`
	TrustedMetadata        map[string]any `json:"trusted_metadata,omitempty"`
	UntrustedMetadata      map[string]any `json:"untrusted_metadata,omitempty"`
	Name                   Name           `json:"name,omitempty"`
}
type PasswordsAuthenticateParams struct {
	Email                  string         `json:"email,omitempty"`
	Password               string         `json:"password,omitempty"`
	SessionToken           string         `json:"session_token,omitempty"`
	SessionDurationMinutes int32          `json:"session_duration_minutes,omitempty"`
	SessionJwt             string         `json:"session_jwt,omitempty"`
	SessionCustomClaims    map[string]any `json:"session_custom_claims,omitempty"`
}
type PasswordsPasswordsemailresetstartParams struct {
	Email                          string     `json:"email,omitempty"`
	ResetPasswordRedirectUrl       string     `json:"reset_password_redirect_url,omitempty"`
	ResetPasswordExpirationMinutes int32      `json:"reset_password_expiration_minutes,omitempty"`
	CodeChallenge                  string     `json:"code_challenge,omitempty"`
	Attributes                     Attributes `json:"attributes,omitempty"`
	LoginRedirectUrl               string     `json:"login_redirect_url,omitempty"`
	Locale                         string     `json:"locale,omitempty"`
	ResetPasswordTemplateID        string     `json:"reset_password_template_id,omitempty"`
}
type PasswordsPasswordsemailresetParams struct {
	Token                  string         `json:"token,omitempty"`
	Password               string         `json:"password,omitempty"`
	SessionToken           string         `json:"session_token,omitempty"`
	SessionDurationMinutes int32          `json:"session_duration_minutes,omitempty"`
	SessionJwt             string         `json:"session_jwt,omitempty"`
	CodeVerifier           string         `json:"code_verifier,omitempty"`
	SessionCustomClaims    map[string]any `json:"session_custom_claims,omitempty"`
	Attributes             Attributes     `json:"attributes,omitempty"`
	Options                Options        `json:"options,omitempty"`
}
type PasswordsPasswordsexistingpasswordresetParams struct {
	Email                  string         `json:"email,omitempty"`
	ExistingPassword       string         `json:"existing_password,omitempty"`
	NewPassword            string         `json:"new_password,omitempty"`
	SessionToken           string         `json:"session_token,omitempty"`
	SessionDurationMinutes int32          `json:"session_duration_minutes,omitempty"`
	SessionJwt             string         `json:"session_jwt,omitempty"`
	SessionCustomClaims    map[string]any `json:"session_custom_claims,omitempty"`
}
type PasswordsPasswordssessionresetParams struct {
	Password     string `json:"password,omitempty"`
	SessionToken string `json:"session_token,omitempty"`
	SessionJwt   string `json:"session_jwt,omitempty"`
}
type PasswordsMigrateParams struct {
	Email             string         `json:"email,omitempty"`
	Hash              string         `json:"hash,omitempty"`
	HashType          string         `json:"hash_type,omitempty"`
	Md5Config         MD5Config      `json:"md_5_config,omitempty"`
	Argon2Config      Argon2Config   `json:"argon_2_config,omitempty"`
	Sha1Config        SHA1Config     `json:"sha_1_config,omitempty"`
	ScryptConfig      ScryptConfig   `json:"scrypt_config,omitempty"`
	FirstName         string         `json:"first_name,omitempty"`
	LastName          string         `json:"last_name,omitempty"`
	TrustedMetadata   map[string]any `json:"trusted_metadata,omitempty"`
	UntrustedMetadata map[string]any `json:"untrusted_metadata,omitempty"`
	SetEmailVerified  bool           `json:"set_email_verified,omitempty"`
}
type PasswordsPasswordsstrengthcheckParams struct {
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

type PasswordsCreateResponse struct {
	RequestID    string  `json:"request_id,omitempty"`
	UserID       string  `json:"user_id,omitempty"`
	EmailID      string  `json:"email_id,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
	Session      Session `json:"session,omitempty"`
	SessionJwt   string  `json:"session_jwt,omitempty"`
	User         User    `json:"user,omitempty"`
}
type PasswordsAuthenticateResponse struct {
	RequestID    string  `json:"request_id,omitempty"`
	UserID       string  `json:"user_id,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
	Session      Session `json:"session,omitempty"`
	SessionJwt   string  `json:"session_jwt,omitempty"`
	User         User    `json:"user,omitempty"`
}
type PasswordsPasswordsemailresetstartResponse struct {
	RequestID string `json:"request_id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	EmailID   string `json:"email_id,omitempty"`
}
type PasswordsPasswordsemailresetResponse struct {
	RequestID    string  `json:"request_id,omitempty"`
	UserID       string  `json:"user_id,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
	Session      Session `json:"session,omitempty"`
	SessionJwt   string  `json:"session_jwt,omitempty"`
	User         User    `json:"user,omitempty"`
}
type PasswordsPasswordsexistingpasswordresetResponse struct {
	RequestID    string  `json:"request_id,omitempty"`
	UserID       string  `json:"user_id,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
	Session      Session `json:"session,omitempty"`
	SessionJwt   string  `json:"session_jwt,omitempty"`
	User         User    `json:"user,omitempty"`
}
type PasswordsPasswordssessionresetResponse struct {
	RequestID string  `json:"request_id,omitempty"`
	UserID    string  `json:"user_id,omitempty"`
	Session   Session `json:"session,omitempty"`
	User      User    `json:"user,omitempty"`
}
type PasswordsMigrateResponse struct {
	RequestID   string `json:"request_id,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	EmailID     string `json:"email_id,omitempty"`
	UserCreated bool   `json:"user_created,omitempty"`
	User        User   `json:"user,omitempty"`
}
type PasswordsPasswordsstrengthcheckResponse struct {
	RequestID               string   `json:"request_id,omitempty"`
	ValidPassword           bool     `json:"valid_password,omitempty"`
	Score                   int32    `json:"score,omitempty"`
	BreachedPassword        bool     `json:"breached_password,omitempty"`
	Feedback                Feedback `json:"feedback,omitempty"`
	StrengthPolicy          string   `json:"strength_policy,omitempty"`
	BreachDetectionOnCreate bool     `json:"breach_detection_on_create,omitempty"`
}
