package organizations

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

// ActiveSSOConnection:
type ActiveSSOConnection struct {
	// ConnectionID: Globally unique UUID that identifies a specific SSO `connection_id` for a Member.
	ConnectionID string `json:"connection_id,omitempty"`
	// DisplayName: A human-readable display name for the connection.
	DisplayName string `json:"display_name,omitempty"`
}

// CreateParams: Request type for `Organizations.Create`.
type CreateParams struct {
	// OrganizationName: The name of the Organization. Must be between 1 and 128 characters in length.
	OrganizationName string `json:"organization_name,omitempty"`
	// OrganizationSlug: The unique URL slug of the Organization. The slug only accepts alphanumeric characters
	// and the following reserved characters: `-` `.` `_` `~`. Must be between 2 and 128 characters in length.
	OrganizationSlug string `json:"organization_slug,omitempty"`
	// OrganizationLogoURL: The image URL of the Organization logo.
	OrganizationLogoURL string `json:"organization_logo_url,omitempty"`
	// TrustedMetadata: An arbitrary JSON object for storing application-specific data or
	// identity-provider-specific data.
	TrustedMetadata map[string]any `json:"trusted_metadata,omitempty"`
	// SSOJITProvisioning: The authentication setting that controls the JIT provisioning of Members when
	// authenticating via SSO. The accepted values are:
	//
	//   `ALL_ALLOWED` – new Members will be automatically provisioned upon successful authentication via any
	// of the Organization's `sso_active_connections`.
	//
	//   `RESTRICTED` – only new Members with SSO logins that comply with
	// `sso_jit_provisioning_allowed_connections` can be provisioned upon authentication.
	//
	//   `NOT_ALLOWED` – disable JIT provisioning via SSO.
	//
	SSOJITProvisioning string `json:"sso_jit_provisioning,omitempty"`
	// EmailAllowedDomains: An array of email domains that allow invites or JIT provisioning for new Members.
	// This list is enforced when either `email_invites` or `email_jit_provisioning` is set to `RESTRICTED`.
	//
	//
	//     Common domains such as `gmail.com` are not allowed. See the
	// [common email domains resource](https://stytch.com/docs/b2b/api/common-email-domains) for the full list.
	EmailAllowedDomains []string `json:"email_allowed_domains,omitempty"`
	// EmailJITProvisioning: The authentication setting that controls how a new Member can be provisioned by
	// authenticating via Email Magic Link. The accepted values are:
	//
	//   `RESTRICTED` – only new Members with verified emails that comply with `email_allowed_domains` can be
	// provisioned upon authentication via Email Magic Link.
	//
	//   `NOT_ALLOWED` – disable JIT provisioning via Email Magic Link.
	//
	EmailJITProvisioning string `json:"email_jit_provisioning,omitempty"`
	// EmailInvites: The authentication setting that controls how a new Member can be invited to an
	// organization by email. The accepted values are:
	//
	//   `ALL_ALLOWED` – any new Member can be invited to join via email.
	//
	//   `RESTRICTED` – only new Members with verified emails that comply with `email_allowed_domains` can be
	// invited via email.
	//
	//   `NOT_ALLOWED` – disable email invites.
	//
	EmailInvites string `json:"email_invites,omitempty"`
	// AuthMethods: The setting that controls which authentication methods can be used by Members of an
	// Organization. The accepted values are:
	//
	//   `ALL_ALLOWED` – the default setting which allows all authentication methods to be used.
	//
	//   `RESTRICTED` – only methods that comply with `allowed_auth_methods` can be used for authentication.
	// This setting does not apply to Members with `is_breakglass` set to `true`.
	//
	AuthMethods string `json:"auth_methods,omitempty"`
	// AllowedAuthMethods:
	//   An array of allowed authentication methods. This list is enforced when `auth_methods` is set to
	// `RESTRICTED`.
	//   The list's accepted values are: `sso`, `magic_link`, `password`, `google_oauth`, and `microsoft_oauth`.
	//
	AllowedAuthMethods []string `json:"allowed_auth_methods,omitempty"`
	// MFAPolicy: The setting that controls the MFA policy for all Members in the Organization. The accepted
	// values are:
	//
	//   `REQUIRED_FOR_ALL` – All Members within the Organization will be required to complete MFA every time
	// they wish to log in.
	//
	//   `OPTIONAL` – The default value. The Organization does not require MFA by default for all Members.
	// Members will be required to complete MFA only if their `mfa_enrolled` status is set to true.
	//
	MFAPolicy string `json:"mfa_policy,omitempty"`
}

// DeleteParams: Request type for `Organizations.Delete`.
type DeleteParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
}

// GetParams: Request type for `Organizations.Get`.
type GetParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
}

// Member:
type Member struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// MemberID: Globally unique UUID that identifies a specific Member. The `member_id` is critical to perform
	// operations on a Member, so be sure to preserve this value.
	MemberID string `json:"member_id,omitempty"`
	// EmailAddress: The email address of the Member.
	EmailAddress string `json:"email_address,omitempty"`
	// Status: The status of the Member. The possible values are: `pending`, `invited`, `active`, or `deleted`.
	Status string `json:"status,omitempty"`
	// Name: The name of the Member.
	Name string `json:"name,omitempty"`
	// SSORegistrations: An array of registered [SAML Connection](saml-connection-object) objects the Member
	// has authenticated with.
	SSORegistrations []SSORegistration `json:"sso_registrations,omitempty"`
	// IsBreakglass: Identifies the Member as a break glass user - someone who has permissions to authenticate
	// into an Organization by bypassing the Organization's settings. A break glass account is typically used
	// for emergency purposes to gain access outside of normal authentication procedures. Refer to the
	// [Organization object](organization-object) and its `auth_methods` and `allowed_auth_methods` fields for
	// more details.
	IsBreakglass bool `json:"is_breakglass,omitempty"`
	// MemberPasswordID: Globally unique UUID that identifies a Member's password.
	MemberPasswordID string `json:"member_password_id,omitempty"`
	// OAuthRegistrations: A list of OAuth registrations for this member.
	OAuthRegistrations []OAuthRegistration `json:"oauth_registrations,omitempty"`
	// EmailAddressVerified: Whether or not the Member's email address is verified.
	EmailAddressVerified bool `json:"email_address_verified,omitempty"`
	// MFAPhoneNumberVerified: Whether or not the Member's phone number is verified.
	MFAPhoneNumberVerified bool `json:"mfa_phone_number_verified,omitempty"`
	// MFAEnrolled: Sets whether the Member is enrolled in MFA. If true, the Member must complete an MFA step
	// whenever they wish to log in to their Organization. If false, the Member only needs to complete an MFA
	// step if the Organization's MFA policy is set to `REQUIRED_FOR_ALL`.
	MFAEnrolled bool `json:"mfa_enrolled,omitempty"`
	// MFAPhoneNumber: The Member's phone number. A Member may only have one phone number.
	MFAPhoneNumber string `json:"mfa_phone_number,omitempty"`
	// TrustedMetadata: An arbitrary JSON object for storing application-specific data or
	// identity-provider-specific data.
	TrustedMetadata map[string]any `json:"trusted_metadata,omitempty"`
	// UntrustedMetadata: An arbitrary JSON object of application-specific data. These fields can be edited
	// directly by the
	//   frontend SDK, and should not be used to store critical information. See the
	// [Metadata resource](https://stytch.com/docs/b2b/api/metadata)
	//   for complete field behavior details.
	UntrustedMetadata map[string]any `json:"untrusted_metadata,omitempty"`
}

// OAuthRegistration:
type OAuthRegistration struct {
	// ProviderType: Denotes the OAuth identity provider that the user has authenticated with, e.g. Google,
	// Microsoft, GitHub etc.
	ProviderType string `json:"provider_type,omitempty"`
	// ProviderSubject: The unique identifier for the User within a given OAuth provider. Also commonly called
	// the `sub` or "Subject field" in OAuth protocols.
	ProviderSubject string `json:"provider_subject,omitempty"`
	// MemberOAuthRegistrationID: The unique ID of an OAuth registration.
	MemberOAuthRegistrationID string `json:"member_oauth_registration_id,omitempty"`
	// ProfilePictureURL: If available, the `profile_picture_url` is a URL of the User's profile picture set in
	// OAuth identity the provider that the User has authenticated with, e.g. Google profile picture.
	ProfilePictureURL string `json:"profile_picture_url,omitempty"`
	// Locale: If available, the `locale` is the Member's locale set in the OAuth identity provider that the
	// user has authenticated with.
	Locale string `json:"locale,omitempty"`
}

// Organization:
type Organization struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// OrganizationName: The name of the Organization. Must be between 1 and 128 characters in length.
	OrganizationName string `json:"organization_name,omitempty"`
	// OrganizationLogoURL: The image URL of the Organization logo.
	OrganizationLogoURL string `json:"organization_logo_url,omitempty"`
	// OrganizationSlug: The unique URL slug of the Organization. The slug only accepts alphanumeric characters
	// and the following reserved characters: `-` `.` `_` `~`. Must be between 2 and 128 characters in length.
	OrganizationSlug string `json:"organization_slug,omitempty"`
	// SSOJITProvisioning: The authentication setting that controls the JIT provisioning of Members when
	// authenticating via SSO. The accepted values are:
	//
	//   `ALL_ALLOWED` – new Members will be automatically provisioned upon successful authentication via any
	// of the Organization's `sso_active_connections`.
	//
	//   `RESTRICTED` – only new Members with SSO logins that comply with
	// `sso_jit_provisioning_allowed_connections` can be provisioned upon authentication.
	//
	//   `NOT_ALLOWED` – disable JIT provisioning via SSO.
	//
	SSOJITProvisioning string `json:"sso_jit_provisioning,omitempty"`
	// SSOJITProvisioningAllowedConnections: An array of `connection_id`s that reference
	// [SAML Connection objects](https://stytch.com/docs/b2b/api/saml-connection-object).
	//   Only these connections will be allowed to JIT provision Members via SSO when `sso_jit_provisioning` is
	// set to `RESTRICTED`.
	SSOJITProvisioningAllowedConnections []string `json:"sso_jit_provisioning_allowed_connections,omitempty"`
	// SSOActiveConnections: An array of active
	// [SAML Connection references](https://stytch.com/docs/b2b/api/saml-connection-object).
	SSOActiveConnections []ActiveSSOConnection `json:"sso_active_connections,omitempty"`
	// EmailAllowedDomains: An array of email domains that allow invites or JIT provisioning for new Members.
	// This list is enforced when either `email_invites` or `email_jit_provisioning` is set to `RESTRICTED`.
	//
	//
	//     Common domains such as `gmail.com` are not allowed. See the
	// [common email domains resource](https://stytch.com/docs/b2b/api/common-email-domains) for the full list.
	EmailAllowedDomains []string `json:"email_allowed_domains,omitempty"`
	// EmailJITProvisioning: The authentication setting that controls how a new Member can be provisioned by
	// authenticating via Email Magic Link. The accepted values are:
	//
	//   `RESTRICTED` – only new Members with verified emails that comply with `email_allowed_domains` can be
	// provisioned upon authentication via Email Magic Link.
	//
	//   `NOT_ALLOWED` – disable JIT provisioning via Email Magic Link.
	//
	EmailJITProvisioning string `json:"email_jit_provisioning,omitempty"`
	// EmailInvites: The authentication setting that controls how a new Member can be invited to an
	// organization by email. The accepted values are:
	//
	//   `ALL_ALLOWED` – any new Member can be invited to join via email.
	//
	//   `RESTRICTED` – only new Members with verified emails that comply with `email_allowed_domains` can be
	// invited via email.
	//
	//   `NOT_ALLOWED` – disable email invites.
	//
	EmailInvites string `json:"email_invites,omitempty"`
	// AuthMethods: The setting that controls which authentication methods can be used by Members of an
	// Organization. The accepted values are:
	//
	//   `ALL_ALLOWED` – the default setting which allows all authentication methods to be used.
	//
	//   `RESTRICTED` – only methods that comply with `allowed_auth_methods` can be used for authentication.
	// This setting does not apply to Members with `is_breakglass` set to `true`.
	//
	AuthMethods string `json:"auth_methods,omitempty"`
	// AllowedAuthMethods:
	//   An array of allowed authentication methods. This list is enforced when `auth_methods` is set to
	// `RESTRICTED`.
	//   The list's accepted values are: `sso`, `magic_link`, `password`, `google_oauth`, and `microsoft_oauth`.
	//
	AllowedAuthMethods []string `json:"allowed_auth_methods,omitempty"`
	MFAPolicy          string   `json:"mfa_policy,omitempty"`
	// TrustedMetadata: An arbitrary JSON object for storing application-specific data or
	// identity-provider-specific data.
	TrustedMetadata map[string]any `json:"trusted_metadata,omitempty"`
	// SSODefaultConnectionID: The default connection used for SSO when there are multiple active connections.
	SSODefaultConnectionID string `json:"sso_default_connection_id,omitempty"`
}

// ResultsMetadata:
type ResultsMetadata struct {
	// Total: The total number of results returned by your search query.
	Total int32 `json:"total,omitempty"`
	// NextCursor: The `next_cursor` string is returned when your search result contains more than one page of
	// results. This value is passed into your next search call in the `cursor` field.
	NextCursor string `json:"next_cursor,omitempty"`
}

// SSORegistration:
type SSORegistration struct {
	// ConnectionID: Globally unique UUID that identifies a specific SSO `connection_id` for a Member.
	ConnectionID string `json:"connection_id,omitempty"`
	// ExternalID: The ID of the member given by the identity provider.
	ExternalID string `json:"external_id,omitempty"`
	// RegistrationID: The unique ID of an SSO Registration.
	RegistrationID string `json:"registration_id,omitempty"`
	// SSOAttributes: An object for storing SSO attributes brought over from the identity provider.
	SSOAttributes map[string]any `json:"sso_attributes,omitempty"`
}

// SearchParams: Request type for `Organizations.Search`.
type SearchParams struct {
	// Cursor: The `cursor` field allows you to paginate through your results. Each result array is limited to
	// 1000 results. If your query returns more than 1000 results, you will need to paginate the responses
	// using the `cursor`. If you receive a response that includes a non-null `next_cursor` in the
	// `results_metadata` object, repeat the search call with the `next_cursor` value set to the `cursor` field
	// to retrieve the next page of results. Continue to make search calls until the `next_cursor` in the
	// response is null.
	Cursor string `json:"cursor,omitempty"`
	// Limit: The number of search results to return per page. The default limit is 100. A maximum of 1000
	// results can be returned by a single search request. If the total size of your result set is greater than
	// one page size, you must paginate the response. See the `cursor` field.
	Limit uint32 `json:"limit,omitempty"`
	// Query: The optional query object contains the operator, i.e. `AND` or `OR`, and the operands that will
	// filter your results. Only an operator is required. If you include no operands, no filtering will be
	// applied. If you include no query object, it will return all Organizations with no filtering applied.
	Query *SearchQuery `json:"query,omitempty"`
}

// SearchQuery:
type SearchQuery struct {
	// Operator: The action to perform on the operands. The accepted value are:
	//
	//   `AND` – all the operand values provided must match.
	//
	//   `OR` – the operator will return any matches to at least one of the operand values you supply.
	Operator SearchQueryOperator `json:"operator,omitempty"`
	// Operands: An array of operand objects that contains all of the filters and values to apply to your
	// search query.
	Operands []map[string]any `json:"operands,omitempty"`
}

// UpdateParams: Request type for `Organizations.Update`.
type UpdateParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// OrganizationName: The name of the Organization. Must be between 1 and 128 characters in length.
	OrganizationName string `json:"organization_name,omitempty"`
	// OrganizationSlug: The unique URL slug of the Organization. The slug only accepts alphanumeric characters
	// and the following reserved characters: `-` `.` `_` `~`. Must be between 2 and 128 characters in length.
	OrganizationSlug string `json:"organization_slug,omitempty"`
	// OrganizationLogoURL: The image URL of the Organization logo.
	OrganizationLogoURL string `json:"organization_logo_url,omitempty"`
	// TrustedMetadata: An arbitrary JSON object for storing application-specific data or
	// identity-provider-specific data.
	TrustedMetadata map[string]any `json:"trusted_metadata,omitempty"`
	// SSODefaultConnectionID: The default connection used for SSO when there are multiple active connections.
	SSODefaultConnectionID string `json:"sso_default_connection_id,omitempty"`
	// SSOJITProvisioning: The authentication setting that controls the JIT provisioning of Members when
	// authenticating via SSO. The accepted values are:
	//
	//   `ALL_ALLOWED` – new Members will be automatically provisioned upon successful authentication via any
	// of the Organization's `sso_active_connections`.
	//
	//   `RESTRICTED` – only new Members with SSO logins that comply with
	// `sso_jit_provisioning_allowed_connections` can be provisioned upon authentication.
	//
	//   `NOT_ALLOWED` – disable JIT provisioning via SSO.
	//
	SSOJITProvisioning string `json:"sso_jit_provisioning,omitempty"`
	// SSOJITProvisioningAllowedConnections: An array of `connection_id`s that reference
	// [SAML Connection objects](https://stytch.com/docs/b2b/api/saml-connection-object).
	//   Only these connections will be allowed to JIT provision Members via SSO when `sso_jit_provisioning` is
	// set to `RESTRICTED`.
	SSOJITProvisioningAllowedConnections []string `json:"sso_jit_provisioning_allowed_connections,omitempty"`
	// EmailAllowedDomains: An array of email domains that allow invites or JIT provisioning for new Members.
	// This list is enforced when either `email_invites` or `email_jit_provisioning` is set to `RESTRICTED`.
	//
	//
	//     Common domains such as `gmail.com` are not allowed. See the
	// [common email domains resource](https://stytch.com/docs/b2b/api/common-email-domains) for the full list.
	EmailAllowedDomains []string `json:"email_allowed_domains,omitempty"`
	// EmailJITProvisioning: The authentication setting that controls how a new Member can be provisioned by
	// authenticating via Email Magic Link. The accepted values are:
	//
	//   `RESTRICTED` – only new Members with verified emails that comply with `email_allowed_domains` can be
	// provisioned upon authentication via Email Magic Link.
	//
	//   `NOT_ALLOWED` – disable JIT provisioning via Email Magic Link.
	//
	EmailJITProvisioning string `json:"email_jit_provisioning,omitempty"`
	// EmailInvites: The authentication setting that controls how a new Member can be invited to an
	// organization by email. The accepted values are:
	//
	//   `ALL_ALLOWED` – any new Member can be invited to join via email.
	//
	//   `RESTRICTED` – only new Members with verified emails that comply with `email_allowed_domains` can be
	// invited via email.
	//
	//   `NOT_ALLOWED` – disable email invites.
	//
	EmailInvites string `json:"email_invites,omitempty"`
	// AuthMethods: The setting that controls which authentication methods can be used by Members of an
	// Organization. The accepted values are:
	//
	//   `ALL_ALLOWED` – the default setting which allows all authentication methods to be used.
	//
	//   `RESTRICTED` – only methods that comply with `allowed_auth_methods` can be used for authentication.
	// This setting does not apply to Members with `is_breakglass` set to `true`.
	//
	AuthMethods string `json:"auth_methods,omitempty"`
	// AllowedAuthMethods:
	//   An array of allowed authentication methods. This list is enforced when `auth_methods` is set to
	// `RESTRICTED`.
	//   The list's accepted values are: `sso`, `magic_link`, `password`, `google_oauth`, and `microsoft_oauth`.
	//
	AllowedAuthMethods []string `json:"allowed_auth_methods,omitempty"`
	// MFAPolicy: The setting that controls the MFA policy for all Members in the Organization. The accepted
	// values are:
	//
	//   `REQUIRED_FOR_ALL` – All Members within the Organization will be required to complete MFA every time
	// they wish to log in.
	//
	//   `OPTIONAL` – The default value. The Organization does not require MFA by default for all Members.
	// Members will be required to complete MFA only if their `mfa_enrolled` status is set to true.
	//
	MFAPolicy string `json:"mfa_policy,omitempty"`
}

// CreateResponse: Response type for `Organizations.Create`.
type CreateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization Organization `json:"organization,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// DeleteResponse: Response type for `Organizations.Delete`.
type DeleteResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// GetResponse: Response type for `Organizations.Get`.
type GetResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization Organization `json:"organization,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// SearchResponse: Response type for `Organizations.Search`.
type SearchResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// Organizations: An array of [Organization objects](https://stytch.com/docs/b2b/api/organization-object).
	Organizations []Organization `json:"organizations,omitempty"`
	// ResultsMetadata: The search `results_metadata` object contains metadata relevant to your specific query
	// like `total` and `next_cursor`.
	ResultsMetadata ResultsMetadata `json:"results_metadata,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// UpdateResponse: Response type for `Organizations.Update`.
type UpdateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// Organization: The [Organization object](https://stytch.com/docs/b2b/api/organization-object).
	Organization Organization `json:"organization,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

type SearchQueryOperator string

const (
	SearchQueryOperatorOR  SearchQueryOperator = "OR"
	SearchQueryOperatorAND SearchQueryOperator = "AND"
)
