package stytch

import "time"

type ProviderType string

const (
	ProviderTypeAmazon    ProviderType = "Amazon"
	ProviderTypeApple     ProviderType = "Apple"
	ProviderTypeBitbucket ProviderType = "Bitbucket"
	ProviderTypeCoinbase  ProviderType = "Coinbase"
	ProviderTypeDiscord   ProviderType = "Discord"
	ProviderTypeFacebook  ProviderType = "Facebook"
	ProviderTypeFigma     ProviderType = "Figma"
	ProviderTypeGithub    ProviderType = "Github"
	ProviderTypeGitlab    ProviderType = "GitLab"
	ProviderTypeGoogle    ProviderType = "Google"
	ProviderTypeLinkedIn  ProviderType = "LinkedIn"
	ProviderTypeMicrosoft ProviderType = "Microsoft"
	ProviderTypeSlack     ProviderType = "Slack"
	ProviderTypeSnapchat  ProviderType = "Snapchat"
	ProviderTypeSpotify   ProviderType = "Spotify"
	ProviderTypeTikTok    ProviderType = "TikTok"
	ProviderTypeTwitch    ProviderType = "Twitch"
	ProviderTypeTwitter   ProviderType = "Twitter"
)

type OAuthAuthenticateParams struct {
	Token                  string                 `json:"token,omitempty"`
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
	CodeVerifier           string                 `json:"code_verifier,omitempty"`
}

type OAuthAuthenticateResponse struct {
	RequestID               string         `json:"request_id,omitempty"`
	StatusCode              int            `json:"status_code,omitempty"`
	UserID                  string         `json:"user_id,omitempty"`
	OAuthUserRegistrationID string         `json:"oauth_user_registration_id,omitempty"`
	ProviderSubject         string         `json:"provider_subject,omitempty"`
	ProviderType            ProviderType   `json:"provider_type,omitempty"`
	Session                 *Session       `json:"session,omitempty"`
	SessionToken            string         `json:"session_token,omitempty"`
	SessionJWT              string         `json:"session_jwt,omitempty"`
	ProviderValues          ProviderValues `json:"provider_values,omitempty"`
	ResetSessions           bool           `json:"reset_sessions,omitempty"`
}

type ProviderValues struct {
	AccessToken  string     `json:"access_token,omitempty"`
	RefreshToken string     `json:"refresh_token,omitempty"`
	IDToken      string     `json:"id_token,omitempty"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty"`
	Scopes       []string   `json:"scopes,omitempty"`
}
