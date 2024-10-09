package oidc

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v15/stytch/b2b/sso"
	"github.com/stytchauth/stytch-go/v15/stytch/methodoptions"
)

// CreateConnectionParams: Request type for `OIDC.CreateConnection`.
type CreateConnectionParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// DisplayName: A human-readable display name for the connection.
	DisplayName string `json:"display_name,omitempty"`
	// IdentityProvider: The identity provider of this connection. For OIDC, the accepted values are `generic`,
	// `okta`, and `microsoft-entra`. For SAML, the accepted values are `generic`, `okta`, `microsoft-entra`,
	// and `google-workspace`.
	IdentityProvider CreateConnectionRequestIdentityProvider `json:"identity_provider,omitempty"`
}

// UpdateConnectionParams: Request type for `OIDC.UpdateConnection`.
type UpdateConnectionParams struct {
	// OrganizationID: Globally unique UUID that identifies a specific Organization. The `organization_id` is
	// critical to perform operations on an Organization, so be sure to preserve this value.
	OrganizationID string `json:"organization_id,omitempty"`
	// ConnectionID: Globally unique UUID that identifies a specific SSO `connection_id` for a Member.
	ConnectionID string `json:"connection_id,omitempty"`
	// DisplayName: A human-readable display name for the connection.
	DisplayName string `json:"display_name,omitempty"`
	// ClientID: The OAuth2.0 client ID used to authenticate login attempts. This will be provided by the IdP.
	ClientID string `json:"client_id,omitempty"`
	// ClientSecret: The secret belonging to the OAuth2.0 client used to authenticate login attempts. This will
	// be provided by the IdP.
	ClientSecret string `json:"client_secret,omitempty"`
	// Issuer: A case-sensitive `https://` URL that uniquely identifies the IdP. This will be provided by the
	// IdP.
	Issuer string `json:"issuer,omitempty"`
	// AuthorizationURL: The location of the URL that starts an OAuth login at the IdP. This will be provided
	// by the IdP.
	AuthorizationURL string `json:"authorization_url,omitempty"`
	// TokenURL: The location of the URL that issues OAuth2.0 access tokens and OIDC ID tokens. This will be
	// provided by the IdP.
	TokenURL string `json:"token_url,omitempty"`
	// UserinfoURL: The location of the IDP's
	// [UserInfo Endpoint](https://openid.net/specs/openid-connect-core-1_0.html#UserInfo). This will be
	// provided by the IdP.
	UserinfoURL string `json:"userinfo_url,omitempty"`
	// JWKSURL: The location of the IdP's JSON Web Key Set, used to verify credentials issued by the IdP. This
	// will be provided by the IdP.
	JWKSURL string `json:"jwks_url,omitempty"`
	// IdentityProvider: The identity provider of this connection. For OIDC, the accepted values are `generic`,
	// `okta`, and `microsoft-entra`. For SAML, the accepted values are `generic`, `okta`, `microsoft-entra`,
	// and `google-workspace`.
	IdentityProvider UpdateConnectionRequestIdentityProvider `json:"identity_provider,omitempty"`
	// CustomScopes: Include a space-separated list of custom scopes that you'd like to include. Note that this
	// list must be URL encoded, e.g. the spaces must be expressed as %20.
	CustomScopes string `json:"custom_scopes,omitempty"`
	// AttributeMapping: An object that represents the attributes used to identify a Member. This object will
	// map the IdP-defined User attributes to Stytch-specific values, which will appear on the member's Trusted
	// Metadata.
	AttributeMapping map[string]any `json:"attribute_mapping,omitempty"`
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

// CreateConnectionResponse: Response type for `OIDC.CreateConnection`.
type CreateConnectionResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// Connection: The `OIDC Connection` object affected by this API call. See the
	// [OIDC Connection Object](https://stytch.com/docs/b2b/api/oidc-connection-object) for complete response
	// field details.
	Connection *sso.OIDCConnection `json:"connection,omitempty"`
}

// UpdateConnectionResponse: Response type for `OIDC.UpdateConnection`.
type UpdateConnectionResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
	// Connection: The `OIDC Connection` object affected by this API call. See the
	// [OIDC Connection Object](https://stytch.com/docs/b2b/api/oidc-connection-object) for complete response
	// field details.
	Connection *sso.OIDCConnection `json:"connection,omitempty"`
	// Warning: If it is not possible to resolve the well-known metadata document from the OIDC issuer, this
	// field will explain what went wrong if the request is successful otherwise. In other words, even if the
	// overall request succeeds, there could be relevant warnings related to the connection update.
	Warning string `json:"warning,omitempty"`
}

type CreateConnectionRequestIdentityProvider string

const (
	CreateConnectionRequestIdentityProviderGeneric         CreateConnectionRequestIdentityProvider = "generic"
	CreateConnectionRequestIdentityProviderOkta            CreateConnectionRequestIdentityProvider = "okta"
	CreateConnectionRequestIdentityProviderMicrosoftentra  CreateConnectionRequestIdentityProvider = "microsoft-entra"
	CreateConnectionRequestIdentityProviderGoogleworkspace CreateConnectionRequestIdentityProvider = "google-workspace"
)

type UpdateConnectionRequestIdentityProvider string

const (
	UpdateConnectionRequestIdentityProviderGeneric         UpdateConnectionRequestIdentityProvider = "generic"
	UpdateConnectionRequestIdentityProviderOkta            UpdateConnectionRequestIdentityProvider = "okta"
	UpdateConnectionRequestIdentityProviderMicrosoftentra  UpdateConnectionRequestIdentityProvider = "microsoft-entra"
	UpdateConnectionRequestIdentityProviderGoogleworkspace UpdateConnectionRequestIdentityProvider = "google-workspace"
)
