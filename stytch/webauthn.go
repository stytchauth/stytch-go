package stytch

type RegisterStartParams struct {
	UserID            string `json:"user_id,omitempty"`
	Domain            string `json:"domain,omitempty"`
	UserAgent         string `json:"user_agent,omitempty"`
	AuthenticatorType string `json:"authenticator_type,omitempty"`
}

type RegisterStartResponse struct {
	RequestID                          string `json:"request_id,omitempty"`
	StatusCode                         int    `json:"status_code,omitempty"`
	UserID                             string `json:"user_id,omitempty"`
	PublicKeyCredentialCreationOptions string `json:"public_key_credential_creation_options,omitempty"`
}

type RegisterParams struct {
	UserID              string `json:"user_id,omitempty"`
	PublicKeyCredential string `json:"public_key_credential,omitempty"`
}

type RegisterResponse struct {
	RequestID              string `json:"request_id,omitempty"`
	StatusCode             int    `json:"status_code,omitempty"`
	UserID                 string `json:"user_id,omitempty"`
	WebAuthnRegistrationID string `json:"webauthn_registration_id,omitempty"`
}

type AuthenticateStartParams struct {
	UserID string `json:"user_id,omitempty"`
	Domain string `json:"domain,omitempty"`
}

type AuthenticateStartResponse struct {
	RequestID                         string `json:"request_id,omitempty"`
	StatusCode                        int    `json:"status_code,omitempty"`
	UserID                            string `json:"user_id,omitempty"`
	PublicKeyCredentialRequestOptions string `json:"public_key_credential_request_options,omitempty"`
}

type AuthenticateParams struct {
	PublicKeyCredential    string `json:"public_key_credential,omitempty"`
	SessionToken           string `json:"session_token,omitempty"`
	SessionDurationMinutes int32  `json:"session_duration_minutes,omitempty"`
}

type AuthenticateResponse struct {
	RequestID              string  `json:"request_id,omitempty"`
	StatusCode             int     `json:"status_code,omitempty"`
	UserID                 string  `json:"user_id,omitempty"`
	WebAuthnRegistrationID string  `json:"webauthn_registration_id,omitempty"`
	SessionToken           string  `json:"session_token,omitempty"`
	Session                Session `json:"session,omitempty"`
}
