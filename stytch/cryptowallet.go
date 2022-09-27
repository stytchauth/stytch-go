package stytch

type CryptoWalletAuthenticateStartParams struct {
	CryptoWalletAddress string `json:"crypto_wallet_address,omitempty"`
	CryptoWalletType    string `json:"crypto_wallet_type,omitempty"`
	UserID              string `json:"user_id,omitempty"`
	SessionToken        string `json:"session_token,omitempty"`
	SessionJWT          string `json:"session_jwt,omitempty"`
}

type CryptoWalletAuthenticateStartResponse struct {
	RequestID   string `json:"request_id,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	Challenge   string `json:"challenge,omitempty"`
	UserCreated bool   `json:"user_created,omitempty"`
}

type CryptoWalletAuthenticateParams struct {
	CryptoWalletAddress    string                 `json:"crypto_wallet_address,omitempty"`
	CryptoWalletType       string                 `json:"crypto_wallet_type,omitempty"`
	Signature              string                 `json:"signature,omitempty"`
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
}

type CryptoWalletAuthenticateResponse struct {
	RequestID    string  `json:"request_id,omitempty"`
	StatusCode   int     `json:"status_code,omitempty"`
	UserID       string  `json:"user_id,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
	SessionJWT   string  `json:"session_jwt,omitempty"`
	Session      Session `json:"session,omitempty"`
	User         User    `json:"user,omitempty"`
}
