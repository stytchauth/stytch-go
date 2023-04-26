package b2b

import "encoding/json"

type OrganizationCreateParams struct {
	OrganizationName     string         `json:"organization_name"`
	OrganizationSlug     string         `json:"organization_slug"`
	OrganizationLogoURL  string         `json:"organization_logo_url,omitempty"`
	TrustedMetadata      map[string]any `json:"trusted_metadata,omitempty"`
	SsoJitProvisioning   string         `json:"sso_jit_provisioning,omitempty"`
	EmailAllowedDomains  []string       `json:"email_allowed_domains,omitempty"`
	EmailJitProvisioning string         `json:"email_jit_provisioning,omitempty"`
	EmailInvites         string         `json:"email_invites,omitempty"`
	AuthMethods          string         `json:"auth_methods,omitempty"`
	AllowedAuthMethods   []string       `json:"allowed_auth_methods,omitempty"`
}
type OrganizationCreateResponse struct {
	RequestID    string       `json:"request_id,omitempty"`
	StatusCode   int          `json:"status_code,omitempty"`
	Organization Organization `json:"organization,omitempty"`
}

type OrganizationGetResponse struct {
	RequestID    string       `json:"request_id,omitempty"`
	StatusCode   int          `json:"status_code,omitempty"`
	Organization Organization `json:"organization,omitempty"`
}

type OrganizationUpdateParams struct {
	OrganizationName                     string         `json:"organization_name,omitempty"`
	OrganizationSlug                     string         `json:"organization_slug,omitempty"`
	OrganizationLogoURL                  string         `json:"organization_logo_url,omitempty"`
	TrustedMetadata                      map[string]any `json:"trusted_metadata,omitempty"`
	SsoDefaultConnectionID               string         `json:"sso_default_connection_id,omitempty"`
	SsoJitProvisioning                   string         `json:"sso_jit_provisioning,omitempty"`
	SsoJitProvisioningAllowedConnections []string       `json:"sso_jit_provisioning_allowed_connections,omitempty"`
	EmailAllowedDomains                  []string       `json:"email_allowed_domains,omitempty"`
	EmailJitProvisioning                 string         `json:"email_jit_provisioning,omitempty"`
	EmailInvites                         string         `json:"email_invites,omitempty"`
	AuthMethods                          string         `json:"auth_methods,omitempty"`
	AllowedAuthMethods                   []string       `json:"allowed_auth_methods,omitempty"`
}

type OrganizationUpdateResponse struct {
	RequestID    string       `json:"request_id,omitempty"`
	StatusCode   int          `json:"status_code,omitempty"`
	Organization Organization `json:"organization,omitempty"`
}

type OrganizationDeleteResponse struct {
	RequestID      string `json:"request_id,omitempty"`
	StatusCode     int    `json:"status_code,omitempty"`
	OrganizationID string `json:"organization_id,omitempty"`
}

type OrganizationSearchParams struct {
	Cursor string                  `json:"cursor,omitempty"`
	Limit  uint32                  `json:"limit,omitempty"`
	Query  OrganizationSearchQuery `json:"query,omitempty"`
}

type OrganizationSearchResponse struct {
	RequestID       string          `json:"request_id,omitempty"`
	Organizations   []Organization  `json:"organizations,omitempty"`
	ResultsMetadata ResultsMetadata `json:"results_metadata,omitempty"`
}

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

type OrganizationSearchQuery struct {
	Operator OrganizationSearchOperator `json:"operator,omitempty"`
	Operands []json.Marshaler           `json:"operands,omitempty"`
}

type OrganizationSearchOperator string

const (
	OrganizationSearchOperatorOR  OrganizationSearchOperator = "OR"
	OrganizationSearchOperatorAND OrganizationSearchOperator = "AND"
)

type ResultsMetadata struct {
	Total      int    `json:"total,omitempty"`
	NextCursor string `json:"next_cursor,omitempty"`
}
