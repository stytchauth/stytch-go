package stytch

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

type WebAuthnServiceWebauthnregisterstartParams struct {
	UserID            string `json:"user_id,omitempty"`
	Domain            string `json:"domain,omitempty"`
	UserAgent         string `json:"user_agent,omitempty"`
	AuthenticatorType string `json:"authenticator_type,omitempty"`
}
type WebAuthnServiceWebauthnregisterParams struct {
	UserID              string `json:"user_id,omitempty"`
	PublicKeyCredential string `json:"public_key_credential,omitempty"`
}
type WebAuthnServiceWebauthnauthenticatestartParams struct {
	UserID string `json:"user_id,omitempty"`
	Domain string `json:"domain,omitempty"`
}
type WebAuthnServiceWebauthnauthenticateParams struct {
	PublicKeyCredential    string         `json:"public_key_credential,omitempty"`
	SessionToken           string         `json:"session_token,omitempty"`
	SessionDurationMinutes int32          `json:"session_duration_minutes,omitempty"`
	SessionJwt             string         `json:"session_jwt,omitempty"`
	SessionCustomClaims    map[string]any `json:"session_custom_claims,omitempty"`
}

type WebAuthnServiceWebauthnregisterstartResponse struct {
	RequestID                          string `json:"request_id,omitempty"`
	UserID                             string `json:"user_id,omitempty"`
	PublicKeyCredentialCreationOptions string `json:"public_key_credential_creation_options,omitempty"`
}
type WebAuthnServiceWebauthnregisterResponse struct {
	RequestID              string `json:"request_id,omitempty"`
	UserID                 string `json:"user_id,omitempty"`
	WebauthnRegistrationID string `json:"webauthn_registration_id,omitempty"`
}
type WebAuthnServiceWebauthnauthenticatestartResponse struct {
	RequestID                         string `json:"request_id,omitempty"`
	UserID                            string `json:"user_id,omitempty"`
	PublicKeyCredentialRequestOptions string `json:"public_key_credential_request_options,omitempty"`
}
type WebAuthnServiceWebauthnauthenticateResponse struct {
	RequestID              string  `json:"request_id,omitempty"`
	UserID                 string  `json:"user_id,omitempty"`
	WebauthnRegistrationID string  `json:"webauthn_registration_id,omitempty"`
	SessionToken           string  `json:"session_token,omitempty"`
	Session                Session `json:"session,omitempty"`
	SessionJwt             string  `json:"session_jwt,omitempty"`
	User                   User    `json:"user,omitempty"`
}
