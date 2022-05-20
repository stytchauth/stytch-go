package stytch

type OTPsAuthenticateParams struct {
	MethodID               string     `json:"method_id"`
	Code                   string     `json:"code"`
	Options                Options    `json:"options,omitempty"`
	Attributes             Attributes `json:"attributes,omitempty"`
	SessionToken           string     `json:"session_token,omitempty"`
	SessionDurationMinutes int32      `json:"session_duration_minutes,omitempty"`
	SessionJWT             string     `json:"session_jwt,omitempty"`
}

type OTPsAuthenticateResponse struct {
	RequestID    string  `json:"request_id,omitempty"`
	StatusCode   int     `json:"status_code,omitempty"`
	UserID       string  `json:"user_id,omitempty"`
	MethodID     string  `json:"method_id,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
	Session      Session `json:"session,omitempty"`
	SessionJWT   string  `json:"session_jwt,omitempty"`
	User         User    `jspon:"user,omitempty"`
}

// OTP - SMS
type OTPsSMSSendParams struct {
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

type OTPsSMSLoginOrCreateParams struct {
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

// OTP - WhatsApp
type OTPsWhatsAppSendParams struct {
	PhoneNumber       string     `json:"phone_number"`
	ExpirationMinutes int32      `json:"expiration_minutes,omitempty"`
	Attributes        Attributes `json:"attributes,omitempty"`
}

type OTPsWhatsAppSendResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	PhoneID    string `json:"phone_id,omitempty"`
}

type OTPsWhatsAppLoginOrCreateParams struct {
	PhoneNumber         string     `json:"phone_number"`
	ExpirationMinutes   int32      `json:"expiration_minutes,omitempty"`
	Attributes          Attributes `json:"attributes,omitempty"`
	CreateUserAsPending bool       `json:"create_user_as_pending,omitempty"`
}

type OTPsWhatsAppLoginOrCreateResponse struct {
	RequestID   string `json:"request_id,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	PhoneID     string `json:"phone_id,omitempty"`
	UserCreated bool   `json:"user_created,omitempty"`
}

// OTP - Email
type OTPsEmailSendParams struct {
	Email             string     `json:"email"`
	ExpirationMinutes int32      `json:"expiration_minutes,omitempty"`
	Attributes        Attributes `json:"attributes,omitempty"`
}

type OTPsEmailSendResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	EmailID    string `json:"email_id,omitempty"`
}

type OTPsEmailLoginOrCreateParams struct {
	Email               string     `json:"email"`
	ExpirationMinutes   int32      `json:"expiration_minutes,omitempty"`
	Attributes          Attributes `json:"attributes,omitempty"`
	CreateUserAsPending bool       `json:"create_user_as_pending,omitempty"`
}

type OTPsEmailLoginOrCreateResponse struct {
	RequestID   string `json:"request_id,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	EmailID     string `json:"email_id,omitempty"`
	UserCreated bool   `json:"user_created,omitempty"`
}
