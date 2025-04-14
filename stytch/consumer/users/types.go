package users

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"time"

	"github.com/stytchauth/stytch-go/v16/stytch/consumer/attribute"
)

// BiometricRegistration:
type BiometricRegistration struct {
	// BiometricRegistrationID: The unique ID for a biometric registration.
	BiometricRegistrationID string `json:"biometric_registration_id,omitempty"`
	// Verified: The verified boolean denotes whether or not this send method, e.g. phone number, email
	// address, etc., has been successfully authenticated by the User.
	Verified bool `json:"verified,omitempty"`
}

// CreateParams: Request type for `Users.Create`.
type CreateParams struct {
	// Email: The email address of the end user.
	Email string `json:"email,omitempty"`
	// Name: The name of the user. Each field in the name object is optional.
	Name *Name `json:"name,omitempty"`
	// Attributes: Provided attributes help with fraud detection.
	Attributes *attribute.Attributes `json:"attributes,omitempty"`
	// PhoneNumber: The phone number to use for one-time passcodes. The phone number should be in E.164 format
	// (i.e. +1XXXXXXXXXX). You may use +10000000000 to test this endpoint, see
	// [Testing](https://stytch.com/docs/home#resources_testing) for more detail.
	PhoneNumber string `json:"phone_number,omitempty"`
	// CreateUserAsPending: Flag for whether or not to save a user as pending vs active in Stytch. Defaults to
	// false.
	//         If true, users will be saved with status pending in Stytch's backend until authenticated.
	//         If false, users will be created as active. An example usage of
	//         a true flag would be to require users to verify their phone by entering the OTP code before
	// creating
	//         an account for them.
	CreateUserAsPending bool `json:"create_user_as_pending,omitempty"`
	// TrustedMetadata: The `trusted_metadata` field contains an arbitrary JSON object of application-specific
	// data. See the [Metadata](https://stytch.com/docs/api/metadata) reference for complete field behavior
	// details.
	TrustedMetadata map[string]any `json:"trusted_metadata,omitempty"`
	// UntrustedMetadata: The `untrusted_metadata` field contains an arbitrary JSON object of
	// application-specific data. Untrusted metadata can be edited by end users directly via the SDK, and
	// **cannot be used to store critical information.** See the
	// [Metadata](https://stytch.com/docs/api/metadata) reference for complete field behavior details.
	UntrustedMetadata map[string]any `json:"untrusted_metadata,omitempty"`
	// ExternalID: An identifier that can be used in API calls wherever a user_id is expected. This is a string
	// consisting of alphanumeric, `.`, `_`, `-`, or `|` characters with a maximum length of 128 characters.
	// External IDs must be unique within an organization, but may be reused across different organizations in
	// the same project.
	ExternalID string `json:"external_id,omitempty"`
}

// CryptoWallet:
type CryptoWallet struct {
	// CryptoWalletID: The unique ID for a crypto wallet
	CryptoWalletID string `json:"crypto_wallet_id,omitempty"`
	// CryptoWalletAddress: The actual blockchain address of the User's crypto wallet.
	CryptoWalletAddress string `json:"crypto_wallet_address,omitempty"`
	// CryptoWalletType: The blockchain that the User's crypto wallet operates on, e.g. Ethereum, Solana, etc.
	CryptoWalletType string `json:"crypto_wallet_type,omitempty"`
	// Verified: The verified boolean denotes whether or not this send method, e.g. phone number, email
	// address, etc., has been successfully authenticated by the User.
	Verified bool `json:"verified,omitempty"`
}

// DeleteBiometricRegistrationParams: Request type for `Users.DeleteBiometricRegistration`.
type DeleteBiometricRegistrationParams struct {
	// BiometricRegistrationID: The `biometric_registration_id` to be deleted.
	BiometricRegistrationID string `json:"biometric_registration_id,omitempty"`
}

// DeleteCryptoWalletParams: Request type for `Users.DeleteCryptoWallet`.
type DeleteCryptoWalletParams struct {
	// CryptoWalletID: The `crypto_wallet_id` to be deleted.
	CryptoWalletID string `json:"crypto_wallet_id,omitempty"`
}

// DeleteEmailParams: Request type for `Users.DeleteEmail`.
type DeleteEmailParams struct {
	// EmailID: The `email_id` to be deleted.
	EmailID string `json:"email_id,omitempty"`
}

// DeleteOAuthRegistrationParams: Request type for `Users.DeleteOAuthRegistration`.
type DeleteOAuthRegistrationParams struct {
	// OAuthUserRegistrationID: The `oauth_user_registration_id` to be deleted.
	OAuthUserRegistrationID string `json:"oauth_user_registration_id,omitempty"`
}

// DeleteParams: Request type for `Users.Delete`.
type DeleteParams struct {
	// UserID: The unique ID of a specific User. You may use an external_id here if one is set for the user.
	UserID string `json:"user_id,omitempty"`
}

// DeletePasswordParams: Request type for `Users.DeletePassword`.
type DeletePasswordParams struct {
	// PasswordID: The `password_id` to be deleted.
	PasswordID string `json:"password_id,omitempty"`
}

// DeletePhoneNumberParams: Request type for `Users.DeletePhoneNumber`.
type DeletePhoneNumberParams struct {
	// PhoneID: The `phone_id` to be deleted.
	PhoneID string `json:"phone_id,omitempty"`
}

// DeleteTOTPParams: Request type for `Users.DeleteTOTP`.
type DeleteTOTPParams struct {
	// TOTPID: The `totp_id` to be deleted.
	TOTPID string `json:"totp_id,omitempty"`
}

// DeleteWebAuthnRegistrationParams: Request type for `Users.DeleteWebAuthnRegistration`.
type DeleteWebAuthnRegistrationParams struct {
	// WebAuthnRegistrationID: The `webauthn_registration_id` to be deleted.
	WebAuthnRegistrationID string `json:"webauthn_registration_id,omitempty"`
}

// Email:
type Email struct {
	// EmailID: The unique ID of a specific email address.
	EmailID string `json:"email_id,omitempty"`
	// Email: The email address.
	Email string `json:"email,omitempty"`
	// Verified: The verified boolean denotes whether or not this send method, e.g. phone number, email
	// address, etc., has been successfully authenticated by the User.
	Verified bool `json:"verified,omitempty"`
}

// ExchangePrimaryFactorParams: Request type for `Users.ExchangePrimaryFactor`.
type ExchangePrimaryFactorParams struct {
	// UserID: The unique ID of a specific User. You may use an external_id here if one is set for the user.
	UserID string `json:"user_id,omitempty"`
	// EmailAddress: The email address to exchange to.
	EmailAddress string `json:"email_address,omitempty"`
	// PhoneNumber: The phone number to exchange to. The phone number should be in E.164 format (i.e.
	// +1XXXXXXXXXX).
	PhoneNumber string `json:"phone_number,omitempty"`
}

// GetParams: Request type for `Users.Get`.
type GetParams struct {
	// UserID: The unique ID of a specific User. You may use an external_id here if one is set for the user.
	UserID string `json:"user_id,omitempty"`
}

// Name:
type Name struct {
	// FirstName: The first name of the user.
	FirstName string `json:"first_name,omitempty"`
	// MiddleName: The middle name(s) of the user.
	MiddleName string `json:"middle_name,omitempty"`
	// LastName: The last name of the user.
	LastName string `json:"last_name,omitempty"`
}

// OAuthProvider:
type OAuthProvider struct {
	// ProviderType: Denotes the OAuth identity provider that the user has authenticated with, e.g. Google,
	// Facebook, GitHub etc.
	ProviderType string `json:"provider_type,omitempty"`
	// ProviderSubject: The unique identifier for the User within a given OAuth provider. Also commonly called
	// the "sub" or "Subject field" in OAuth protocols.
	ProviderSubject string `json:"provider_subject,omitempty"`
	// ProfilePictureURL: If available, the `profile_picture_url` is a url of the User's profile picture set in
	// OAuth identity the provider that the User has authenticated with, e.g. Facebook profile picture.
	ProfilePictureURL string `json:"profile_picture_url,omitempty"`
	// Locale: If available, the `locale` is the User's locale set in the OAuth identity provider that the user
	// has authenticated with.
	Locale string `json:"locale,omitempty"`
	// OAuthUserRegistrationID: The unique ID for an OAuth registration.
	OAuthUserRegistrationID string `json:"oauth_user_registration_id,omitempty"`
}

// Password:
type Password struct {
	// PasswordID: The unique ID of a specific password
	PasswordID string `json:"password_id,omitempty"`
	// RequiresReset: Indicates whether this password requires a password reset
	RequiresReset bool `json:"requires_reset,omitempty"`
}

// PhoneNumber:
type PhoneNumber struct {
	// PhoneID: The unique ID for the phone number.
	PhoneID string `json:"phone_id,omitempty"`
	// PhoneNumber: The phone number.
	PhoneNumber string `json:"phone_number,omitempty"`
	// Verified: The verified boolean denotes whether or not this send method, e.g. phone number, email
	// address, etc., has been successfully authenticated by the User.
	Verified bool `json:"verified,omitempty"`
}

// ResultsMetadata:
type ResultsMetadata struct {
	// Total: The total number of results returned by your search query. If totals have been disabled for your
	// Stytch Workspace to improve search performance, the value will always be -1.
	Total int32 `json:"total,omitempty"`
	// NextCursor: The `next_cursor` string is returned when your search result contains more than one page of
	// results. This value is passed into your next search call in the `cursor` field.
	NextCursor string `json:"next_cursor,omitempty"`
}

// SearchParams: Request type for `Users.Search`.
type SearchParams struct {
	// Cursor: The `cursor` field allows you to paginate through your results. Each result array is limited to
	// 1000 results. If your query returns more than 1000 results, you will need to paginate the responses
	// using the `cursor`. If you receive a response that includes a non-null `next_cursor` in the
	// `results_metadata` object, repeat the search call with the `next_cursor` value set to the `cursor` field
	// to retrieve the next page of results. Continue to make search calls until the `next_cursor` in the
	// response is null.
	Cursor string `json:"cursor,omitempty"`
	// Limit: The number of search results to return per page. The default limit is 100. A maximum of 1000
	// results can be returned by a single search request. If the total size of your result set is greater than
	// one page size, you must paginate the response. See the `cursor` field.
	Limit uint32 `json:"limit,omitempty"`
	// Query: The optional query object contains the operator, i.e. `AND` or `OR`, and the operands that will
	// filter your results. Only an operator is required. If you include no operands, no filtering will be
	// applied. If you include no query object, it will return all results with no filtering applied.
	Query *SearchUsersQuery `json:"query,omitempty"`
}

// SearchUsersQuery:
type SearchUsersQuery struct {
	// Operator: The action to perform on the operands. The accepted value are:
	//
	//   `AND` – all the operand values provided must match.
	//
	//   `OR` – the operator will return any matches to at least one of the operand values you supply.
	Operator SearchUsersQueryOperator `json:"operator,omitempty"`
	// Operands: An array of operand objects that contains all of the filters and values to apply to your
	// search search query.
	Operands []map[string]any `json:"operands,omitempty"`
}

// TOTP:
type TOTP struct {
	// TOTPID: The unique ID for a TOTP instance.
	TOTPID string `json:"totp_id,omitempty"`
	// Verified: The verified boolean denotes whether or not this send method, e.g. phone number, email
	// address, etc., has been successfully authenticated by the User.
	Verified bool `json:"verified,omitempty"`
}

// UpdateParams: Request type for `Users.Update`.
type UpdateParams struct {
	// UserID: The unique ID of a specific User. You may use an external_id here if one is set for the user.
	UserID string `json:"user_id,omitempty"`
	// Name: The name of the user. Each field in the name object is optional.
	Name *Name `json:"name,omitempty"`
	// Attributes: Provided attributes help with fraud detection.
	Attributes *attribute.Attributes `json:"attributes,omitempty"`
	// TrustedMetadata: The `trusted_metadata` field contains an arbitrary JSON object of application-specific
	// data. See the [Metadata](https://stytch.com/docs/api/metadata) reference for complete field behavior
	// details.
	TrustedMetadata map[string]any `json:"trusted_metadata,omitempty"`
	// UntrustedMetadata: The `untrusted_metadata` field contains an arbitrary JSON object of
	// application-specific data. Untrusted metadata can be edited by end users directly via the SDK, and
	// **cannot be used to store critical information.** See the
	// [Metadata](https://stytch.com/docs/api/metadata) reference for complete field behavior details.
	UntrustedMetadata map[string]any `json:"untrusted_metadata,omitempty"`
	// ExternalID: An identifier that can be used in API calls wherever a user_id is expected. This is a string
	// consisting of alphanumeric, `.`, `_`, `-`, or `|` characters with a maximum length of 128 characters.
	// External IDs must be unique within an organization, but may be reused across different organizations in
	// the same project.
	ExternalID string `json:"external_id,omitempty"`
}

// User:
type User struct {
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// Emails: An array of email objects for the User.
	Emails []Email `json:"emails,omitempty"`
	// Status: The status of the User. The possible values are `pending` and `active`.
	Status string `json:"status,omitempty"`
	// PhoneNumbers: An array of phone number objects linked to the User.
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
	// WebAuthnRegistrations: An array that contains a list of all Passkey or WebAuthn registrations for a
	// given User in the Stytch API.
	WebAuthnRegistrations []WebAuthnRegistration `json:"webauthn_registrations,omitempty"`
	// Providers: An array of OAuth `provider` objects linked to the User.
	Providers []OAuthProvider `json:"providers,omitempty"`
	// TOTPs: An array containing a list of all TOTP instances for a given User in the Stytch API.
	TOTPs []TOTP `json:"totps,omitempty"`
	// CryptoWallets: An array contains a list of all crypto wallets for a given User in the Stytch API.
	CryptoWallets []CryptoWallet `json:"crypto_wallets,omitempty"`
	// BiometricRegistrations: An array that contains a list of all biometric registrations for a given User in
	// the Stytch API.
	BiometricRegistrations []BiometricRegistration `json:"biometric_registrations,omitempty"`
	// Name: The name of the User. Each field in the `name` object is optional.
	Name *Name `json:"name,omitempty"`
	// CreatedAt: The timestamp of the User's creation. Values conform to the RFC 3339 standard and are
	// expressed in UTC, e.g. `2021-12-29T12:33:09Z`.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Password: The password object is returned for users with a password.
	Password *Password `json:"password,omitempty"`
	// TrustedMetadata: The `trusted_metadata` field contains an arbitrary JSON object of application-specific
	// data. See the [Metadata](https://stytch.com/docs/api/metadata) reference for complete field behavior
	// details.
	TrustedMetadata map[string]any `json:"trusted_metadata,omitempty"`
	// UntrustedMetadata: The `untrusted_metadata` field contains an arbitrary JSON object of
	// application-specific data. Untrusted metadata can be edited by end users directly via the SDK, and
	// **cannot be used to store critical information.** See the
	// [Metadata](https://stytch.com/docs/api/metadata) reference for complete field behavior details.
	UntrustedMetadata map[string]any `json:"untrusted_metadata,omitempty"`
	ExternalID        string         `json:"external_id,omitempty"`
}

// WebAuthnRegistration:
type WebAuthnRegistration struct {
	// WebAuthnRegistrationID: The unique ID for the Passkey or WebAuthn registration.
	WebAuthnRegistrationID string `json:"webauthn_registration_id,omitempty"`
	// Domain: The `domain` on which Passkey or WebAuthn registration was started. This will be the domain of
	// your app.
	Domain string `json:"domain,omitempty"`
	// UserAgent: The user agent of the User.
	UserAgent string `json:"user_agent,omitempty"`
	// Verified: The verified boolean denotes whether or not this send method, e.g. phone number, email
	// address, etc., has been successfully authenticated by the User.
	Verified bool `json:"verified,omitempty"`
	// AuthenticatorType: The `authenticator_type` string displays the requested authenticator type of the
	// Passkey or WebAuthn device. The two valid types are "platform" and "cross-platform". If no value is
	// present, the Passkey or WebAuthn device was created without an authenticator type preference.
	AuthenticatorType string `json:"authenticator_type,omitempty"`
	// Name: The `name` of the Passkey or WebAuthn registration.
	Name string `json:"name,omitempty"`
}

// CreateResponse: Response type for `Users.Create`.
type CreateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// EmailID: The unique ID of a specific email address.
	EmailID string `json:"email_id,omitempty"`
	// Status: The status of the User. The possible values are `pending` and `active`.
	Status string `json:"status,omitempty"`
	// PhoneID: The unique ID for the phone number.
	PhoneID string `json:"phone_id,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User User `json:"user,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// DeleteBiometricRegistrationResponse: Response type for `Users.DeleteBiometricRegistration`.
type DeleteBiometricRegistrationResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User User `json:"user,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// DeleteCryptoWalletResponse: Response type for `Users.DeleteCryptoWallet`.
type DeleteCryptoWalletResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User User `json:"user,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// DeleteEmailResponse: Response type for `Users.DeleteEmail`.
type DeleteEmailResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User User `json:"user,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// DeleteOAuthRegistrationResponse: Response type for `Users.DeleteOAuthRegistration`.
type DeleteOAuthRegistrationResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User User `json:"user,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// DeletePasswordResponse: Response type for `Users.DeletePassword`.
type DeletePasswordResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User User `json:"user,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// DeletePhoneNumberResponse: Response type for `Users.DeletePhoneNumber`.
type DeletePhoneNumberResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User User `json:"user,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// DeleteResponse: Response type for `Users.Delete`.
type DeleteResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the deleted User.
	UserID string `json:"user_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// DeleteTOTPResponse: Response type for `Users.DeleteTOTP`.
type DeleteTOTPResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User User `json:"user,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// DeleteWebAuthnRegistrationResponse: Response type for `Users.DeleteWebAuthnRegistration`.
type DeleteWebAuthnRegistrationResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User User `json:"user,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// ExchangePrimaryFactorResponse: Response type for `Users.ExchangePrimaryFactor`.
type ExchangePrimaryFactorResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User User `json:"user,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// GetResponse: Response type for `Users.Get`.
type GetResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the returned User.
	UserID string `json:"user_id,omitempty"`
	// Emails: An array of email objects for the User.
	Emails []Email `json:"emails,omitempty"`
	// Status: The status of the User. The possible values are `pending` and `active`.
	Status string `json:"status,omitempty"`
	// PhoneNumbers: An array of phone number objects linked to the User.
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
	// WebAuthnRegistrations: An array that contains a list of all Passkey or WebAuthn registrations for a
	// given User in the Stytch API.
	WebAuthnRegistrations []WebAuthnRegistration `json:"webauthn_registrations,omitempty"`
	// Providers: An array of OAuth `provider` objects linked to the User.
	Providers []OAuthProvider `json:"providers,omitempty"`
	// TOTPs: An array containing a list of all TOTP instances for a given User in the Stytch API.
	TOTPs []TOTP `json:"totps,omitempty"`
	// CryptoWallets: An array contains a list of all crypto wallets for a given User in the Stytch API.
	CryptoWallets []CryptoWallet `json:"crypto_wallets,omitempty"`
	// BiometricRegistrations: An array that contains a list of all biometric registrations for a given User in
	// the Stytch API.
	BiometricRegistrations []BiometricRegistration `json:"biometric_registrations,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// Name: The name of the User. Each field in the `name` object is optional.
	Name *Name `json:"name,omitempty"`
	// CreatedAt: The timestamp of the User's creation. Values conform to the RFC 3339 standard and are
	// expressed in UTC, e.g. `2021-12-29T12:33:09Z`.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Password: The password object is returned for users with a password.
	Password *Password `json:"password,omitempty"`
	// TrustedMetadata: The `trusted_metadata` field contains an arbitrary JSON object of application-specific
	// data. See the [Metadata](https://stytch.com/docs/api/metadata) reference for complete field behavior
	// details.
	TrustedMetadata map[string]any `json:"trusted_metadata,omitempty"`
	// UntrustedMetadata: The `untrusted_metadata` field contains an arbitrary JSON object of
	// application-specific data. Untrusted metadata can be edited by end users directly via the SDK, and
	// **cannot be used to store critical information.** See the
	// [Metadata](https://stytch.com/docs/api/metadata) reference for complete field behavior details.
	UntrustedMetadata map[string]any `json:"untrusted_metadata,omitempty"`
	ExternalID        string         `json:"external_id,omitempty"`
}

// SearchResponse: Response type for `Users.Search`.
type SearchResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// Results: An array of results that match your search query.
	Results []User `json:"results,omitempty"`
	// ResultsMetadata: The search `results_metadata` object contains metadata relevant to your specific query
	// like total and `next_cursor`.
	ResultsMetadata ResultsMetadata `json:"results_metadata,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// UpdateResponse: Response type for `Users.Update`.
type UpdateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the updated User.
	UserID string `json:"user_id,omitempty"`
	// Emails: An array of email objects for the User.
	Emails []Email `json:"emails,omitempty"`
	// PhoneNumbers: An array of phone number objects linked to the User.
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
	// CryptoWallets: An array contains a list of all crypto wallets for a given User in the Stytch API.
	CryptoWallets []CryptoWallet `json:"crypto_wallets,omitempty"`
	// User: The `user` object affected by this API call. See the
	// [Get user endpoint](https://stytch.com/docs/api/get-user) for complete response field details.
	User User `json:"user,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

type SearchUsersQueryOperator string

const (
	SearchUsersQueryOperatorOR  SearchUsersQueryOperator = "OR"
	SearchUsersQueryOperatorAND SearchUsersQueryOperator = "AND"
)
