package stytch

type SessionManagementType string

const (
	SessionManagementTypeIdP    SessionManagementType = "idp"
	SessionManagementTypeNone   SessionManagementType = "none"
	SessionManagementTypeStytch SessionManagementType = "stytch"
)

type ProviderType string

const (
	ProviderTypeGoogle ProviderType = "Google"
)

type OAuthAuthenticateParams struct {
	Token                 string                `json:"token,omitempty"`
	SessionManagementType SessionManagementType `json:"session_management_type,omitempty"`
}

type OAuthSessionIdp struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type OAuthSession struct {
	Idp *OAuthSessionIdp `json:"idp,omitempty"`
}

type OAuthAuthenticateResponse struct {
	RequestID       string        `json:"request_id,omitempty"`
	StatusCode      int           `json:"status_code,omitempty"`
	UserID          string        `json:"user_id,omitempty"`
	ProviderSubject string        `json:"provider_subject,omitempty"`
	ProviderType    ProviderType  `json:"provider_type,omitempty"`
	Session         *OAuthSession `json:"session,omitempty"`
}
