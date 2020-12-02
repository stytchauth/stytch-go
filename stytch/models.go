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

type CreateUser struct {
	// The email to use for email magic links. This can be changed later via the update endpoint.
	Email      string     `json:"email"`
	Name       Name       `json:"name,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
}

type CreateUserResponse struct {
	RequestID string `json:"request_id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	// The id for the created email.
	EmailID string `json:"email_id,omitempty"`
}

type GetUserResponse struct {
	RequestID string  `json:"request_id,omitempty"`
	UserID    string  `json:"user_id,omitempty"`
	Name      Name    `json:"name,omitempty"`
	Emails    []Email `json:"emails,omitempty"`
}

type UpdateUser struct {
	Name Name `json:"name,omitempty"`
	// Multiple emails can exist for one user. Add additional emails via this endpoint.
	// To delete an email, use the delete endpoint.
	Emails     []Email    `json:"emails,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
}

type UpdateUserResponse struct {
	RequestID string  `json:"request_id,omitempty"`
	UserID    string  `json:"user_id,omitempty"`
	Emails    []Email `json:"emails,omitempty"`
}

type DeleteUserResponse struct {
	RequestID string `json:"request_id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
}

type DeleteEmailResponse struct {
	RequestID string `json:"request_id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	EmailID   string `json:"email_id,omitempty"`
}

type SendEmailVerification struct {
	UserID  string `json:"user_id"`
	EmailID string `json:"email_id,omitempty"`
	// The url the user clicks from the email magic link. This should be a url that your app
	// receives and parses and subsequently send an api request to authenticate the magic
	// link and log in the user.
	MagicLinkURL string `json:"magic_link_url,omitempty"`
	// Set the expiration for the email magic link, in minutes. By default, it expires in 1 hour.
	// The minimum expiration is 5 minutes and the maximum is 7 days (10080 mins).
	ExpirationMinutes int32 `json:"expiration_minutes,omitempty"`
}

type SendEmailVerificationResponse struct {
	RequestID string `json:"request_id,omitempty"`
}

type VerifyEmailResponse struct {
	RequestID string `json:"request_id,omitempty"`
	EmailID   string `json:"email_id,omitempty"`
}

type SendMagicLink struct {
	UserID string `json:"user_id"`
	// The method id for where to send the magic link, such as an email_id.
	MethodID string `json:"method_id"`
	// The url the user clicks from the email magic link. This should be a url that your app
	// receives and parses and subsequently send an api request to authenticate the magic
	// link and log in the user.
	MagicLinkURL string `json:"magic_link_url,omitempty"`
	// Set the expiration for the email magic link, in minutes. By default, it expires in 1 hour.
	// The minimum expiration is 5 minutes and the maximum is 7 days (10080 mins).
	ExpirationMinutes int32 `json:"expiration_minutes,omitempty"`
	// The template id to use for the magic link, for example the template_id that corresponds
	// to a specific email format.
	TemplateID string     `json:"template_id,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
}

type SendMagicLinkResponse struct {
	RequestID string `json:"request_id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
}

type SendMagicLinkByEmail struct {
	// The email the user enters to sign in with.
	Email string `json:"email"`
	// The url the user clicks from the email magic link. This should be a url that your
	// app receives and parses and subsequently send an api request to authenticate the
	// magic link and log in the user.
	MagicLinkURL string `json:"magic_link_url,omitempty"`
	// Set the expiration for the email magic link, in minutes. By default, it expires in 1 hour.
	// The minimum expiration is 5 minutes and the maximum is 7 days (10080 mins).
	ExpirationMinutes int32 `json:"expiration_minutes,omitempty"`
	// The template id to use for the magic link, for example the template_id
	// that corresponds to a specific email format.
	TemplateID string     `json:"template_id,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
}

type AuthenticateMagicLink struct {
	Options    Options    `json:"options,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
}

type AuthenticateMagicLinkResponse struct {
	RequestID string `json:"request_id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
}

type LoginOrCreateUser struct {
	// The email the user enters to login or sign up with.
	Email string `json:"email"`
	// The url the user clicks from the login email magic link. This should be a url that your
	// app receives and parses and subsequently send an api request to authenticate the
	// magic link and log in the user.
	LoginMagicLinkURL string `json:"login_magic_link_url,omitempty"`
	// The url the user clicks from the sign up email magic link. This should be a url that your
	// app receives and parses and subsequently send an api request to authenticate the
	// magic link and sign the user up.
	SignUpMagicLinkURL string `json:"signup_magic_link_url,omitempty"`
	// Set the expiration for the login email magic link, in minutes. By default, it expires in 1 hour.
	// The minimum expiration is 5 minutes and the maximum is 7 days (10080 mins).
	LoginExpirationMinutes int32 `json:"login_expiration_minutes,omitempty"`
	// Set the expiration for the sign up email magic link, in minutes.
	// By default, it expires in 1 week. The minimum expiration is 5 minutes and
	// the maximum is 7 days (10080 mins).
	SignUpExpirationMinutes int32 `json:"signup_expiration_minutes,omitempty"`
	// The template id to use for the login magic link, for example the template_id
	// that corresponds to a specific login email format.
	LoginTemplateID string `json:"login_template_id,omitempty"`
	// The template id to use for the sign up magic link, for example the template_id
	// that corresponds to a specific sign up email format.
	SignUpTemplateID string     `json:"signup_template_id,omitempty"`
	Attributes       Attributes `json:"attributes,omitempty"`
}

type LoginOrInviteByEmail struct {
	// The email the user enters to login or be invited with.
	Email string `json:"email"`
	// The url the user clicks from the login email magic link. This should be a url that your
	// app receives and parses and subsequently send an api request to authenticate the
	// magic link and log in the user.
	LoginMagicLinkURL string `json:"login_magic_link_url,omitempty"`
	// The url the user clicks from the invite email magic link. This should be a url that your
	// app receives and parses and subsequently send an api request to authenticate the
	// magic link and finish creating a users account.
	InviteMagicLinkURL string `json:"invite_magic_link_url,omitempty"`
	// Set the expiration for the login email magic link, in minutes. By default, it expires in 1 hour.
	// The minimum expiration is 5 minutes and the maximum is 7 days (10080 mins).
	LoginExpirationMinutes int32 `json:"login_expiration_minutes,omitempty"`
	// Set the expiration for the invite email magic link, in minutes.
	// By default, it expires in 1 week. The minimum expiration is 5 minutes
	// and the maximum is 7 days (10080 mins).
	InviteExpirationMinutes int32 `json:"invite_expiration_minutes,omitempty"`
	// The template id to use for the login magic link, for example the template_id
	// that corresponds to a specific login email format.
	LoginTemplateID string `json:"login_template_id,omitempty"`
	// The template id to use for the invite magic link, for example the template_id
	// that corresponds to a specific invite email format.
	InviteTemplateID string     `json:"invite_template_id,omitempty"`
	Attributes       Attributes `json:"attributes,omitempty"`
}

type LoginOrCreateResponse struct {
	RequestID   string `json:"request_id,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	EmailID     string `json:"email_id,omitempty"`
	UserCreated bool   `json:"user_created,omitempty"`
}
