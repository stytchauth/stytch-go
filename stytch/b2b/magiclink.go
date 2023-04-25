package b2b

type MagicLinksAuthenticateParams struct {
	MagicLinksToken        string                 `json:"magic_links_token"`
	SessionToken           string                 `json:"session_token,omitempty"`
	SessionJWT             string                 `json:"session_jwt,omitempty"`
	SessionDurationMinutes int32                  `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims    map[string]interface{} `json:"session_custom_claims,omitempty"`
	PKCECodeVerifier       string                 `json:"pkce_code_verifier,omitempty"`
}

type MagicLinksAuthenticateResponse struct {
	RequestID      string        `json:"request_id,omitempty"`
	StatusCode     int           `json:"status_code,omitempty"`
	MemberID       string        `json:"member_id,omitempty"`
	MethodID       string        `json:"method_id,omitempty"`
	ResetSessions  bool          `json:"reset_sessions,omitempty"`
	OrganizationID string        `json:"organization_id,omitempty"`
	Member         Member        `json:"member,omitempty"`
	SessionToken   string        `json:"session_token,omitempty"`
	SessionJWT     string        `json:"session_jwt,omitempty"`
	MemberSession  MemberSession `json:"member_session,omitempty"`
	Organization   Organization  `json:"organization,omitempty"`
}

// MAGIC LINK - EMAIL
type MagicLinksEmailLoginOrSignupParams struct {
	OrganizationID         string `json:"organization_id"`
	EmailAddress           string `json:"email_address"`
	LoginMagicLinkURL      string `json:"login_magic_link_url,omitempty"`
	SignupMagicLinkURL     string `json:"signup_magic_link_url,omitempty"`
	LoginExpirationMinutes int32  `json:"login_expiration_minutes,omitempty"`
	PKCECodeChallenge      string `json:"pkce_code_challenge,omitempty"`
	LoginTemplateID        string `json:"login_template_id,omitempty"`
	SignupTemplateID       string `json:"signup_template_id,omitempty"`
	Locale                 string `json:"locale,omitempty"`
}

type MagicLinksEmailLoginOrSignupResponse struct {
	RequestID     string       `json:"request_id,omitempty"`
	StatusCode    int          `json:"status_code,omitempty"`
	MemberID      string       `json:"member_id,omitempty"`
	MemberCreated bool         `json:"member_created,omitempty"`
	Member        Member       `json:"member,omitempty"`
	Organization  Organization `json:"organization,omitempty"`
}
