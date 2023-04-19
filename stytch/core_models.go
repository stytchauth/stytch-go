package stytch

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Name struct {
	FirstName  string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
}

type Email struct {
	EmailID  string `json:"email_id,omitempty"`
	Email    string `json:"email,omitempty"`
	Verified bool   `json:"verified,omitempty"`
}

type EmailString struct {
	Email string `json:"email,omitempty"`
}

type PhoneNumber struct {
	PhoneID     string `json:"phone_id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Verified    bool   `json:"verified,omitempty"`
}

type PhoneNumberString struct {
	PhoneNumber string `json:"phone_number,omitempty"`
}

type CryptoWalletString struct {
	CryptoWalletAddress string `json:"crypto_wallet_address,omitempty"`
	CryptoWalletType    string `json:"crypto_wallet_type,omitempty"`
}

type WebAuthnRegistration struct {
	WebAuthnRegistrationID string `json:"webauthn_registration_id,omitempty"`
	Domain                 string `json:"domain,omitempty"`
	UserAgent              string `json:"user_agent,omitempty"`
	Verified               bool   `json:"verified,omitempty"`
	AuthenticatorType      string `json:"authenticator_type,omitempty"`
}

type OAuthProvider struct {
	OAuthUserRegistrationID string `json:"oauth_user_registration_id,omitempty"`
	ProviderType            string `json:"provider_type,omitempty"`
	ProviderSubject         string `json:"provider_subject,omitempty"`
	ProfilePictureURL       string `json:"profile_picture_url,omitempty"`
	Locale                  string `json:"locale,omitempty"`
}

type UserTOTP struct {
	TOTPID   string `json:"totp_id,omitempty"`
	Verified bool   `json:"verified,omitempty"`
}

type CryptoWallet struct {
	CryptoWalletID      string `json:"crypto_wallet_id,omitempty"`
	CryptoWalletAddress string `json:"crypto_wallet_address,omitempty"`
	CryptoWalletType    string `json:"crypto_wallet_type,omitempty"`
	Verified            bool   `json:"verified,omitempty"`
}

type Password struct {
	PasswordID    string `json:"password_id,omitempty"`
	RequiresReset bool   `json:"requires_reset,omitempty"`
}

type Feedback struct {
	Warning          string           `json:"warning,omitempty"`
	Suggestions      []string         `json:"suggestions,omitempty"`
	LUDSRequirements LUDSRequirements `json:"luds_requirements,omitempty"`
}

type LUDSRequirements struct {
	HasLowerCase      bool `json:"has_lower_case,omitempty"`
	HasUpperCase      bool `json:"has_upper_case,omitempty"`
	HasDigit          bool `json:"has_digit,omitempty"`
	HasSymbol         bool `json:"has_symbol,omitempty"`
	MissingComplexity int  `json:"missing_complexity,omitempty"`
	MissingCharacters int  `json:"missing_characters,omitempty"`
}

type MD5Config struct {
	PrependSalt string `json:"prepend_salt,omitempty"`
	AppendSalt  string `json:"append_salt,omitempty"`
}

type Argon2Config struct {
	Salt            string `json:"salt"`
	IterationAmount int    `json:"iteration_amount"`
	Memory          int    `json:"memory"`
	Threads         int    `json:"threads"`
	KeyLength       int    `json:"key_length"`
}

type SHA1Config struct {
	PrependSalt string `json:"prepend_salt,omitempty"`
	AppendSalt  string `json:"append_salt,omitempty"`
}

type ScryptConfig struct {
	Salt       string `json:"salt"`
	NParameter int    `json:"n_parameter"`
	RParameter int    `json:"r_parameter"`
	PParameter int    `json:"p_parameter"`
	KeyLength  int    `json:"key_length"`
}

type Key struct {
	Typ     string   `json:"kty"`
	Use     string   `json:"use"`
	KeyOps  []string `json:"key_ops"`
	Alg     string   `json:"alg"`
	KeyID   string   `json:"kid"`
	X5C     []string `json:"x5c"`
	X5TS256 string   `json:"x5tS256"`
	N       string   `json:"n"`
	E       string   `json:"e"`
}

type SessionClaim struct {
	ID                    string                 `json:"id"`
	StartedAt             string                 `json:"started_at"`
	LastAccessedAt        string                 `json:"last_accessed_at"`
	ExpiresAt             string                 `json:"expires_at"`
	Attributes            Attributes             `json:"attributes"`
	AuthenticationFactors []AuthenticationFactor `json:"authentication_factors"`
}

type Claims struct {
	StytchSession SessionClaim `json:"https://stytch.com/session"`
	jwt.RegisteredClaims
}

type HashType string

const (
	HashTypeBcrypt   HashType = "bcrypt"
	HashTypeMD5      HashType = "md_5"
	HashTypeArgon2I  HashType = "argon_2i"
	HashTypeArgon2ID HashType = "argon_2id"
	HashTypeSHA1     HashType = "sha_1"
	HashTypeScrypt   HashType = "scrypt"
	HashTypePHPass   HashType = "phpass"
)

type UsersSearchParams struct {
	Limit  int32             `json:"limit,omitempty"`
	Query  *UsersSearchQuery `json:"query,omitempty"`
	Cursor string            `json:"cursor,omitempty"`
}

type UsersSearchOperator string

const (
	UserSearchOperatorOR  UsersSearchOperator = "OR"
	UserSearchOperatorAND UsersSearchOperator = "AND"
)

type UsersSearchQuery struct {
	Operator UsersSearchOperator `json:"operator,omitempty"`
	Operands []json.Marshaler    `json:"operands,omitempty"`
}

type UsersSearchQueryCreatedAtGreaterThanFilter struct {
	CreatedAtGreaterThan time.Time
}

func marshalFilter(filterName string, filterValue interface{}) ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter_name":  filterName,
		"filter_value": filterValue,
	})
}

func (q UsersSearchQueryCreatedAtGreaterThanFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("created_at_greater_than", q.CreatedAtGreaterThan)
}

type UsersSearchQueryCreatedAtLessThanFilter struct {
	CreatedAtLessThan time.Time
}

func (q UsersSearchQueryCreatedAtLessThanFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("created_at_less_than", q.CreatedAtLessThan)
}

type UsersSearchQueryCreatedAtBetweenFilter struct {
	GreaterThan time.Time
	LessThan    time.Time
}

func (q UsersSearchQueryCreatedAtBetweenFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("created_at_between", map[string]time.Time{
		"greater_than": q.GreaterThan,
		"less_than":    q.LessThan,
	})
}

/* User Filters */
type UsersSearchQueryStatusFilter struct {
	Status string
}

func (q UsersSearchQueryStatusFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("status", q.Status)
}

type UsersSearchQueryUserIDFilter struct {
	UserIDs []string
}

func (q UsersSearchQueryUserIDFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("user_id", q.UserIDs)
}

type UsersSearchQueryFullNameFuzzyFilter struct {
	FullNameFuzzy string
}

func (q UsersSearchQueryFullNameFuzzyFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("full_name_fuzzy", q.FullNameFuzzy)
}

/* Phone Number Filters */

type UsersSearchQueryPhoneNumberFilter struct {
	PhoneNumbers []string
}

func (q UsersSearchQueryPhoneNumberFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("phone_number", q.PhoneNumbers)
}

type UsersSearchQueryPhoneIDFilter struct {
	PhoneIDs []string
}

func (q UsersSearchQueryPhoneIDFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("phone_id", q.PhoneIDs)
}

type UsersSearchQueryPhoneVerifiedFilter struct {
	PhoneVerified bool
}

func (q UsersSearchQueryPhoneVerifiedFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("phone_verified", q.PhoneVerified)
}

type UsersSearchQueryPhoneNumberFuzzyFilter struct {
	PhoneNumberFuzzy string
}

func (q UsersSearchQueryPhoneNumberFuzzyFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("phone_number_fuzzy", q.PhoneNumberFuzzy)
}

/* Email Filters */

type UsersSearchQueryEmailAddressFilter struct {
	EmailAddresses []string
}

func (q UsersSearchQueryEmailAddressFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("email_address", q.EmailAddresses)
}

type UsersSearchQueryEmailIDFilter struct {
	EmailIDs []string
}

func (q UsersSearchQueryEmailIDFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("email_id", q.EmailIDs)
}

type UsersSearchQueryEmailVerifiedFilter struct {
	EmailVerified bool
}

func (q UsersSearchQueryEmailVerifiedFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("email_verified", q.EmailVerified)
}

type UsersSearchQueryEmailAddressFuzzyFilter struct {
	EmailAddressFuzzy string
}

func (q UsersSearchQueryEmailAddressFuzzyFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("email_address_fuzzy", q.EmailAddressFuzzy)
}

/* WebAuthn Filters */

type UsersSearchQueryWebAuthnRegistrationVerifiedFilter struct {
	WebAuthnRegistrationVerified bool
}

func (q UsersSearchQueryWebAuthnRegistrationVerifiedFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("webauthn_registration_verified", q.WebAuthnRegistrationVerified)
}

type UsersSearchQueryWebAuthnRegistrationIDFilter struct {
	WebAuthnRegistrationIDs []string
}

func (q UsersSearchQueryWebAuthnRegistrationIDFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("webauthn_registration_id", q.WebAuthnRegistrationIDs)
}

/* Crypto Wallet Filters */

type UsersSearchQueryCryptoWalletIDFilter struct {
	CryptoWalletIDs []string
}

func (q UsersSearchQueryCryptoWalletIDFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("crypto_wallet_id", q.CryptoWalletIDs)
}

type UsersSearchQueryCryptoWalletAddressFilter struct {
	CryptoWalletAddresses []string
}

func (q UsersSearchQueryCryptoWalletAddressFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("crypto_wallet_address", q.CryptoWalletAddresses)
}

type UsersSearchQueryCryptoWalletVerifiedFilter struct {
	CryptoWalletVerified bool
}

func (q UsersSearchQueryCryptoWalletVerifiedFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("crypto_wallet_verified", q.CryptoWalletVerified)
}

type UsersSearchQueryPasswordExistsFilter struct {
	PasswordExists bool
}

func (q UsersSearchQueryPasswordExistsFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("password_exists", q.PasswordExists)
}

/* OAuth Filters */

type UsersSearchQueryOAuthProviderFilter struct {
	OAuthProviders []string
}

func (q UsersSearchQueryOAuthProviderFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("oauth_provider", q.OAuthProviders)
}

/* TOTP Filters */

type UsersSearchQueryTOTPIDFilter struct {
	TOTPIDs []string
}

func (q UsersSearchQueryTOTPIDFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("totp_id", q.TOTPIDs)
}

type UsersSearchQueryTOTPVerifiedFilter struct {
	TOTPVerified bool
}

func (q UsersSearchQueryTOTPVerifiedFilter) MarshalJSON() ([]byte, error) {
	return marshalFilter("totp_verified", q.TOTPVerified)
}

type User struct {
	UserID                string                 `json:"user_id,omitempty"`
	Name                  Name                   `json:"name,omitempty"`
	Emails                []Email                `json:"emails,omitempty"`
	PhoneNumbers          []PhoneNumber          `json:"phone_numbers,omitempty"`
	WebAuthnRegistrations []WebAuthnRegistration `json:"webauthn_registrations,omitempty"`
	OAuthProviders        []OAuthProvider        `json:"providers,omitempty"`
	TOTPs                 []UserTOTP             `json:"totps,omitempty"`
	Password              Password               `json:"password,omitempty"`
	Status                string                 `json:"status,omitempty"`
	CreatedAt             time.Time              `json:"created_at,omitempty"`
	TrustedMetadata       map[string]interface{} `json:"trusted_metadata,omitempty"`
	UntrustedMetadata     map[string]interface{} `json:"untrusted_metadata,omitempty"`
}

type PendingUser struct {
	UserID        string         `json:"user_id,omitempty"`
	Name          Name           `json:"name,omitempty"`
	Emails        []Email        `json:"emails,omitempty"`
	PhoneNumbers  []PhoneNumber  `json:"phone_numbers,omitempty"`
	TOTPs         []UserTOTP     `json:"totps,omitempty"`
	CryptoWallets []CryptoWallet `json:"crypto_wallets,omitempty"`
	Password      Password       `json:"password,omitempty"`
	Status        string         `json:"status,omitempty"`
	InvitedAt     string         `json:"invited_at,omitempty"`
}

type Member struct{}

type Organization struct{}

type DiscoveredOrganization struct{}

type MemberSession struct{}

type LudsFeedback struct{}

type ZxcvbnFeedback struct{}

type ProviderValues struct {
	AccessToken  string     `json:"access_token,omitempty"`
	RefreshToken string     `json:"refresh_token,omitempty"`
	IDToken      string     `json:"id_token,omitempty"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty"`
	Scopes       []string   `json:"scopes,omitempty"`
}

type ExternalSearchQuery struct{}

type ResultsMetadata struct {
	NextCursor string `json:"next_cursor,omitempty"`
	Total      int    `json:"total,omitempty"`
}

type JWK struct{}

type OIDCConnection struct{}

type SAMLConnection struct{}

type TOTP struct {
	TOTPID        string   `json:"totp_id,omitempty"`
	Verified      bool     `json:"verified,omitempty"`
	RecoveryCodes []string `json:"recovery_codes,omitempty"`
}

type Birthday struct {
	Month int32 `json:"month"`
	Day   int32 `json:"day"`
	Year  int32 `json:"year"`
}
