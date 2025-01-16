package saml

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v17/stytch/b2b/sso"
	"github.com/stytchauth/stytch-go/v17/stytch/methodoptions"
)

// CreateConnectionParams: Request type for `SAML.CreateConnection`.
type CreateConnectionParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// DisplayName: A human-readable display name for the connection.
	DisplayName string `json:"display_name,omitempty"`
	// IdentityProvider: Name of the IdP. Enum with possible values: `classlink`, `cyberark`, `duo`,
	// `google-workspace`, `jumpcloud`, `keycloak`, `miniorange`, `microsoft-entra`, `okta`, `onelogin`,
	// `pingfederate`, `rippling`, `salesforce`, `shibboleth`, or `generic`.
	//
	// Specifying a known provider allows Stytch to handle any provider-specific logic.
	IdentityProvider *CreateConnectionRequestIdentityProvider `json:"identity_provider,omitempty"`
}

// DeleteVerificationCertificateParams: Request type for `SAML.DeleteVerificationCertificate`.
type DeleteVerificationCertificateParams struct {
	// OrganizationID: The organization ID that the SAML connection belongs to.
	OrganizationID string `json:"organization_id,omitempty"`
	// ConnectionID: The ID of the SAML connection.
	ConnectionID string `json:"connection_id,omitempty"`
	// CertificateID: The ID of the certificate to be deleted.
	CertificateID string `json:"certificate_id,omitempty"`
}

// UpdateByURLParams: Request type for `SAML.UpdateByURL`.
type UpdateByURLParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// ConnectionID: Globally unique UUID that identifies a specific SSO `connection_id` for a Member.
	ConnectionID string `json:"connection_id,omitempty"`
	// MetadataURL: A URL that points to the IdP metadata. This will be provided by the IdP.
	MetadataURL string `json:"metadata_url,omitempty"`
}

// UpdateConnectionParams: Request type for `SAML.UpdateConnection`.
type UpdateConnectionParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// ConnectionID: Globally unique UUID that identifies a specific SSO `connection_id` for a Member.
	ConnectionID string `json:"connection_id,omitempty"`
	// IdpEntityID: A globally unique name for the IdP. This will be provided by the IdP.
	IdpEntityID string `json:"idp_entity_id,omitempty"`
	// DisplayName: A human-readable display name for the connection.
	DisplayName string `json:"display_name,omitempty"`
	// AttributeMapping: An object that represents the attributes used to identify a Member. This object will
	// map the IdP-defined User attributes to Stytch-specific values. Required attributes: `email` and one of
	// `full_name` or `first_name` and `last_name`.
	AttributeMapping map[string]any `json:"attribute_mapping,omitempty"`
	// X509Certificate: A certificate that Stytch will use to verify the sign-in assertion sent by the IdP, in
	// [PEM](https://en.wikipedia.org/wiki/Privacy-Enhanced_Mail) format. See our
	// [X509 guide](https://stytch.com/docs/b2b/api/saml-certificates) for more info.
	X509Certificate string `json:"x509_certificate,omitempty"`
	// IdpSSOURL: The URL for which assertions for login requests will be sent. This will be provided by the
	// IdP.
	IdpSSOURL string `json:"idp_sso_url,omitempty"`
	// SAMLConnectionImplicitRoleAssignments: All Members who log in with this SAML connection will implicitly
	// receive the specified Roles. See the
	// [RBAC guide](https://stytch.com/docs/b2b/guides/rbac/role-assignment) for more information about role
	// assignment.
	SAMLConnectionImplicitRoleAssignments []*sso.SAMLConnectionImplicitRoleAssignment `json:"saml_connection_implicit_role_assignments,omitempty"`
	// SAMLGroupImplicitRoleAssignments: Defines the names of the SAML groups
	//  that grant specific role assignments. For each group-Role pair, if a Member logs in with this SAML
	// connection and
	//  belongs to the specified SAML group, they will be granted the associated Role. See the
	//  [RBAC guide](https://stytch.com/docs/b2b/guides/rbac/role-assignment) for more information about role
	// assignment. Before adding any group implicit role assignments, you must add a "groups" key to your SAML
	// connection's
	//          `attribute_mapping`. Make sure that your IdP is configured to correctly send the group
	// information.
	SAMLGroupImplicitRoleAssignments []*sso.SAMLGroupImplicitRoleAssignment `json:"saml_group_implicit_role_assignments,omitempty"`
	// AlternativeAudienceURI: An alternative URL to use for the Audience Restriction. This value can be used
	// when you wish to migrate an existing SAML integration to Stytch with zero downtime. Read our
	// [SSO migration guide](https://stytch.com/docs/b2b/guides/migrations/additional-migration-considerations)
	// for more info.
	AlternativeAudienceURI string `json:"alternative_audience_uri,omitempty"`
	// IdentityProvider: Name of the IdP. Enum with possible values: `classlink`, `cyberark`, `duo`,
	// `google-workspace`, `jumpcloud`, `keycloak`, `miniorange`, `microsoft-entra`, `okta`, `onelogin`,
	// `pingfederate`, `rippling`, `salesforce`, `shibboleth`, or `generic`.
	//
	// Specifying a known provider allows Stytch to handle any provider-specific logic.
	IdentityProvider *UpdateConnectionRequestIdentityProvider `json:"identity_provider,omitempty"`
}

// CreateConnectionRequestOptions:
type CreateConnectionRequestOptions struct {
	// Authorization: Optional authorization object.
	// Pass in an active Stytch Member session token or session JWT and the request
	// will be run using that member's permissions.
	Authorization methodoptions.Authorization `json:"authorization,omitempty"`
}

func (o *CreateConnectionRequestOptions) AddHeaders(headers map[string][]string) map[string][]string {
	headers = o.Authorization.AddHeaders(headers)
	return headers
}

// DeleteVerificationCertificateRequestOptions:
type DeleteVerificationCertificateRequestOptions struct {
	// Authorization: Optional authorization object.
	// Pass in an active Stytch Member session token or session JWT and the request
	// will be run using that member's permissions.
	Authorization methodoptions.Authorization `json:"authorization,omitempty"`
}

func (o *DeleteVerificationCertificateRequestOptions) AddHeaders(headers map[string][]string) map[string][]string {
	headers = o.Authorization.AddHeaders(headers)
	return headers
}

// UpdateByURLRequestOptions:
type UpdateByURLRequestOptions struct {
	// Authorization: Optional authorization object.
	// Pass in an active Stytch Member session token or session JWT and the request
	// will be run using that member's permissions.
	Authorization methodoptions.Authorization `json:"authorization,omitempty"`
}

func (o *UpdateByURLRequestOptions) AddHeaders(headers map[string][]string) map[string][]string {
	headers = o.Authorization.AddHeaders(headers)
	return headers
}

// UpdateConnectionRequestOptions:
type UpdateConnectionRequestOptions struct {
	// Authorization: Optional authorization object.
	// Pass in an active Stytch Member session token or session JWT and the request
	// will be run using that member's permissions.
	Authorization methodoptions.Authorization `json:"authorization,omitempty"`
}

func (o *UpdateConnectionRequestOptions) AddHeaders(headers map[string][]string) map[string][]string {
	headers = o.Authorization.AddHeaders(headers)
	return headers
}

// CreateConnectionResponse: Response type for `SAML.CreateConnection`.
type CreateConnectionResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// Connection: The `SAML Connection` object affected by this API call. See the
	// [SAML Connection Object](https://stytch.com/docs/b2b/api/saml-connection-object) for complete response
	// field details.
	Connection *sso.SAMLConnection `json:"connection,omitempty"`
}

// DeleteVerificationCertificateResponse: Response type for `SAML.DeleteVerificationCertificate`.
type DeleteVerificationCertificateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// CertificateID: The ID of the certificate that was deleted.
	CertificateID string `json:"certificate_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// UpdateByURLResponse: Response type for `SAML.UpdateByURL`.
type UpdateByURLResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// Connection: The `SAML Connection` object affected by this API call. See the
	// [SAML Connection Object](https://stytch.com/docs/b2b/api/saml-connection-object) for complete response
	// field details.
	Connection *sso.SAMLConnection `json:"connection,omitempty"`
}

// UpdateConnectionResponse: Response type for `SAML.UpdateConnection`.
type UpdateConnectionResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// Connection: The `SAML Connection` object affected by this API call. See the
	// [SAML Connection Object](https://stytch.com/docs/b2b/api/saml-connection-object) for complete response
	// field details.
	Connection *sso.SAMLConnection `json:"connection,omitempty"`
}

type CreateConnectionRequestIdentityProvider string

const (
	CreateConnectionRequestIdentityProviderClasslink       CreateConnectionRequestIdentityProvider = "classlink"
	CreateConnectionRequestIdentityProviderCyberark        CreateConnectionRequestIdentityProvider = "cyberark"
	CreateConnectionRequestIdentityProviderDuo             CreateConnectionRequestIdentityProvider = "duo"
	CreateConnectionRequestIdentityProviderGeneric         CreateConnectionRequestIdentityProvider = "generic"
	CreateConnectionRequestIdentityProviderGoogleworkspace CreateConnectionRequestIdentityProvider = "google-workspace"
	CreateConnectionRequestIdentityProviderJumpcloud       CreateConnectionRequestIdentityProvider = "jumpcloud"
	CreateConnectionRequestIdentityProviderKeycloak        CreateConnectionRequestIdentityProvider = "keycloak"
	CreateConnectionRequestIdentityProviderMiniorange      CreateConnectionRequestIdentityProvider = "miniorange"
	CreateConnectionRequestIdentityProviderMicrosoftentra  CreateConnectionRequestIdentityProvider = "microsoft-entra"
	CreateConnectionRequestIdentityProviderOkta            CreateConnectionRequestIdentityProvider = "okta"
	CreateConnectionRequestIdentityProviderOnelogin        CreateConnectionRequestIdentityProvider = "onelogin"
	CreateConnectionRequestIdentityProviderPingfederate    CreateConnectionRequestIdentityProvider = "pingfederate"
	CreateConnectionRequestIdentityProviderRippling        CreateConnectionRequestIdentityProvider = "rippling"
	CreateConnectionRequestIdentityProviderSalesforce      CreateConnectionRequestIdentityProvider = "salesforce"
	CreateConnectionRequestIdentityProviderShibboleth      CreateConnectionRequestIdentityProvider = "shibboleth"
)

type UpdateConnectionRequestIdentityProvider string

const (
	UpdateConnectionRequestIdentityProviderClasslink       UpdateConnectionRequestIdentityProvider = "classlink"
	UpdateConnectionRequestIdentityProviderCyberark        UpdateConnectionRequestIdentityProvider = "cyberark"
	UpdateConnectionRequestIdentityProviderDuo             UpdateConnectionRequestIdentityProvider = "duo"
	UpdateConnectionRequestIdentityProviderGeneric         UpdateConnectionRequestIdentityProvider = "generic"
	UpdateConnectionRequestIdentityProviderGoogleworkspace UpdateConnectionRequestIdentityProvider = "google-workspace"
	UpdateConnectionRequestIdentityProviderJumpcloud       UpdateConnectionRequestIdentityProvider = "jumpcloud"
	UpdateConnectionRequestIdentityProviderKeycloak        UpdateConnectionRequestIdentityProvider = "keycloak"
	UpdateConnectionRequestIdentityProviderMiniorange      UpdateConnectionRequestIdentityProvider = "miniorange"
	UpdateConnectionRequestIdentityProviderMicrosoftentra  UpdateConnectionRequestIdentityProvider = "microsoft-entra"
	UpdateConnectionRequestIdentityProviderOkta            UpdateConnectionRequestIdentityProvider = "okta"
	UpdateConnectionRequestIdentityProviderOnelogin        UpdateConnectionRequestIdentityProvider = "onelogin"
	UpdateConnectionRequestIdentityProviderPingfederate    UpdateConnectionRequestIdentityProvider = "pingfederate"
	UpdateConnectionRequestIdentityProviderRippling        UpdateConnectionRequestIdentityProvider = "rippling"
	UpdateConnectionRequestIdentityProviderSalesforce      UpdateConnectionRequestIdentityProvider = "salesforce"
	UpdateConnectionRequestIdentityProviderShibboleth      UpdateConnectionRequestIdentityProvider = "shibboleth"
)
