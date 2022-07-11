package stytch

type PasswordsCreateParams struct {
	Email                  string `json:"email"`
	Password               string `json:"password"`
	SessionDurationMinutes int32  `json:"session_duration_minutes,omitempty"`
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
	Email                  string `json:"email"`
	Password               string `json:"password"`
	SessionToken           string `json:"session_token,omitempty"`
	SessionJWT             string `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32  `json:"session_duration_minutes,omitempty"`
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
	Email       string `json:"email"`
	Hash        string `json:"hash"`
	HashType    string `json:"hash_type"`
	PrependSalt string `json:"prepend_salt,omitempty"`
	AppendSalt  string `json:"append_salt,omitempty"`
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
	LoginExpirationMinutes         int32      `json:"login_expiration_minutes,omitempty"`
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
	Token                  string     `json:"token,omitempty"`
	Password               string     `json:"password,omitempty"`
	SessionToken           string     `json:"session_token,omitempty"`
	SessionJWT             string     `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32      `json:"session_duration_minutes,omitempty"`
	Options                Options    `json:"options,omitempty"`
	Attributes             Attributes `json:"attributes,omitempty"`
	CodeVerifier           string     `json:"code_verifier,omitempty"`
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