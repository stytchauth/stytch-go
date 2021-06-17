package stytch

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

/*
 * Structure for the custom type Name
 */
type Name struct {
	// The first name of the user.
	FirstName string `json:"firstName,omitempty"`
	// The middle name(s) of the user.
	MiddleName string `json:"middleName,omitempty"`
	// The last name of the user.
	LastName string `json:"lastName,omitempty"`
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

type CreateUser struct {
	// The email to use for email magic links. This can be changed later via the update endpoint.
	Email               string     `json:"email,omitempty"`
	PhoneNumber         string     `json:"phone_number,omitempty"`
	Name                Name       `json:"name,omitempty"`
	CreateUserAsPending bool       `json:"create_user_as_pending,omitempty"`
	Attributes          Attributes `json:"attributes,omitempty"`
}

type CreateUserResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	// The id for the created email.
	EmailID string `json:"email_id,omitempty"`
	PhoneID string `json:"phone_id,omitempty"`
	Status  string `json:"status,omitempty"`
}

type GetUserResponse struct {
	RequestID    string        `json:"request_id,omitempty"`
	StatusCode   int           `json:"status_code,omitempty"`
	UserID       string        `json:"user_id,omitempty"`
	Name         Name          `json:"name,omitempty"`
	Emails       []Email       `json:"emails,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Status       string        `json:"status,omitempty"`
}

type UpdateUser struct {
	Name Name `json:"name,omitempty"`
	// Multiple emails can exist for one user. Add additional emails via this endpoint.
	// To delete an email, use the delete endpoint.
	Emails       []EmailString       `json:"emails,omitempty"`
	PhoneNumbers []PhoneNumberString `json:"phone_numbers,omitempty"`
	Attributes   Attributes          `json:"attributes,omitempty"`
}

type UpdateUserResponse struct {
	RequestID    string        `json:"request_id,omitempty"`
	StatusCode   int           `json:"status_code,omitempty"`
	UserID       string        `json:"user_id,omitempty"`
	Emails       []Email       `json:"emails,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
}

type DeleteUserResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
}

type DeleteUserEmailResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
}

type DeleteUserPhoneNumberResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
}

type MagicLinksEmailSend struct {
	// The email the user enters to sign in with.
	Email string `json:"email"`
	// The url the user clicks from the login email magic link. This should be a url that your
	// app receives and parses and subsequently send an api request to authenticate the
	// magic link and log in the user.
	LoginMagicLinkURL string `json:"login_magic_link_url"`
	// The url the user clicks from the sign up email magic link. This should be a url that your
	// app receives and parses and subsequently send an api request to authenticate the
	// magic link and sign the user up.
	SignupMagicLinkURL string `json:"signup_magic_link_url"`
	// Set the expiration for the login email magic link, in minutes. By default, it expires in 1 hour.
	// The minimum expiration is 5 minutes and the maximum is 7 days (10080 mins).
	LoginExpirationMinutes int32 `json:"login_expiration_minutes,omitempty"`
	// Set the expiration for the sign up email magic link, in minutes.
	// By default, it expires in 1 week. The minimum expiration is 5 minutes and
	// the maximum is 7 days (10080 mins).
	SignupExpirationMinutes int32      `json:"signup_expiration_minutes,omitempty"`
	Attributes              Attributes `json:"attributes,omitempty"`
}

type MagicLinksEmailSendResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	EmailID    string `json:"email_id,omitempty"`
}

type MagicLinksAuthenticate struct {
	Token      string     `json:"token"`
	Options    Options    `json:"options,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
}

type MagicLinksAuthenticateResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	MethodID   string `json:"method_id,omitempty"`
}

type MagicLinksEmailLoginOrCreate struct {
	// The email the user enters to login or sign up with.
	Email string `json:"email"`
	// The url the user clicks from the login email magic link. This should be a url that your
	// app receives and parses and subsequently send an api request to authenticate the
	// magic link and log in the user.
	LoginMagicLinkURL string `json:"login_magic_link_url"`
	// The url the user clicks from the sign up email magic link. This should be a url that your
	// app receives and parses and subsequently send an api request to authenticate the
	// magic link and sign the user up.
	SignupMagicLinkURL string `json:"signup_magic_link_url"`
	// Set the expiration for the login email magic link, in minutes. By default, it expires in 1 hour.
	// The minimum expiration is 5 minutes and the maximum is 7 days (10080 mins).
	LoginExpirationMinutes int32 `json:"login_expiration_minutes,omitempty"`
	// Set the expiration for the sign up email magic link, in minutes.
	// By default, it expires in 1 week. The minimum expiration is 5 minutes and
	// the maximum is 7 days (10080 mins).
	SignupExpirationMinutes int32 `json:"signup_expiration_minutes,omitempty"`
	// Boolean flag for whether or not to save a user as pending vs active in Stytch.
	// Defaults to false. If true, users will be saved with status pending.
	CreateUserAsPending bool       `json:"create_user_as_pending,omitempty"`
	Attributes          Attributes `json:"attributes,omitempty"`
}

type MagicLinksEmailLoginOrCreateResponse struct {
	RequestID   string `json:"request_id,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	EmailID     string `json:"email_id,omitempty"`
	UserCreated bool   `json:"user_created,omitempty"`
}

type MagicLinksEmailInvite struct {
	// The email the user enters to be invited with.
	Email string `json:"email"`
	// The url the user clicks from the invite email magic link. This should be a url that your
	// app receives and parses and subsequently send an api request to authenticate the
	// magic link and log in the user.
	InviteMagicLinkURL string `json:"magic_link_url"`
	// Set the expiration for the invite email magic link, in minutes. By default,
	// it expires in 1 hour. The minimum expiration is 5 minutes and the maximum
	// is 7 days (10080 mins).
	InviteExpirationMinutes int32      `json:"login_expiration_minutes,omitempty"`
	Name                    Name       `json:"name,omitempty"`
	Attributes              Attributes `json:"attributes,omitempty"`
}

type MagicLinksEmailInviteResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	EmailID    string `json:"email_id,omitempty"`
}

type MagicLinksEmailRevokeInvite struct {
	// The email of the user who's invite should be revoked.
	Email string `json:"email"`
}

type MagicLinksEmailRevokeInviteResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
}

type PendingUsers struct {
	UserID       string        `json:"user_id,omitempty"`
	Name         Name          `json:"name,omitempty"`
	Emails       []Email       `json:"emails,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Status       string        `json:"status,omitempty"`
	InvitedAt    string        `json:"invited_at,omitempty"`
}

type GetPendingUsers struct {
	Limit           int32  `json:"limit,omitempty"`
	StartingAfterID string `json:"starting_after_id,omitempty"`
}

type GetPendingUsersResponse struct {
	RequestID       string         `json:"request_id,omitempty"`
	StatusCode      int            `json:"status_code,omitempty"`
	Users           []PendingUsers `json:"users,omitempty"`
	HasMore         bool           `json:"has_more,omitempty"`
	StartingAfterID string         `json:"starting_after_id,omitempty"`
	Total           int            `json:"total,omitempty"`
}

type OTPsSMSSend struct {
	PhoneNumber       string     `json:"phone_number"`
	ExpirationMinutes int32      `json:"expiration_minutes,omitempty"`
	Attributes        Attributes `json:"attributes,omitempty"`
}

type OTPsSMSSendResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	PhoneID    string `json:"phone_id,omitempty"`
}

type OTPsSMSLoginOrCreate struct {
	PhoneNumber         string     `json:"phone_number"`
	ExpirationMinutes   int32      `json:"expiration_minutes,omitempty"`
	Attributes          Attributes `json:"attributes,omitempty"`
	CreateUserAsPending bool       `json:"create_user_as_pending,omitempty"`
}

type OTPsSMSLoginOrCreateResponse struct {
	RequestID   string `json:"request_id,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	PhoneID     string `json:"phone_id,omitempty"`
	UserCreated bool   `json:"user_created,omitempty"`
}

type OTPsAuthenticate struct {
	MethodID   string     `json:"method_id"`
	Code       string     `json:"code"`
	Options    Options    `json:"options,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
}

type OTPsAuthenticateResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	MethodID   string `json:"method_id,omitempty"`
}
