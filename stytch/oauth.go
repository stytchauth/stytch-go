package stytch

type SessionManagementType string

const (
	SessionManagementTypeIDP    SessionManagementType = "idp"
	SessionManagementTypeNone   SessionManagementType = "none"
	SessionManagementTypeStytch SessionManagementType = "stytch"
)

type ProviderType string

const (
	ProviderTypeGoogle ProviderType = "Google"
)

type OAuthAuthenticateParams struct {
	Token                  string                `json:"token,omitempty"`
	SessionManagementType  SessionManagementType `json:"session_management_type,omitempty"`
	SessionToken           string                `json:"session_token,omitempty"`
	SessionDurationMinutes int32                 `json:"session_duration_minutes,omitempty"`
}

type OAuthSessionIdp struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type StytchSession struct {
	Session      Session `json:"session,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
}

type OAuthSession struct {
	Idp           *OAuthSessionIdp `json:"idp,omitempty"`
	StytchSession *StytchSession   `json:"stytch_session,omitempty"`
}

type OAuthAuthenticateResponse struct {
	RequestID       string        `json:"request_id,omitempty"`
	StatusCode      int           `json:"status_code,omitempty"`
	UserID          string        `json:"user_id,omitempty"`
	ProviderSubject string        `json:"provider_subject,omitempty"`
	ProviderType    ProviderType  `json:"provider_type,omitempty"`
	Session         *OAuthSession `json:"session,omitempty"`
}
