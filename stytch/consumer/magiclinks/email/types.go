package email

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v10/stytch/consumer/attribute"
	"github.com/stytchauth/stytch-go/v10/stytch/consumer/users"
)

// InviteParams: Request type for `Email.Invite`.
type InviteParams struct {
	// Email: The email address of the User to send the invite Magic Link to.
	Email string `json:"email,omitempty"`
	// InviteTemplateID: Use a custom template for invite emails. By default, it will use your default email
	// template. The template must be a template using our built-in customizations or a custom HTML email for
	// Magic links - Invite.
	InviteTemplateID string `json:"invite_template_id,omitempty"`
	// Attributes: Provided attributes help with fraud detection.
	Attributes attribute.Attributes `json:"attributes,omitempty"`
	// Name: The name of the user. Each field in the name object is optional.
	Name users.Name `json:"name,omitempty"`
	// InviteMagicLinkURL: The URL the end user clicks from the Email Magic Link. This should be a URL that
	// your app receives and parses and subsequently sends an API request to authenticate the Magic Link and
	// log in the User. If this value is not passed, the default invite redirect URL that you set in your
	// Dashboard is used. If you have not set a default sign-up redirect URL, an error is returned.
	InviteMagicLinkURL string `json:"invite_magic_link_url,omitempty"`
	// InviteExpirationMinutes: Set the expiration for the email magic link, in minutes. By default, it expires
	// in 1 hour. The minimum expiration is 5 minutes and the maximum is 7 days (10080 mins).
	InviteExpirationMinutes int32 `json:"invite_expiration_minutes,omitempty"`
	// Locale: Used to determine which language to use when sending the user this delivery method. Parameter is
	// a [IETF BCP 47 language tag](https://www.w3.org/International/articles/language-tags/), e.g. `"en"`.
	//
	// Currently supported languages are English (`"en"`), Spanish (`"es"`), and Brazilian Portuguese
	// (`"pt-br"`); if no value is provided, the copy defaults to English.
	//
	// Request support for additional languages
	// [here](https://docs.google.com/forms/d/e/1FAIpQLScZSpAu_m2AmLXRT3F3kap-s_mcV6UTBitYn6CdyWP0-o7YjQ/viewform?usp=sf_link")!
	//
	Locale InviteRequestLocale `json:"locale,omitempty"`
}

// LoginOrCreateParams: Request type for `Email.LoginOrCreate`.
type LoginOrCreateParams struct {
	// Email: The email address of the end user.
	Email string `json:"email,omitempty"`
	// LoginMagicLinkURL: The URL the end user clicks from the login Email Magic Link. This should be a URL
	// that your app receives and parses and subsequently send an API request to authenticate the Magic Link
	// and log in the User. If this value is not passed, the default login redirect URL that you set in your
	// Dashboard is used. If you have not set a default login redirect URL, an error is returned.
	LoginMagicLinkURL string `json:"login_magic_link_url,omitempty"`
	// SignupMagicLinkURL: The URL the end user clicks from the sign-up Email Magic Link. This should be a URL
	// that your app receives and parses and subsequently send an API request to authenticate the Magic Link
	// and sign-up the User. If this value is not passed, the default sign-up redirect URL that you set in your
	// Dashboard is used. If you have not set a default sign-up redirect URL, an error is returned.
	SignupMagicLinkURL string `json:"signup_magic_link_url,omitempty"`
	// LoginExpirationMinutes: Set the expiration for the login email magic link, in minutes. By default, it
	// expires in 1 hour. The minimum expiration is 5 minutes and the maximum is 7 days (10080 mins).
	LoginExpirationMinutes int32 `json:"login_expiration_minutes,omitempty"`
	// SignupExpirationMinutes: Set the expiration for the sign-up email magic link, in minutes. By default, it
	// expires in 1 week. The minimum expiration is 5 minutes and the maximum is 7 days (10080 mins).
	SignupExpirationMinutes int32 `json:"signup_expiration_minutes,omitempty"`
	// LoginTemplateID: Use a custom template for login emails. By default, it will use your default email
	// template. The template must be a template using our built-in customizations or a custom HTML email for
	// Magic links - Login.
	LoginTemplateID string `json:"login_template_id,omitempty"`
	// SignupTemplateID: Use a custom template for sign-up emails. By default, it will use your default email
	// template. The template must be a template using our built-in customizations or a custom HTML email for
	// Magic links - Sign-up.
	SignupTemplateID string `json:"signup_template_id,omitempty"`
	// Attributes: Provided attributes help with fraud detection.
	Attributes attribute.Attributes `json:"attributes,omitempty"`
	// CreateUserAsPending: Flag for whether or not to save a user as pending vs active in Stytch. Defaults to
	// false.
	//         If true, users will be saved with status pending in Stytch's backend until authenticated.
	//         If false, users will be created as active. An example usage of
	//         a true flag would be to require users to verify their phone by entering the OTP code before
	// creating
	//         an account for them.
	CreateUserAsPending bool `json:"create_user_as_pending,omitempty"`
	// CodeChallenge: A base64url encoded SHA256 hash of a one time secret used to validate that the request
	// starts and ends on the same device.
	CodeChallenge string `json:"code_challenge,omitempty"`
	// Locale: Used to determine which language to use when sending the user this delivery method. Parameter is
	// a [IETF BCP 47 language tag](https://www.w3.org/International/articles/language-tags/), e.g. `"en"`.
	//
	// Currently supported languages are English (`"en"`), Spanish (`"es"`), and Brazilian Portuguese
	// (`"pt-br"`); if no value is provided, the copy defaults to English.
	//
	// Request support for additional languages
	// [here](https://docs.google.com/forms/d/e/1FAIpQLScZSpAu_m2AmLXRT3F3kap-s_mcV6UTBitYn6CdyWP0-o7YjQ/viewform?usp=sf_link")!
	//
	Locale LoginOrCreateRequestLocale `json:"locale,omitempty"`
}

// RevokeInviteParams: Request type for `Email.RevokeInvite`.
type RevokeInviteParams struct {
	// Email: The email of the user.
	Email string `json:"email,omitempty"`
}

// SendParams: Request type for `Email.Send`.
type SendParams struct {
	// Email: The email address of the User to send the Magic Link to.
	Email string `json:"email,omitempty"`
	// LoginTemplateID: Use a custom template for login emails. By default, it will use your default email
	// template. The template must be a template using our built-in customizations or a custom HTML email for
	// Magic links - Login.
	LoginTemplateID string `json:"login_template_id,omitempty"`
	// Attributes: Provided attributes help with fraud detection.
	Attributes attribute.Attributes `json:"attributes,omitempty"`
	// LoginMagicLinkURL: The URL the end user clicks from the login Email Magic Link. This should be a URL
	// that your app receives and parses and subsequently send an API request to authenticate the Magic Link
	// and log in the User. If this value is not passed, the default login redirect URL that you set in your
	// Dashboard is used. If you have not set a default login redirect URL, an error is returned.
	LoginMagicLinkURL string `json:"login_magic_link_url,omitempty"`
	// SignupMagicLinkURL: The URL the end user clicks from the sign-up Email Magic Link. This should be a URL
	// that your app receives and parses and subsequently send an API request to authenticate the Magic Link
	// and sign-up the User. If this value is not passed, the default sign-up redirect URL that you set in your
	// Dashboard is used. If you have not set a default sign-up redirect URL, an error is returned.
	SignupMagicLinkURL string `json:"signup_magic_link_url,omitempty"`
	// LoginExpirationMinutes: Set the expiration for the login email magic link, in minutes. By default, it
	// expires in 1 hour. The minimum expiration is 5 minutes and the maximum is 7 days (10080 mins).
	LoginExpirationMinutes int32 `json:"login_expiration_minutes,omitempty"`
	// SignupExpirationMinutes: Set the expiration for the sign-up email magic link, in minutes. By default, it
	// expires in 1 week. The minimum expiration is 5 minutes and the maximum is 7 days (10080 mins).
	SignupExpirationMinutes int32 `json:"signup_expiration_minutes,omitempty"`
	// CodeChallenge: A base64url encoded SHA256 hash of a one time secret used to validate that the request
	// starts and ends on the same device.
	CodeChallenge string `json:"code_challenge,omitempty"`
	// UserID: The unique ID of a specific User.
	UserID string `json:"user_id,omitempty"`
	// SessionToken: The `session_token` of the user to associate the email with.
	SessionToken string `json:"session_token,omitempty"`
	// SessionJWT: The `session_jwt` of the user to associate the email with.
	SessionJWT string `json:"session_jwt,omitempty"`
	// Locale: Used to determine which language to use when sending the user this delivery method. Parameter is
	// a [IETF BCP 47 language tag](https://www.w3.org/International/articles/language-tags/), e.g. `"en"`.
	//
	// Currently supported languages are English (`"en"`), Spanish (`"es"`), and Brazilian Portuguese
	// (`"pt-br"`); if no value is provided, the copy defaults to English.
	//
	// Request support for additional languages
	// [here](https://docs.google.com/forms/d/e/1FAIpQLScZSpAu_m2AmLXRT3F3kap-s_mcV6UTBitYn6CdyWP0-o7YjQ/viewform?usp=sf_link")!
	//
	Locale SendRequestLocale `json:"locale,omitempty"`
	// SignupTemplateID: Use a custom template for sign-up emails. By default, it will use your default email
	// template. The template must be a template using our built-in customizations or a custom HTML email for
	// Magic links - Sign-up.
	SignupTemplateID string `json:"signup_template_id,omitempty"`
}

// InviteResponse: Response type for `Email.Invite`.
type InviteResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// EmailID: The unique ID of a specific email address.
	EmailID string `json:"email_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// LoginOrCreateResponse: Response type for `Email.LoginOrCreate`.
type LoginOrCreateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// EmailID: The unique ID of a specific email address.
	EmailID string `json:"email_id,omitempty"`
	// UserCreated: In `login_or_create` endpoints, this field indicates whether or not a User was just created.
	UserCreated bool `json:"user_created,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// RevokeInviteResponse: Response type for `Email.RevokeInvite`.
type RevokeInviteResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// SendResponse: Response type for `Email.Send`.
type SendResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// UserID: The unique ID of the affected User.
	UserID string `json:"user_id,omitempty"`
	// EmailID: The unique ID of a specific email address.
	EmailID string `json:"email_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

type InviteRequestLocale string

const (
	InviteRequestLocaleEn   InviteRequestLocale = "en"
	InviteRequestLocaleEs   InviteRequestLocale = "es"
	InviteRequestLocalePtbr InviteRequestLocale = "pt-br"
)

type LoginOrCreateRequestLocale string

const (
	LoginOrCreateRequestLocaleEn   LoginOrCreateRequestLocale = "en"
	LoginOrCreateRequestLocaleEs   LoginOrCreateRequestLocale = "es"
	LoginOrCreateRequestLocalePtbr LoginOrCreateRequestLocale = "pt-br"
)

type SendRequestLocale string

const (
	SendRequestLocaleEn   SendRequestLocale = "en"
	SendRequestLocaleEs   SendRequestLocale = "es"
	SendRequestLocalePtbr SendRequestLocale = "pt-br"
)
