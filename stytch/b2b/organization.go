package b2b

type Organization struct {
	OrganizationID      string                 `json:"organization_id,omitempty"`
	OrganizationName    string                 `json:"organization_name,omitempty"`
	OrganizationLogoURL string                 `json:"organization_logo_url,omitempty"`
	OrganizationSlug    string                 `json:"organization_slug,omitempty"`
	TrustedMetadata     map[string]interface{} `json:"trusted_metadata,omitempty"`

	SSODefaultConnectionID             string                `json:"sso_default_connection_id,omitempty"`
	SSOJITProvisioning                 string                `json:"sso_jit_provisioning,omitempty"`
	SSOJITProvisioningAllowConnections []string              `json:"sso_jit_provisioning_allowed_connections,omitempty"`
	SSOActiveConnections               []ActiveSSOConnection `json:"sso_active_connections,omitempty"`

	EmailAllowedDomains  []string `json:"email_allowed_domains,omitempty"`
	EmailJITProvisioning string   `json:"email_jit_provisioning,omitempty"`
	EmailInvites         string   `json:"email_invites,omitempty"`

	AuthMethods        string   `json:"auth_methods,omitempty"`
	AllowedAuthMethods []string `json:"allowed_auth_methods,omitempty"`
}

type Member struct {
	OrganizationID    string                 `json:"organization_id,omitempty"`
	MemberID          string                 `json:"member_id,omitempty"`
	EmailAddress      string                 `json:"email_address,omitempty"`
	Status            string                 `json:"status,omitempty"`
	Name              string                 `json:"name,omitempty"`
	SSORegistration   []SSORegistrations     `json:"sso_registrations,omitempty"`
	TrustedMetadata   map[string]interface{} `json:"trusted_metadata,omitempty"`
	UntrustedMetadata map[string]interface{} `json:"untrusted_metadata,omitempty"`
	IsBreakglass      bool                   `json:"is_breakglass,omitempty"`
	MemberPasswordID  string                 `json:"member_password_id,omitempty"`
}

type ActiveSSOConnection struct {
	ConnectionID string `json:"connection_id,omitempty"`
	DisplayName  string `json:"display_name,omitempty"`
}

type SSORegistrations struct {
	ConnectionID   string                 `json:"connection_id,omitempty"`
	ExternalID     string                 `json:"external_id,omitempty"`
	RegistrationID string                 `json:"registration_id,omitempty"`
	SSOAttributes  map[string]interface{} `json:"sso_attributes,omitempty"`
}
