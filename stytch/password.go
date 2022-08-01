package stytch

type PasswordsCreateParams struct {
	Email                  string                 `json:"email"`
	Password               string                 `json:"password"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
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
	RequestID        string   `json:"request_id,omitempty"`
	StatusCode       int      `json:"status_code,omitempty"`
	ValidPassword    bool     `json:"valid_password,omitempty"`
	Score            int      `json:"score,omitempty"`
	BreachedPassword bool     `json:"breached_password,omitempty"`
	Feedback         Feedback `json:"feedback,omitempty"`
}

type Feedback struct {
	Warning     string   `json:"warning,omitempty"`
	Suggestions []string `json:"suggestions,omitempty"`
}

type PasswordsMigrateParams struct {
	Email        string       `json:"email"`
	Hash         string       `json:"hash"`
	HashType     HashType     `json:"hash_type"`
	MD5Config    MD5Config    `json:"md_5_config,omitempty"`
	Argon2Config Argon2Config `json:"argon_2_config,omitempty"`
}

type MD5Config struct {
	PrependSalt string `json:"prepend_salt,omitempty"`
	AppendSalt  string `json:"append_salt,omitempty"`
}

type Argon2Config struct {
	Salt            string `json:"salt"`
	IterationAmount string `json:"iteration_amount"`
	Memory          string `json:"memory"`
	Threads         string `json:"threads"`
	KeyLength       string `json:"key_length"`
}

type PasswordsMigrateResponse struct {
	RequestID   string `json:"request_id,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	EmailID     string `json:"email_id,omitempty"`
	UserCreated bool   `json:"user_created,omitempty"`
}

type HashType string

const (
	HashTypeBcrypt   HashType = "bcrypt"
	HashTypeMD5      HashType = "md_5"
	HashTypeArgon2I  HashType = "argon_2i"
	HashTypeArgon2ID HashType = "argon_2id"
)

// PASSWORD - EMAIL
type PasswordEmailResetStartParams struct {
	Email                          string     `json:"email"`
	LoginRedirectURL               string     `json:"login_redirect_url,omitempty"`
	ResetPasswordRedirectURL       string     `json:"reset_password_redirect_url,omitempty"`
	ResetPasswordExpirationMinutes int32      `json:"reset_password_expiration_minutes,omitempty"`
	Attributes                     Attributes `json:"attributes,omitempty"`
	CodeChallenge                  string     `json:"code_challenge,omitempty"`
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
