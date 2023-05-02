package b2b

type DiscoveryListOrganizationsParams struct {
	IntermediateSessionToken string `json:"intermediate_session_token,omitempty"`
	SessionToken             string `json:"session_token,omitempty"`
	SessionJwt               string `json:"session_jwt,omitempty"`
}

type DiscoveryListOrganizationsResponse struct {
	RequestID               string                   `json:"request_id,omitempty"`
	StatusCode              int                      `json:"status_code,omitempty"`
	EmailAddress            string                   `json:"email_address,omitempty"`
	DiscoveredOrganizations []DiscoveredOrganization `json:"discovered_organizations,omitempty"`
}

type DiscoveryIntermediateSessionExchangeParams struct {
	IntermediateSessionToken string         `json:"intermediate_session_token"`
	OrganizationID           string         `json:"organization_id"`
	SessionDurationMinutes   int32          `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims      map[string]any `json:"session_custom_claims,omitempty"`
}

type DiscoveryIntermediateSessionExchangeResponse struct {
	RequestID     string        `json:"request_id,omitempty"`
	StatusCode    int           `json:"status_code,omitempty"`
	MemberID      string        `json:"member_id,omitempty"`
	MemberSession MemberSession `json:"member_session,omitempty"`
	SessionToken  string        `json:"session_token,omitempty"`
	SessionJWT    string        `json:"session_jwt,omitempty"`
	Member        Member        `json:"member,omitempty"`
	Organization  Organization  `json:"organization,omitempty"`
}

type DiscoveryOrganizationCreateParams struct {
	IntermediateSessionToken string         `json:"intermediate_session_token"`
	SessionDurationMinutes   int32          `json:"session_duration_minutes,omitempty"`
	SessionCustomClaims      map[string]any `json:"session_custom_claims,omitempty"`
	OrganizationName         string         `json:"organization_name,omitempty"`
	OrganizationSlug         string         `json:"organization_slug,omitempty"`
	OrganizationLogoURL      string         `json:"organization_logo_url,omitempty"`
	TrustedMetadata          map[string]any `json:"trusted_metadata,omitempty"`
	SsoJitProvisioning       string         `json:"sso_jit_provisioning,omitempty"`
	EmailAllowedDomains      []string       `json:"email_allowed_domains,omitempty"`
	EmailJitProvisioning     string         `json:"email_jit_provisioning,omitempty"`
	EmailInvites             string         `json:"email_invites,omitempty"`
	AuthMethods              string         `json:"auth_methods,omitempty"`
	AllowedAuthMethods       []string       `json:"allowed_auth_methods,omitempty"`
}

type DiscoveryOrganizationCreateResponse struct {
	RequestID     string        `json:"request_id,omitempty"`
	StatusCode    int           `json:"status_code,omitempty"`
	MemberID      string        `json:"member_id,omitempty"`
	MemberSession MemberSession `json:"member_session,omitempty"`
	SessionToken  string        `json:"session_token,omitempty"`
	SessionJWT    string        `json:"session_jwt,omitempty"`
	Member        Member        `json:"member,omitempty"`
	Organization  Organization  `json:"organization,omitempty"`
}

type DiscoveredOrganization struct {
	Organization        Organization `json:"organization,omitempty"`
	Membership          Membership   `json:"membership,omitempty"`
	MemberAuthenticated bool         `json:"member_authenticated,omitempty"`
}

type Membership struct {
	Type    string                 `json:"type,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Member  Member                 `json:"member,omitempty"`
}
