package sessions

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/attribute"
	"github.com/stytchauth/stytch-go/v8/stytch/b2c/users"
)

type AmazonOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}
type AppleOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

// AuthenticateParams: Request type for `Authenticate`.
// Fields:
//
//   - SessionToken: The session token to authenticate.
//
//   - SessionDurationMinutes: Set the session lifetime to be this many minutes from now; minimum of 5 and
//     a maximum of 527040 minutes (366 days). Note that a successful authentication will continue to extend
//     the session this many minutes.
//
//   - SessionJWT: The JWT to authenticate. You may provide a JWT that has expired according to its `exp`
//     claim and needs to be refreshed. If the signature is valid and the underlying session is still active
//     then Stytch will return a new JWT.
//
//   - SessionCustomClaims: Add a custom claims map to the Session being authenticated. Claims are only
//     created if a Session is initialized by providing a value in `session_duration_minutes`. Claims will be
//     included on the Session object and in the JWT. To update a key in an existing Session, supply a new
//     value. To delete a key, supply a null value.
//
//     Custom claims made with reserved claims ("iss", "sub", "aud", "exp", "nbf", "iat", "jti") will be
//     ignored. Total custom claims size cannot exceed four kilobytes.
type AuthenticateParams struct {
	SessionToken           string         `json:"session_token,omitempty"`
	SessionDurationMinutes int32          `json:"session_duration_minutes,omitempty"`
	SessionJWT             string         `json:"session_jwt,omitempty"`
	SessionCustomClaims    map[string]any `json:"session_custom_claims,omitempty"`
}
type AuthenticationFactor struct {
	Type                      AuthenticationFactorType           `json:"type,omitempty"`
	DeliveryMethod            AuthenticationFactorDeliveryMethod `json:"delivery_method,omitempty"`
	LastAuthenticatedAt       *time.Time                         `json:"last_authenticated_at,omitempty"`
	CreatedAt                 *time.Time                         `json:"created_at,omitempty"`
	UpdatedAt                 *time.Time                         `json:"updated_at,omitempty"`
	EmailFactor               EmailFactor                        `json:"email_factor,omitempty"`
	PhoneNumberFactor         PhoneNumberFactor                  `json:"phone_number_factor,omitempty"`
	GoogleOauthFactor         GoogleOAuthFactor                  `json:"google_oauth_factor,omitempty"`
	MicrosoftOauthFactor      MicrosoftOAuthFactor               `json:"microsoft_oauth_factor,omitempty"`
	AppleOauthFactor          AppleOAuthFactor                   `json:"apple_oauth_factor,omitempty"`
	WebauthnFactor            WebAuthnFactor                     `json:"webauthn_factor,omitempty"`
	AuthenticatorAppFactor    AuthenticatorAppFactor             `json:"authenticator_app_factor,omitempty"`
	GithubOauthFactor         GithubOAuthFactor                  `json:"github_oauth_factor,omitempty"`
	RecoveryCodeFactor        RecoveryCodeFactor                 `json:"recovery_code_factor,omitempty"`
	FacebookOauthFactor       FacebookOAuthFactor                `json:"facebook_oauth_factor,omitempty"`
	CryptoWalletFactor        CryptoWalletFactor                 `json:"crypto_wallet_factor,omitempty"`
	AmazonOauthFactor         AmazonOAuthFactor                  `json:"amazon_oauth_factor,omitempty"`
	BitbucketOauthFactor      BitbucketOAuthFactor               `json:"bitbucket_oauth_factor,omitempty"`
	CoinbaseOauthFactor       CoinbaseOAuthFactor                `json:"coinbase_oauth_factor,omitempty"`
	DiscordOauthFactor        DiscordOAuthFactor                 `json:"discord_oauth_factor,omitempty"`
	FigmaOauthFactor          FigmaOAuthFactor                   `json:"figma_oauth_factor,omitempty"`
	GitLabOauthFactor         GitLabOAuthFactor                  `json:"git_lab_oauth_factor,omitempty"`
	InstagramOauthFactor      InstagramOAuthFactor               `json:"instagram_oauth_factor,omitempty"`
	LinkedInOauthFactor       LinkedInOAuthFactor                `json:"linked_in_oauth_factor,omitempty"`
	ShopifyOauthFactor        ShopifyOAuthFactor                 `json:"shopify_oauth_factor,omitempty"`
	SlackOauthFactor          SlackOAuthFactor                   `json:"slack_oauth_factor,omitempty"`
	SnapchatOauthFactor       SnapchatOAuthFactor                `json:"snapchat_oauth_factor,omitempty"`
	SpotifyOauthFactor        SpotifyOAuthFactor                 `json:"spotify_oauth_factor,omitempty"`
	SteamOauthFactor          SteamOAuthFactor                   `json:"steam_oauth_factor,omitempty"`
	TikTokOauthFactor         TikTokOAuthFactor                  `json:"tik_tok_oauth_factor,omitempty"`
	TwitchOauthFactor         TwitchOAuthFactor                  `json:"twitch_oauth_factor,omitempty"`
	TwitterOauthFactor        TwitterOAuthFactor                 `json:"twitter_oauth_factor,omitempty"`
	EmbeddableMagicLinkFactor EmbeddableMagicLinkFactor          `json:"embeddable_magic_link_factor,omitempty"`
	BiometricFactor           BiometricFactor                    `json:"biometric_factor,omitempty"`
	SamlSsoFactor             SAMLSSOFactor                      `json:"saml_sso_factor,omitempty"`
	OidcSsoFactor             OIDCSSOFactor                      `json:"oidc_sso_factor,omitempty"`
	SalesforceOauthFactor     SalesforceOAuthFactor              `json:"salesforce_oauth_factor,omitempty"`
}
type AuthenticatorAppFactor struct {
	TotpID string `json:"totp_id,omitempty"`
}
type BiometricFactor struct {
	BiometricRegistrationID string `json:"biometric_registration_id,omitempty"`
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
type CryptoWalletFactor struct {
	CryptoWalletID      string `json:"crypto_wallet_id,omitempty"`
	CryptoWalletAddress string `json:"crypto_wallet_address,omitempty"`
	CryptoWalletType    string `json:"crypto_wallet_type,omitempty"`
}
type DiscordOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}
type EmailFactor struct {
	EmailID      string `json:"email_id,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
}
type EmbeddableMagicLinkFactor struct {
	EmbeddedID string `json:"embedded_id,omitempty"`
}
type FacebookOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}
type FigmaOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

// GetParams: Request type for `Get`.
// Fields:
//   - UserID: The `user_id` to get active Sessions for.
type GetParams struct {
	UserID string `json:"user_id,omitempty"`
}
type GitLabOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}
type GithubOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}
type GoogleOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}
type InstagramOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}
type JWK struct {
	Kty     string   `json:"kty,omitempty"`
	Use     string   `json:"use,omitempty"`
	KeyOps  []string `json:"key_ops,omitempty"`
	Alg     string   `json:"alg,omitempty"`
	Kid     string   `json:"kid,omitempty"`
	X5c     []string `json:"x5c,omitempty"`
	X5tS256 string   `json:"x5t_s256,omitempty"`
	N       string   `json:"n,omitempty"`
	E       string   `json:"e,omitempty"`
}

// JWKSParams: Request type for `GetJWKS`.
// Fields:
//   - ProjectID: The `project_id` to get the JWKS for.
type JWKSParams struct {
	ProjectID string `json:"project_id,omitempty"`
}
type LinkedInOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}
type MicrosoftOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}
type OIDCSSOFactor struct {
	ID         string `json:"id,omitempty"`
	ProviderID string `json:"provider_id,omitempty"`
	ExternalID string `json:"external_id,omitempty"`
}
type PhoneNumberFactor struct {
	PhoneID     string `json:"phone_id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
type RecoveryCodeFactor struct {
	TotpRecoveryCodeID string `json:"totp_recovery_code_id,omitempty"`
}

// RevokeParams: Request type for `Revoke`.
// Fields:
//   - SessionID: The `session_id` to revoke.
//   - SessionToken: The session token to revoke.
//   - SessionJWT: A JWT for the session to revoke.
type RevokeParams struct {
	SessionID    string `json:"session_id,omitempty"`
	SessionToken string `json:"session_token,omitempty"`
	SessionJWT   string `json:"session_jwt,omitempty"`
}
type SAMLSSOFactor struct {
	ID         string `json:"id,omitempty"`
	ProviderID string `json:"provider_id,omitempty"`
	ExternalID string `json:"external_id,omitempty"`
}
type SalesforceOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

// Session: Fields:
//   - SessionID: A unique identifier for a specific Session.
//   - UserID: The unique ID of the affected User.
//   - AuthenticationFactors: An array of different authentication factors that have initiated a Session.
//   - StartedAt: The timestamp of the Session's creation. Values conform to the RFC 3339 standard and are
//     expressed in UTC, e.g. `2021-12-29T12:33:09Z`.
//   - LastAccessedAt: The timestamp when the Session was last accessed. Values conform to the RFC 3339
//     standard and are expressed in UTC, e.g. `2021-12-29T12:33:09Z`.
//   - ExpiresAt: The timestamp when the Session expires. Values conform to the RFC 3339 standard and are
//     expressed in UTC, e.g. `2021-12-29T12:33:09Z`.
//   - Attributes: Provided attributes help with fraud detection.
//   - CustomClaims: The custom claims map for a Session. Claims can be added to a session during a
//     Sessions authenticate call.
type Session struct {
	SessionID             string                 `json:"session_id,omitempty"`
	UserID                string                 `json:"user_id,omitempty"`
	AuthenticationFactors []AuthenticationFactor `json:"authentication_factors,omitempty"`
	StartedAt             *time.Time             `json:"started_at,omitempty"`
	LastAccessedAt        *time.Time             `json:"last_accessed_at,omitempty"`
	ExpiresAt             *time.Time             `json:"expires_at,omitempty"`
	Attributes            attribute.Attributes   `json:"attributes,omitempty"`
	CustomClaims          map[string]any         `json:"custom_claims,omitempty"`
}
type ShopifyOAuthFactor struct {
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
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}
type SpotifyOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}
type SteamOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}
type TikTokOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}
type TwitchOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}
type TwitterOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}
type WebAuthnFactor struct {
	WebauthnRegistrationID string `json:"webauthn_registration_id,omitempty"`
	Domain                 string `json:"domain,omitempty"`
	UserAgent              string `json:"user_agent,omitempty"`
}

// AuthenticateResponse: Response type for `Authenticate`.
// Fields:
//
//   - RequestID: Globally unique UUID that is returned with every API call. This value is important to log
//     for debugging purposes; we may ask for this value to help identify a specific API call when helping you
//     debug an issue.
//
//   - SessionToken: A secret token for a given Stytch Session.
//
//   - SessionJWT: The JSON Web Token (JWT) for a given Stytch Session.
//
//   - User: The `user` object affected by this API call. See the [Get user
//     endpoint](https://invalid.dev.stytch.com/docs/api/get-user) for complete response field details.
//
//   - StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
//     patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
//     are server errors.
//
//   - Session: If you initiate a Session, by including `session_duration_minutes` in your authenticate
//     call, you'll receive a full Session object in the response.
//
//     See [GET sessions](https://stytch.com/docs/api/session-get) for complete response fields.
type AuthenticateResponse struct {
	RequestID    string     `json:"request_id,omitempty"`
	SessionToken string     `json:"session_token,omitempty"`
	SessionJWT   string     `json:"session_jwt,omitempty"`
	User         users.User `json:"user,omitempty"`
	StatusCode   int32      `json:"status_code,omitempty"`
	Session      Session    `json:"session,omitempty"`
}

// GetResponse: Response type for `Get`.
// Fields:
//   - RequestID: Globally unique UUID that is returned with every API call. This value is important to log
//     for debugging purposes; we may ask for this value to help identify a specific API call when helping you
//     debug an issue.
//   - Sessions: An array of Session objects.
//   - StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
//     patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
//     are server errors.
type GetResponse struct {
	RequestID  string    `json:"request_id,omitempty"`
	Sessions   []Session `json:"sessions,omitempty"`
	StatusCode int32     `json:"status_code,omitempty"`
}

// JWKSResponse: Response type for `GetJWKS`.
// Fields:
//   - Keys: The JWK
//   - RequestID: Globally unique UUID that is returned with every API call. This value is important to log
//     for debugging purposes; we may ask for this value to help identify a specific API call when helping you
//     debug an issue.
//   - StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
//     patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
//     are server errors.
type JWKSResponse struct {
	Keys       []JWK  `json:"keys,omitempty"`
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int32  `json:"status_code,omitempty"`
}

// RevokeResponse: Response type for `Revoke`.
// Fields:
//   - RequestID: Globally unique UUID that is returned with every API call. This value is important to log
//     for debugging purposes; we may ask for this value to help identify a specific API call when helping you
//     debug an issue.
//   - StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
//     patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
//     are server errors.
type RevokeResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int32  `json:"status_code,omitempty"`
}

type AuthenticationFactorDeliveryMethod string

const (
	AuthenticationFactorDeliveryMethodEmail                AuthenticationFactorDeliveryMethod = "email"
	AuthenticationFactorDeliveryMethodSms                  AuthenticationFactorDeliveryMethod = "sms"
	AuthenticationFactorDeliveryMethodWhatsapp             AuthenticationFactorDeliveryMethod = "whatsapp"
	AuthenticationFactorDeliveryMethodEmbedded             AuthenticationFactorDeliveryMethod = "embedded"
	AuthenticationFactorDeliveryMethodOauthGoogle          AuthenticationFactorDeliveryMethod = "oauth_google"
	AuthenticationFactorDeliveryMethodOauthMicrosoft       AuthenticationFactorDeliveryMethod = "oauth_microsoft"
	AuthenticationFactorDeliveryMethodOauthApple           AuthenticationFactorDeliveryMethod = "oauth_apple"
	AuthenticationFactorDeliveryMethodWebauthnRegistration AuthenticationFactorDeliveryMethod = "webauthn_registration"
	AuthenticationFactorDeliveryMethodAuthenticatorApp     AuthenticationFactorDeliveryMethod = "authenticator_app"
	AuthenticationFactorDeliveryMethodOauthGithub          AuthenticationFactorDeliveryMethod = "oauth_github"
	AuthenticationFactorDeliveryMethodRecoveryCode         AuthenticationFactorDeliveryMethod = "recovery_code"
	AuthenticationFactorDeliveryMethodOauthFacebook        AuthenticationFactorDeliveryMethod = "oauth_facebook"
	AuthenticationFactorDeliveryMethodCryptoWallet         AuthenticationFactorDeliveryMethod = "crypto_wallet"
	AuthenticationFactorDeliveryMethodOauthAmazon          AuthenticationFactorDeliveryMethod = "oauth_amazon"
	AuthenticationFactorDeliveryMethodOauthBitbucket       AuthenticationFactorDeliveryMethod = "oauth_bitbucket"
	AuthenticationFactorDeliveryMethodOauthCoinbase        AuthenticationFactorDeliveryMethod = "oauth_coinbase"
	AuthenticationFactorDeliveryMethodOauthDiscord         AuthenticationFactorDeliveryMethod = "oauth_discord"
	AuthenticationFactorDeliveryMethodOauthFigma           AuthenticationFactorDeliveryMethod = "oauth_figma"
	AuthenticationFactorDeliveryMethodOauthGitlab          AuthenticationFactorDeliveryMethod = "oauth_gitlab"
	AuthenticationFactorDeliveryMethodOauthInstagram       AuthenticationFactorDeliveryMethod = "oauth_instagram"
	AuthenticationFactorDeliveryMethodOauthLinkedin        AuthenticationFactorDeliveryMethod = "oauth_linkedin"
	AuthenticationFactorDeliveryMethodOauthShopify         AuthenticationFactorDeliveryMethod = "oauth_shopify"
	AuthenticationFactorDeliveryMethodOauthSlack           AuthenticationFactorDeliveryMethod = "oauth_slack"
	AuthenticationFactorDeliveryMethodOauthSnapchat        AuthenticationFactorDeliveryMethod = "oauth_snapchat"
	AuthenticationFactorDeliveryMethodOauthSpotify         AuthenticationFactorDeliveryMethod = "oauth_spotify"
	AuthenticationFactorDeliveryMethodOauthSteam           AuthenticationFactorDeliveryMethod = "oauth_steam"
	AuthenticationFactorDeliveryMethodOauthTiktok          AuthenticationFactorDeliveryMethod = "oauth_tiktok"
	AuthenticationFactorDeliveryMethodOauthTwitch          AuthenticationFactorDeliveryMethod = "oauth_twitch"
	AuthenticationFactorDeliveryMethodOauthTwitter         AuthenticationFactorDeliveryMethod = "oauth_twitter"
	AuthenticationFactorDeliveryMethodKnowledge            AuthenticationFactorDeliveryMethod = "knowledge"
	AuthenticationFactorDeliveryMethodBiometric            AuthenticationFactorDeliveryMethod = "biometric"
	AuthenticationFactorDeliveryMethodSsoSaml              AuthenticationFactorDeliveryMethod = "sso_saml"
	AuthenticationFactorDeliveryMethodSsoOidc              AuthenticationFactorDeliveryMethod = "sso_oidc"
	AuthenticationFactorDeliveryMethodOauthSalesforce      AuthenticationFactorDeliveryMethod = "oauth_salesforce"
)

type AuthenticationFactorType string

const (
	AuthenticationFactorTypeMagicLink          AuthenticationFactorType = "magic_link"
	AuthenticationFactorTypeOTP                AuthenticationFactorType = "otp"
	AuthenticationFactorTypeOauth              AuthenticationFactorType = "oauth"
	AuthenticationFactorTypeWebauthn           AuthenticationFactorType = "webauthn"
	AuthenticationFactorTypeTOTP               AuthenticationFactorType = "totp"
	AuthenticationFactorTypeCrypto             AuthenticationFactorType = "crypto"
	AuthenticationFactorTypePassword           AuthenticationFactorType = "password"
	AuthenticationFactorTypeSignatureChallenge AuthenticationFactorType = "signature_challenge"
	AuthenticationFactorTypeSSO                AuthenticationFactorType = "sso"
)

// MANUAL(Types)(TYPES)
// ADDIMPORT: "errors"
// ADDIMPORT: "strings"
// ADDIMPORT: "github.com/golang-jwt/jwt/v5"

var ErrJWTTooOld = errors.New("JWT too old")

type SessionClaim struct {
	ID                    string                 `json:"id"`
	StartedAt             string                 `json:"started_at"`
	LastAccessedAt        string                 `json:"last_accessed_at"`
	ExpiresAt             string                 `json:"expires_at"`
	Attributes            attribute.Attributes   `json:"attributes"`
	AuthenticationFactors []AuthenticationFactor `json:"authentication_factors"`
}

type Claims struct {
	StytchSession SessionClaim `json:"https://stytch.com/session"`
	jwt.RegisteredClaims
}

type ClaimsWrapper struct {
	Claims map[string]any `json:"custom_claims"`
}

type SessionWrapper struct {
	Session ClaimsWrapper `json:"session"`
}

// IsValid returns an error if there is an issuer or audience mismatch in the claims.
//
// Deprecated: JWT claims are validated when the token is parsed. There is no need to call this method.
func (c Claims) IsValid(projectID string) error {
	var errs []error

	if !c.verifyIssuer(projectID) {
		errs = append(errs, jwt.ErrTokenInvalidIssuer)
	}

	if !c.verifyAudience(projectID) {
		errs = append(errs, jwt.ErrTokenInvalidAudience)
	}

	if len(errs) == 0 {
		return nil
	}
	return multiError{errs}
}

type multiError struct {
	errs []error
}

func (me multiError) Error() string {
	var msgs []string
	for _, err := range me.errs {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, ", ")
}

func (me multiError) Is(target error) bool {
	for _, err := range me.errs {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}

func (c *Claims) verifyIssuer(cmp string) bool {
	issuerSplit := strings.Split(c.RegisteredClaims.Issuer, "/")
	return len(issuerSplit) == 2 && issuerSplit[1] == cmp
}

func (c *Claims) verifyAudience(cmp string) bool {
	return len(c.RegisteredClaims.Audience) == 1 && c.RegisteredClaims.Audience[0] == cmp
}

// ENDMANUAL(Types)
