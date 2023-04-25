package b2c

type WebAuthnRegisterStartParams struct {
	UserID            string `json:"user_id,omitempty"`
	Domain            string `json:"domain,omitempty"`
	UserAgent         string `json:"user_agent,omitempty"`
	AuthenticatorType string `json:"authenticator_type,omitempty"`
}

type WebAuthnRegisterStartResponse struct {
	RequestID                          string `json:"request_id,omitempty"`
	StatusCode                         int    `json:"status_code,omitempty"`
	UserID                             string `json:"user_id,omitempty"`
	PublicKeyCredentialCreationOptions string `json:"public_key_credential_creation_options,omitempty"`
}

type WebAuthnRegisterParams struct {
	UserID              string `json:"user_id,omitempty"`
	PublicKeyCredential string `json:"public_key_credential,omitempty"`
}

type WebAuthnRegisterResponse struct {
	RequestID              string `json:"request_id,omitempty"`
	StatusCode             int    `json:"status_code,omitempty"`
	UserID                 string `json:"user_id,omitempty"`
	WebAuthnRegistrationID string `json:"webauthn_registration_id,omitempty"`
}

type WebAuthnAuthenticateStartParams struct {
	UserID string `json:"user_id,omitempty"`
	Domain string `json:"domain,omitempty"`
}

type WebAuthnAuthenticateStartResponse struct {
	RequestID                         string `json:"request_id,omitempty"`
	StatusCode                        int    `json:"status_code,omitempty"`
	UserID                            string `json:"user_id,omitempty"`
	PublicKeyCredentialRequestOptions string `json:"public_key_credential_request_options,omitempty"`
}

type WebAuthnAuthenticateParams struct {
	PublicKeyCredential    string                 `json:"public_key_credential,omitempty"`
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
}

type WebAuthnAuthenticateResponse struct {
	RequestID              string  `json:"request_id,omitempty"`
	StatusCode             int     `json:"status_code,omitempty"`
	UserID                 string  `json:"user_id,omitempty"`
	WebAuthnRegistrationID string  `json:"webauthn_registration_id,omitempty"`
	SessionToken           string  `json:"session_token,omitempty"`
	SessionJWT             string  `json:"session_jwt,omitempty"`
	Session                Session `json:"session,omitempty"`
	User                   User    `json:"user,omitempty"`
}
