package stytch

type SessionsGetParams struct {
	UserID string `json:"user_id"`
}

type SessionsGetResponse struct {
	RequestID string    `json:"request_id,omitempty"`
	Sessions  []Session `json:"sessions,omitempty"`
}

type SessionsAuthenticateParams struct {
	SessionToken           string `json:"session_token,omitempty"`
	SessionDurationMinutes int32  `json:"session_duration_minutes,omitempty"`
}

type SessionsAuthenticateResponse struct {
	RequestID    string  `json:"request_id,omitempty"`
	Session      Session `json:"session,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
}

type SessionsRevokeParams struct {
	SessionID    string `json:"session_id,omitempty"`
	SessionToken string `json:"session_token,omitempty"`
}

type SessionsRevokeResponse struct {
	RequestID string `json:"request_id,omitempty"`
}
