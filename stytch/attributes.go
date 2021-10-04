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

type EmailFactor struct {
	EmailID      string `json:"email_id,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
}

type PhoneNumberFactor struct {
	PhoneID     string `json:"phone_id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type GoogleOAuthFactor struct {
	ID              string `json:"id,omitempty"`
	EmailID         string `json:"email_id,omitempty"`
	ProviderSubject string `json:"provider_subject,omitempty"`
}

type AuthenticationFactor struct {
	Type              string            `json:"type,omitempty"`
	DeliveryMethod    string            `json:"delivery_method,omitempty"`
	EmailFactor       EmailFactor       `json:"email_factor,omitempty"`
	PhoneNumberFactor PhoneNumberFactor `json:"phone_number_factor,omitempty"`
	GoogleOAuthFactor GoogleOAuthFactor `json:"google_oauth_factor,omitempty"`
}

/*
 * Structure for the custom type Session
 */
type Session struct {
	SessionID             string                  `json:"session_id,omitempty"`
	UserID                string                  `json:"user_id,omitempty"`
	StartedAt             string                  `json:"started_at,omitempty"`
	LastAccessedAt        string                  `json:"last_accessed_at,omitempty"`
	ExpiresAt             string                  `json:"expires_at,omitempty"`
	Attributes            Attributes              `json:"attributes,omitempty"`
	AuthenticationFactors []*AuthenticationFactor `json:"authentication_factors,omitempty"`
}
