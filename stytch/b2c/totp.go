package b2c

type TOTP struct {
	TOTPID        string   `json:"totp_id,omitempty"`
	Verified      bool     `json:"verified,omitempty"`
	RecoveryCodes []string `json:"recovery_codes,omitempty"`
}

type TOTPsCreateParams struct {
	UserID            string `json:"user_id"`
	ExpirationMinutes int32  `json:"expiration_minutes,omitempty"`
}

type TOTPsCreateResponse struct {
	RequestID     string   `json:"request_id,omitempty"`
	StatusCode    int      `json:"status_code,omitempty"`
	Secret        string   `json:"secret,omitempty"`
	TOTPID        string   `json:"totp_id,omitempty"`
	QRCode        string   `json:"qr_code,omitempty"`
	RecoveryCodes []string `json:"recovery_codes,omitempty"`
	User          User     `json:"user,omitempty"`
	UserID        string   `json:"user_id,omitempty"`
}

type TOTPsAuthenticateParams struct {
	UserID                 string                 `json:"user_id"`
	TOTPCode               string                 `json:"totp_code"`
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
}

type TOTPsAuthenticateResponse struct {
	RequestID    string  `json:"request_id,omitempty"`
	StatusCode   int     `json:"status_code,omitempty"`
	UserID       string  `json:"user_id,omitempty"`
	TOTPID       string  `json:"totp_id,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
	SessionJWT   string  `json:"session_jwt,omitempty"`
	Session      Session `json:"session,omitempty"`
	User         User    `json:"user,omitempty"`
}

type TOTPsRecoveryCodesParams struct {
	UserID string `json:"user_id"`
}

type TOTPsRecoveryCodesResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	TOTPs      []TOTP `json:"totps,omitempty"`
}

type TOTPsRecoverParams struct {
	UserID                 string                 `json:"user_id"`
	RecoveryCode           string                 `json:"recovery_code"`
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
}

type TOTPsRecoverResponse struct {
	RequestID    string  `json:"request_id,omitempty"`
	StatusCode   int     `json:"status_code,omitempty"`
	UserID       string  `json:"user_id,omitempty"`
	TOTPID       string  `json:"totp_id,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
	SessionJWT   string  `json:"session_jwt,omitempty"`
	Session      Session `json:"session,omitempty"`
	User         User    `json:"user,omitempty"`
}
