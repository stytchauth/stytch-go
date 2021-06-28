package stytch

type MagicLinksAuthenticateParams struct {
	Token      string     `json:"token,omitempty"`
	Options    Options    `json:"options,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
}

type MagicLinksAuthenticateResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	MethodID   string `json:"method_id,omitempty"`
}

// MAGIC LINK - EMAIL
type MagicLinksEmailSendParams struct {
	Email                   string     `json:"email"`
	LoginMagicLinkURL       string     `json:"login_magic_link_url"`
	SignupMagicLinkURL      string     `json:"signup_magic_link_url"`
	LoginExpirationMinutes  int32      `json:"login_expiration_minutes,omitempty"`
	SignupExpirationMinutes int32      `json:"signup_expiration_minutes,omitempty"`
	Attributes              Attributes `json:"attributes,omitempty"`
}

type MagicLinksEmailSendResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	EmailID    string `json:"email_id,omitempty"`
}

type MagicLinksEmailLoginOrCreateParams struct {
	Email                   string     `json:"email"`
	LoginMagicLinkURL       string     `json:"login_magic_link_url"`
	SignupMagicLinkURL      string     `json:"signup_magic_link_url"`
	LoginExpirationMinutes  int32      `json:"login_expiration_minutes,omitempty"`
	SignupExpirationMinutes int32      `json:"signup_expiration_minutes,omitempty"`
	CreateUserAsPending     bool       `json:"create_user_as_pending,omitempty"`
	Attributes              Attributes `json:"attributes,omitempty"`
}

type MagicLinksEmailLoginOrCreateResponse struct {
	RequestID   string `json:"request_id,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	EmailID     string `json:"email_id,omitempty"`
	UserCreated bool   `json:"user_created,omitempty"`
}

type MagicLinksEmailInviteParams struct {
	Email                   string     `json:"email"`
	InviteMagicLinkURL      string     `json:"invite_magic_link_url"`
	InviteExpirationMinutes int32      `json:"invite_expiration_minutes,omitempty"`
	Name                    Name       `json:"name,omitempty"`
	Attributes              Attributes `json:"attributes,omitempty"`
}

type MagicLinksEmailInviteResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	EmailID    string `json:"email_id,omitempty"`
}

type MagicLinksEmailRevokeInviteParams struct {
	Email string `json:"email"`
}

type MagicLinksEmailRevokeInviteResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
}
