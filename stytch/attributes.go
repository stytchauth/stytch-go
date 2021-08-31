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
 * Structure for the custom type Session
 */
type Session struct {
	SessionID      string     `json:"session_id,omitempty"`
	UserID         string     `json:"user_id,omitempty"`
	StartedAt      string     `json:"started_at,omitempty"`
	LastAccessedAt string     `json:"last_accessed_at,omitempty"`
	ExpiresAt      string     `json:"expires_at,omitempty"`
	Attributes     Attributes `json:"attributes,omitempty"`
}
