package saml

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v9/stytch/b2b/sso"
)

// CreateConnectionParams: Request type for `SAML.CreateConnection`.
type CreateConnectionParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// DisplayName: A human-readable display name for the connection.
	DisplayName string `json:"display_name,omitempty"`
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
	Connection sso.SAMLConnection `json:"connection,omitempty"`
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
	Connection sso.SAMLConnection `json:"connection,omitempty"`
}
