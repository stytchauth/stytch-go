package b2c

import "github.com/golang-jwt/jwt/v5"

/*
 * Structure for the custom type Attributes
 */
type Attributes struct {
	// The ip address of the user.
	IPAddress string `json:"ip_address,omitempty"`
	// The user agent of the user.
	UserAgent string `json:"user_agent,omitempty"`
}

/*
 * Structure for the custom type Options
 */
type Options struct {
	// Require that the ip address the magic link was requested from
	// matches the ip address it's clicked from.
	IPMatchRequired bool `json:"ip_match_required,omitempty"`
	// Require that the user agent the magic link was requested from
	// matches the user agent it's clicked from.
	UserAgentMatchRequired bool `json:"user_agent_match_required,omitempty"`
}

type EmailFactor struct {
	EmailID      string `json:"email_id,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
}

type PhoneNumberFactor struct {
	PhoneID     string `json:"phone_id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type GoogleOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type MicrosoftOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type AppleOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type GithubOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type GitLabOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type FacebookOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type AmazonOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type BitbucketOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type CoinbaseOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type DiscordOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type FigmaOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type LinkedInOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type SalesforceOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type SlackOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type SnapchatOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type SpotifyOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type TikTokOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type TwitchOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type TwitterOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type WebAuthnFactor struct {
	WebAuthnRegistrationID string `json:"webauthn_registration_id,omitempty"`
	Domain                 string `json:"domain,omitempty"`
	UserAgent              string `json:"user_agent,omitempty"`
}

type BiometricFactor struct {
	BiometricRegistrationID string `json:"biometric_registration_id,omitempty"`
}

type AuthenticatorAppFactor struct {
	TOTPID string `json:"totp_id,omitempty"`
}

type RecoveryCodeFactor struct {
	TOTPRecoveryCodeID string `json:"totp_recovery_code_id,omitempty"`
}

type CryptoWalletFactor struct {
	CryptoWalletID      string `json:"crypto_wallet_id,omitempty"`
	CryptoWalletAddress string `json:"crypto_wallet_address,omitempty"`
	CryptoWalletType    string `json:"crypto_wallet_type,omitempty"`
}

type AuthenticationFactor struct {
	Type                string `json:"type,omitempty"`
	DeliveryMethod      string `json:"delivery_method,omitempty"`
	LastAuthenticatedAt string `json:"last_authenticated_at,omitempty"`

	EmailFactor            EmailFactor            `json:"email_factor,omitempty"`
	PhoneNumberFactor      PhoneNumberFactor      `json:"phone_number_factor,omitempty"`
	GoogleOAuthFactor      GoogleOAuthFactor      `json:"google_oauth_factor,omitempty"`
	MicrosoftOAuthFactor   MicrosoftOAuthFactor   `json:"microsoft_oauth_factor,omitempty"`
	AppleOAuthFactor       AppleOAuthFactor       `json:"apple_oauth_factor,omitempty"`
	GithubOAuthFactor      GithubOAuthFactor      `json:"github_oauth_factor,omitempty"`
	GitLabOAuthFactor      GitLabOAuthFactor      `json:"gitlab_oauth_factor,omitempty"`
	FacebookOAuthFactor    FacebookOAuthFactor    `json:"facebook_oauth_factor,omitempty"`
	AmazonOAuthFactor      AmazonOAuthFactor      `json:"amazon_oauth_factor,omitempty"`
	BitbucketOAuthFactor   BitbucketOAuthFactor   `json:"bitbucket_oauth_factor,omitempty"`
	CoinbaseOAuthFactor    CoinbaseOAuthFactor    `json:"coinbase_oauth_factor,omitempty"`
	DiscordOAuthFactor     DiscordOAuthFactor     `json:"discord_oauth_factor,omitempty"`
	FigmaOAuthFactor       FigmaOAuthFactor       `json:"figma_oauth_factor,omitempty"`
	LinkedInOAuthFactor    LinkedInOAuthFactor    `json:"linkedin_oauth_factor,omitempty"`
	SalesforceOAuthFactor  SalesforceOAuthFactor  `json:"salesforce_oauth_factor,omitempty"`
	SlackOAuthFactor       SlackOAuthFactor       `json:"slack_oauth_factor,omitempty"`
	SnapchatOAuthFactor    SnapchatOAuthFactor    `json:"snapchat_oauth_factor,omitempty"`
	SpotifyOAuthFactor     SpotifyOAuthFactor     `json:"spotify_oauth_factor,omitempty"`
	TikTokOAuthFactor      TikTokOAuthFactor      `json:"tiktok_oauth_factor,omitempty"`
	TwitchOAuthFactor      TwitchOAuthFactor      `json:"twitch_oauth_factor,omitempty"`
	TwitterOAuthFactor     TwitterOAuthFactor     `json:"twitter_oauth_factor,omitempty"`
	WebAuthnFactor         WebAuthnFactor         `json:"webauthn_factor,omitempty"`
	BiometricFactor        BiometricFactor        `json:"biometric_factor,omitempty"`
	AuthenticatorAppFactor AuthenticatorAppFactor `json:"authenticator_app_factor,omitempty"`
	RecoveryCodeFactor     RecoveryCodeFactor     `json:"recovery_code_factor,omitempty"`
	CryptoWalletFactor     CryptoWalletFactor     `json:"crypto_wallet_factor,omitempty"`
}

type JWTAuthenticationFactor struct {
	Type                string           `json:"type,omitempty"`
	DeliveryMethod      string           `json:"delivery_method,omitempty"`
	LastAuthenticatedAt *jwt.NumericDate `json:"last_authenticated_at,omitempty"`

	EmailFactor            EmailFactor            `json:"email_factor,omitempty"`
	PhoneNumberFactor      PhoneNumberFactor      `json:"phone_number_factor,omitempty"`
	GoogleOAuthFactor      GoogleOAuthFactor      `json:"google_oauth_factor,omitempty"`
	MicrosoftOAuthFactor   MicrosoftOAuthFactor   `json:"microsoft_oauth_factor,omitempty"`
	AppleOAuthFactor       AppleOAuthFactor       `json:"apple_oauth_factor,omitempty"`
	GithubOAuthFactor      GithubOAuthFactor      `json:"github_oauth_factor,omitempty"`
	FacebookOAuthFactor    FacebookOAuthFactor    `json:"facebook_oauth_factor,omitempty"`
	AmazonOAuthFactor      AmazonOAuthFactor      `json:"amazon_oauth_factor,omitempty"`
	BitbucketOAuthFactor   BitbucketOAuthFactor   `json:"bitbucket_oauth_factor,omitempty"`
	CoinbaseOAuthFactor    CoinbaseOAuthFactor    `json:"coinbase_oauth_factor,omitempty"`
	DiscordOAuthFactor     DiscordOAuthFactor     `json:"discord_oauth_factor,omitempty"`
	FigmaOAuthFactor       FigmaOAuthFactor       `json:"figma_oauth_factor,omitempty"`
	LinkedInOAuthFactor    LinkedInOAuthFactor    `json:"linkedin_oauth_factor,omitempty"`
	SalesforceOAuthFactor  SalesforceOAuthFactor  `json:"salesforce_oauth_factor,omitempty"`
	SlackOAuthFactor       SlackOAuthFactor       `json:"slack_oauth_factor,omitempty"`
	SnapchatOAuthFactor    SnapchatOAuthFactor    `json:"snapchat_oauth_factor,omitempty"`
	SpotifyOAuthFactor     SpotifyOAuthFactor     `json:"spotify_oauth_factor,omitempty"`
	TikTokOAuthFactor      TikTokOAuthFactor      `json:"tiktok_oauth_factor,omitempty"`
	TwitchOAuthFactor      TwitchOAuthFactor      `json:"twitch_oauth_factor,omitempty"`
	TwitterOAuthFactor     TwitterOAuthFactor     `json:"twitter_oauth_factor,omitempty"`
	WebAuthnFactor         WebAuthnFactor         `json:"webauthn_factor,omitempty"`
	BiometricFactor        BiometricFactor        `json:"biometric_factor,omitempty"`
	AuthenticatorAppFactor AuthenticatorAppFactor `json:"authenticator_app_factor,omitempty"`
	RecoveryCodeFactor     RecoveryCodeFactor     `json:"recovery_code_factor,omitempty"`
	CryptoWalletFactor     CryptoWalletFactor     `json:"crypto_wallet_factor,omitempty"`
}

/*
 * Structure for the custom type Session
 */
type Session struct {
	SessionID             string                  `json:"session_id,omitempty"`
	UserID                string                  `json:"user_id,omitempty"`
	StartedAt             string                  `json:"started_at,omitempty"`
	LastAccessedAt        string                  `json:"last_accessed_at,omitempty"`
	ExpiresAt             string                  `json:"expires_at,omitempty"`
	Attributes            Attributes              `json:"attributes,omitempty"`
	AuthenticationFactors []*AuthenticationFactor `json:"authentication_factors,omitempty"`
	CustomClaims          interface{}             `json:"custom_claims"`
}

// SessionWrapper wraps a session object with a custom_claims field
// so that we can unmarshal custom claims from authenticate responses
type SessionWrapper struct {
	Session ClaimsWrapper `json:"session"`
}

type ClaimsWrapper struct {
	Claims interface{} `json:"custom_claims"`
}
