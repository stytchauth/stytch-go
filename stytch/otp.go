package stytch

type OTPsAuthenticateParams struct {
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
