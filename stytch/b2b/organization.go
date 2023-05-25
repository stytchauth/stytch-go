package b2b

import (
	"encoding/json"
)

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
	StatusCode      int             `json:"status_code,omitempty"`
	Organizations   []Organization  `json:"organizations,omitempty"`
	ResultsMetadata ResultsMetadata `json:"results_metadata,omitempty"`
}

// MEMBERS

type OrganizationMemberCreateParams struct {
	EmailAddress          string         `json:"email_address,omitempty"`
	Name                  string         `json:"name,omitempty"`
	TrustedMetadata       map[string]any `json:"trusted_metadata,omitempty"`
	UntrustedMetadata     map[string]any `json:"untrusted_metadata,omitempty"`
	CreateMemberAsPending bool           `json:"create_member_as_pending,omitempty"`
	IsBreakglass          bool           `json:"is_breakglass,omitempty"`
}

type OrganizationMemberCreateResponse struct {
	RequestID    string       `json:"request_id,omitempty"`
	StatusCode   int          `json:"status_code,omitempty"`
	MemberID     string       `json:"member_id,omitempty"`
	Member       Member       `json:"member,omitempty"`
	Organization Organization `json:"organization,omitempty"`
}

type OrganizationMemberUpdateParams struct {
	Name              string         `json:"name,omitempty"`
	TrustedMetadata   map[string]any `json:"trusted_metadata,omitempty"`
	UntrustedMetadata map[string]any `json:"untrusted_metadata,omitempty"`
	IsBreakglass      bool           `json:"is_breakglass,omitempty"`
}
type OrganizationMemberUpdateResponse struct {
	RequestID    string       `json:"request_id,omitempty"`
	StatusCode   int          `json:"status_code,omitempty"`
	MemberID     string       `json:"member_id,omitempty"`
	Member       Member       `json:"member,omitempty"`
	Organization Organization `json:"organization,omitempty"`
}

type OrganizationMemberDeleteResponse struct {
	RequestID  string `json:"request_id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	MemberID   string `json:"member_id,omitempty"`
}

type OrganizationMemberSearchParams struct {
	Cursor          string            `json:"cursor,omitempty"`
	Limit           uint32            `json:"limit,omitempty"`
	Query           MemberSearchQuery `json:"query,omitempty"`
	OrganizationIds []string          `json:"organization_ids,omitempty"`
}

type OrganizationMemberSearchResponse struct {
	RequestID       string                  `json:"request_id,omitempty"`
	StatusCode      int                     `json:"status_code,omitempty"`
	Members         []Member                `json:"members,omitempty"`
	Organizations   map[string]Organization `json:"organizations,omitempty"`
	ResultsMetadata ResultsMetadata         `json:"results_metadata,omitempty"`
}

type OrganizationMemberGetParams struct {
	MemberID     string `json:"member_id,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
}

type OrganizationMemberGetResponse struct {
	RequestID    string       `json:"request_id,omitempty"`
	StatusCode   int          `json:"status_code,omitempty"`
	MemberID     string       `json:"member_id,omitempty"`
	Member       Member       `json:"member,omitempty"`
	Organization Organization `json:"organization,omitempty"`
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
	Operator SearchOperator   `json:"operator,omitempty"`
	Operands []json.Marshaler `json:"operands,omitempty"`
}

type SearchOperator string

const (
	SearchOperatorOR  SearchOperator = "OR"
	SearchOperatorAND SearchOperator = "AND"
)

type ResultsMetadata struct {
	Total      int    `json:"total,omitempty"`
	NextCursor string `json:"next_cursor,omitempty"`
}

type MemberSearchQuery struct {
	Operator SearchOperator   `json:"operator,omitempty"`
	Operands []json.Marshaler `json:"operands,omitempty"`
}

func marshalFilter(filterName string, filterValue interface{}) ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"filter_name":  filterName,
		"filter_value": filterValue,
	})
}

// Organization Search Filters

type OrganizationSearchOrganizationIDsQuery struct {
	OrganizationIDs []string
}

func (q OrganizationSearchOrganizationIDsQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("organization_ids", q.OrganizationIDs)
}

type OrganizationSearchOrganizationSlugsQuery struct {
	OrganizationSlugs []string
}

func (q OrganizationSearchOrganizationSlugsQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("organization_slugs", q.OrganizationSlugs)
}

type OrganizationSearchOrganizationNameFuzzyQuery struct {
	OrganizationNameFuzzy string
}

func (q OrganizationSearchOrganizationNameFuzzyQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("organization_name_fuzzy", q.OrganizationNameFuzzy)
}

type OrganizationSearchOrganizationSlugFuzzyQuery struct {
	OrganizationSlugFuzzy string
}

func (q OrganizationSearchOrganizationSlugFuzzyQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("organization_slug_fuzzy", q.OrganizationSlugFuzzy)
}

type OrganizationSearchMemberEmailsQuery struct {
	MemberEmails []string
}

func (q OrganizationSearchMemberEmailsQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("member_emails", q.MemberEmails)
}

type OrganizationSearchMemberEmailFuzzyQuery struct {
	MemberEmailFuzzy string
}

func (q OrganizationSearchMemberEmailFuzzyQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("member_email_fuzzy", q.MemberEmailFuzzy)
}

type OrganizationSearchAllowedDomainsQuery struct {
	AllowedDomains []string
}

func (q OrganizationSearchAllowedDomainsQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("allowed_domains", q.AllowedDomains)
}

type OrganizationSearchAllowedDomainFuzzyQuery struct {
	AllowedDomainFuzzy string
}

func (q OrganizationSearchAllowedDomainFuzzyQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("allowed_domain_fuzzy", q.AllowedDomainFuzzy)
}

// Member Search Filters

type MemberSearchMemberIDsQuery struct {
	MemberIDs []string
}

func (q MemberSearchMemberIDsQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("member_ids", q.MemberIDs)
}

type MemberSearchMemberEmailsQuery struct {
	MemberEmails []string
}

func (q MemberSearchMemberEmailsQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("member_emails", q.MemberEmails)
}

type MemberSearchMemberEmailFuzzyQuery struct {
	MemberEmailFuzzy string
}

func (q MemberSearchMemberEmailFuzzyQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("member_email_fuzzy", q.MemberEmailFuzzy)
}

type MemberSearchMemberIsBreakglassQuery struct {
	MemberIsBreakglass bool
}

func (q MemberSearchMemberIsBreakglassQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("member_is_breakglass", q.MemberIsBreakglass)
}

type MemberSearchMemberPasswordExistsQuery struct {
	MemberPasswordExists bool
}

func (q MemberSearchMemberPasswordExistsQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("member_password_exists", q.MemberPasswordExists)
}

type MemberSearchOrganizationIDQuery struct {
	OrganizationID string
}

func (q MemberSearchOrganizationIDQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("organization_id", q.OrganizationID)
}

type MemberSearchOrganizationSlugQuery struct {
	OrganizationSlug string
}

func (q MemberSearchOrganizationSlugQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("organization_slug", q.OrganizationSlug)
}

type MemberSearchOrganizationSlugFuzzyQuery struct {
	OrganizationSlugFuzzy string
}

func (q MemberSearchOrganizationSlugFuzzyQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("organization_slug_fuzzy", q.OrganizationSlugFuzzy)
}

type MemberSearchStatusQuery struct {
	Status string
}

func (q MemberSearchStatusQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("status", q.Status)
}

type MemberSearchStatusesQuery struct {
	Statuses []string
}

func (q MemberSearchStatusesQuery) MarshalJSON() ([]byte, error) {
	return marshalFilter("statuses", q.Statuses)
}
