package stytch

type Name struct {
	FirstName  string `json:"firstName,omitempty"`
	MiddleName string `json:"middleName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
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

type WebAuthnRegistration struct {
	WebAuthnRegistrationID string `json:"webauthn_registration_id,omitempty"`
	Domain                 string `json:"domain,omitempty"`
	UserAgent              string `json:"user_agent,omitempty"`
	Verified               bool   `json:"verified,omitempty"`
}

type UsersCreateParams struct {
	Email               string     `json:"email,omitempty"`
	PhoneNumber         string     `json:"phone_number,omitempty"`
	Name                Name       `json:"name,omitempty"`
	CreateUserAsPending bool       `json:"create_user_as_pending,omitempty"`
	Attributes          Attributes `json:"attributes,omitempty"`
}

type UsersCreateResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	EmailID    string `json:"email_id,omitempty"`
	PhoneID    string `json:"phone_id,omitempty"`
	Status     string `json:"status,omitempty"`
}

type UsersGetResponse struct {
	RequestID             string                 `json:"request_id,omitempty"`
	StatusCode            int                    `json:"status_code,omitempty"`
	UserID                string                 `json:"user_id,omitempty"`
	Name                  Name                   `json:"name,omitempty"`
	Emails                []Email                `json:"emails,omitempty"`
	PhoneNumbers          []PhoneNumber          `json:"phone_numbers,omitempty"`
	WebAuthnRegistrations []WebAuthnRegistration `json:"webauthn_registrations,omitempty"`
	Status                string                 `json:"status,omitempty"`
}

type UsersUpdateParams struct {
	Name         Name                `json:"name,omitempty"`
	Emails       []EmailString       `json:"emails,omitempty"`
	PhoneNumbers []PhoneNumberString `json:"phone_numbers,omitempty"`
	Attributes   Attributes          `json:"attributes,omitempty"`
}

type UsersUpdateResponse struct {
	RequestID    string        `json:"request_id,omitempty"`
	StatusCode   int           `json:"status_code,omitempty"`
	UserID       string        `json:"user_id,omitempty"`
	Emails       []Email       `json:"emails,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
}

type UsersDeleteResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
}

type UsersDeleteEmailResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
}

type UsersDeletePhoneNumberResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
}

type DeleteWebAuthnRegistrationResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
}

type PendingUsers struct {
	UserID       string        `json:"user_id,omitempty"`
	Name         Name          `json:"name,omitempty"`
	Emails       []Email       `json:"emails,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Status       string        `json:"status,omitempty"`
	InvitedAt    string        `json:"invited_at,omitempty"`
}

type UsersGetPendingParams struct {
	Limit           int32  `json:"limit,omitempty"`
	StartingAfterID string `json:"starting_after_id,omitempty"`
}

type UsersGetPendingResponse struct {
	RequestID       string         `json:"request_id,omitempty"`
	StatusCode      int            `json:"status_code,omitempty"`
	Users           []PendingUsers `json:"users,omitempty"`
	HasMore         bool           `json:"has_more,omitempty"`
	StartingAfterID string         `json:"starting_after_id,omitempty"`
	Total           int            `json:"total,omitempty"`
}
